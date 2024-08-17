package main

import (
	"fmt"

	"github.com/sanjay-sol/p2p-DistributedFileSystem/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOptions{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		fmt.Printf("In main Error in ListenAndAccept: %s\n", err)
	}
	fmt.Println("Hello ")
	select {}
}
