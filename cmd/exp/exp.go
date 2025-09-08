package main

import (
	"flag"
	"fmt"
)

func main() {
	var port string
	var verbose bool

	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	if verbose {
		fmt.Println("Server is running on port " + port)
	} else {
		fmt.Println(port)
	}
}
