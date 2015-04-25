// A test web server
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "Listen HTTP Address")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", HelloWorldServer)
	log.Fatal(http.ListenAndServe(*listen, nil))
}

func HelloWorldServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}
