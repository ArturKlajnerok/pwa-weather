package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var srv http.Server
	http2.VerboseLogs = true
	srv.Addr = ":8080"
	srv.Handler = http.FileServer(http.Dir(""))
	http2.ConfigureServer(&srv, nil)

	log.Fatal(srv.ListenAndServeTLS("localhost.cert", "localhost.key"))
}
