package p2p

import ( 
  "net"
  "fmt"
)

type TCPTransport struct {
	listenAddress string
  listener net.Listener
	peers map[net.Addr]Peer
}


func Prints()  {
   fmt.Println("Transporter")
}

