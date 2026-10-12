package httpfromtcp

import "fmt"

type Server struct {
	Listen_addr string
	Routes      map[string]HandleFunc
	ResWriter ResponseWriter
	ch chan string
}

type HandleFunc func(*Req, *ResponseWriter)

// NewServer
func NewServer() Server {
	server := &Server{
		Routes: make(map[string]HandleFunc),
		ResWriter: ResponseWriter{},
		ch :make(chan string),
	}

	// initialises the header map
	server.ResWriter.res.Header = make(map[string]any)

	return *server

}

// Handle
func (server *Server) Handle(method_path string, handlefunc HandleFunc) {
	fmt.Println("routes starting to registering")
	server.Routes[method_path] = handlefunc
	fmt.Println("routes registered")
}

// ListenAndServe
func (server *Server) ListenAndServe(listen_addr string) {
	server.Listen_addr = listen_addr

	go Listen(server)
	
	for {
		req := <- server.ch

		// parses req
		parsedReq := Parser(req)

		// matchs routes
		fn := MatchRoutes(parsedReq, server.Routes)
		
		if fn != nil {
			// execute the function
			fn(parsedReq,&server.ResWriter)
		}
	}

}




