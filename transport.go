package httpfromtcp

import (
	"fmt"
	"net"
)

func Listener(server *Server, listen_addr string) Error {

	// defines the listener
	listener, err := net.Listen("tcp", listen_addr)
	if err != nil {
		return Error{
			Msg: "Error while listening tcp",
			Err: err,
		}
	}

	// Running listener in infinite for loop - to accect requests
	for {

		fmt.Println("Listening on port : ",server.Listen_addr)
		// accepts the request
		conn, err := listener.Accept()
		fmt.Println("conn accepted", conn)
		if err != nil {
			fmt.Println(Error{
				Msg: "Error while accepting request",
				Err: err,
			})
		}

		go HandleConnection(server,conn)

	}

}

func HandleConnection(server *Server, conn net.Conn) {
	for {
		// initialises the response writer
		var res_writer ResponseWriter
		// sets conn to server struct 
		res_writer.conn = conn
		// initialises the header map
		res_writer.res.Header = make(map[string]any)
		// creates a buffer to store the req msg temporarily to read
		buff := make([]byte, 2000)

		// reads the req and write it in buffer
		n, err := conn.Read(buff)
		if err != nil {
			return
		}
		req := string(buff[:n])
		// parses the req
		parsed_req := Parse_Request(req)
		// matchs the routes
		fn := MatchRoutes(parsed_req, server)
		
		if fn == nil {
			continue
		}
		// executes the function
		fn(parsed_req,&res_writer)

	}

}
