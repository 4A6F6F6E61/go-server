package main

import (
	"flag"
	. "go-server/api"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTPS network address")
  	// certFile := flag.String("certfile", "cert.pem", "certificate PEM file")
  	// keyFile := flag.String("keyfile", "key.pem", "key PEM file")
	flag.Parse()
	
	mux := http.NewServeMux()

	var routes = map[string]http.Handler{
		"/": http.FileServer(http.Dir("public")),
		"/api/comics": http.HandlerFunc(GetComics),
		"/api/videos": http.HandlerFunc(GetVideos),
		"/api/pictures": http.HandlerFunc(GetPictures),
	};

	for path, handler := range routes {
		mux.Handle(path, handler)
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	log.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
