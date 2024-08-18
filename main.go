package main

import (
	"fmt"

	"github.com/sanjay-sol/p2p-DistributedFileSystem/p2p"
)

func OnPeer(p p2p.Peer) error {
	fmt.Println("Working with new Peer")
	p.Close()
	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Received message: %s\n", string(msg.Payload))

		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		fmt.Printf("In main Error in ListenAndAccept: %s\n", err)
	}
	fmt.Println("Hello ")
	select {}
}
