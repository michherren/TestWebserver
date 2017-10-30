package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

var port int

func main() {
	flag.IntVar(&port, "port", 8080, "specify the webserver port")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf(string(requestDump))
		fmt.Fprintf(w, string(requestDump))
	})

	log.Printf("Starting server on port :%d ...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
