package main

import (
	"fmt"
	. "go-server/api"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	var routes = map[string]http.Handler{
		"/": fs,
		"/api/comics": http.HandlerFunc(GetComics),
		"/api/videos": http.HandlerFunc(GetVideos),
		"/api/pictures": http.HandlerFunc(GetPictures),
	};

	for path, handler := range routes {
		http.Handle(path, handler)
	}

	fmt.Println("Listening on :3000...")
	http.ListenAndServe(":3000", nil)
}
