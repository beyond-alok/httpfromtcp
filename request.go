package httpfromtcp

import (
	"strings"
)

type Req struct {
	Method  string
	Path    string
	Version string
	Header  map[string]any
	PathValue map[string]any
	Body    string
}

func Parser(req string) *Req {
	// splits the full req into two parts - header section and body
	header_section, content_body, _ := strings.Cut(req, "\r\n\r\n")
	// splits the header section into each new line
	header_section_slice := strings.Split(header_section, "\r\n")
	// extracts the request line from header section slice
	requestLine := strings.Split(header_section_slice[0], " ")
	// extracts the only headers from header section slice
	header_slice := header_section_slice[1:]

	// intialises the header map
	header_map := make(map[string]any)

	// cuts each header line and stores it as key value in header_map
	for _, v := range header_slice {
		key, value, _ := strings.Cut(v, ": ")
		header_map[key] = value
	}

	if header_map["Content-Length"] != nil {

	}


	// intialises the req with parsed values
	parsed_req := &Req{
		Method:  requestLine[0],
		Path:    requestLine[1],
		Version: requestLine[2],
		Header:  header_map,
		Body:    content_body,
	}

	// returns parsed request
	return parsed_req

}
