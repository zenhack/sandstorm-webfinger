package main

import (
	"golang.org/x/net/context"
	"net/http"
	"zenhack.net/go/sandstorm/sandstormhttpbridge"
)

func main() {
	http.HandleFunc("/.well-known/webfinger",
		func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(404)
		})
	sandstormhttpbridge.ListenAndServe(context.Background(), nil, ":8000", nil)
}
