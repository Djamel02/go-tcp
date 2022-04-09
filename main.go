package main

import (
	"log"
	"net"
)

func main()  {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp",":8888")
	if err != nil {
		log.Fatalf("Cannot start server on %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Server started on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("cannot accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}