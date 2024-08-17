package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOptions struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
		fmt.Printf("New incoming connection from %+v\n", conn)
		go t.handleConn(conn)
	}
}

// type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP Handshake error: %s\n", err)
		return
	}

	msg := &Message{}
	// buf := make([]byte, 2000)
	for {
		// n, err := conn.Read(buf)

		// if err != nil {
		// 	fmt.Printf("TCP error: %s\n", err)
		// }
		// write some data to the connection
		// _, err := conn.Write([]byte("Hello from server"))

		// if err != nil {
		// 	fmt.Printf("Write error: %s\n", err)
		// 	continue
		// }

		// if err := t.Decoder.Decode(conn, msg); err != nil {
		// 	fmt.Printf("Decode error: %s\n", err)
		// 	continue
		// }

		// fmt.Printf("Received message: %s\n", string(msg.Payload))


		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Decode error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()
		
		fmt.Printf("MEssage received from %s\n", msg.From)
		fmt.Printf("Received message: %s\n", string(msg.Payload))
	}
}
 