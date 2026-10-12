package httpfromtcp

import (
	"fmt"
	"strings"
)

func MatchRoutes(req *Req , routes map[string]HandleFunc) HandleFunc {

	// initialises matched variable with false value
	var matched bool = false

	// gives method path
	methodPath :=  req.Method + " " + req.Path

	// splits method path by /
	method_Path_slice := strings.Split(methodPath,"/")

	for key,value := range routes {
		// spilits the method_path/key
		key_slice := strings.Split(key,"/")

		// initialises the PathValue map
		req.PathValue = make(map[string]any)

		for i := range key_slice {
			k_length := len(key_slice[i])
			if k_length == 0 {
				continue
			}
			
			fmt.Printf("key : %v || length: %v\n",key_slice[i],k_length)
			if key_slice[i][0] == 123 {
				req.PathValue[key_slice[i][1:k_length-1]] = method_Path_slice[i]
				continue
			}
			if method_Path_slice[i] == key_slice[i] {
				matched = true
			} else {
				matched = false
				break
			}
		}
		if matched == true {
			return value
		} 
	}
	return nil

}
