package main

import (
	"fmt"
	"log"
	"github.com/astraxx04/distributedfilestoragesystem/p2p"
)

func main() {
	fmt.Println("Distributed File Storage System")

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}