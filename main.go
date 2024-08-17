package main

import (
	"fmt"
	"github.com/sanjay-sol/p2p-DistributedFileSystem/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		fmt.Printf("In main Error in ListenAndAccept: %s\n", err)
	}
	fmt.Println("Hello ")
	select {}
}
