package httpfromtcp

import (
	"net"
)

func Listen(server *Server) {

	listener, err := net.Listen("tcp", server.Listen_addr)
	if err != nil {
		panic(err)
	}	

	// Running listener in infinite for loop - to accect requests
	for {

		// accepts the request
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(server,conn)

	}

}

func handleConnection(server *Server, conn net.Conn) {
	defer conn.Close()
	for {
		
		// creates a buffer to store the req msg temporarily to read
		buff := make([]byte, 2000)

		// reads the req and write it in buffer
		n, err := conn.Read(buff)
		if err != nil {
			return
		}

		req := string(buff[:n])

		server.ResWriter.conn = conn

		server.ch <- req 

	}

}
