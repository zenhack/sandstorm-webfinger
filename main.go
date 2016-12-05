package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"html/template"
	"log"
	"net/http"
	"os"
	"zenhack.net/go/sandstorm/sandstormhttpbridge"
)

var (
	assetPath = os.Getenv("ASSET_PATH")
	indexTpl  *template.Template
)

func init() {
	indexTpl = template.Must(template.ParseFiles(assetPath + "templates/index.html"))
}

var db = map[string]Jrd{
	"acct:alice@example.net": {
		Subject: "acct:alice@example.net",
	},
}

func main() {
	http.HandleFunc("/.well-known/webfinger",
		func(w http.ResponseWriter, req *http.Request) {
			resource := req.URL.Query().Get("resource")
			ret, ok := db[resource]
			if !ok {
				w.WriteHeader(404)
				return
			}
			w.Header().Set("Content-Type", "application/jrd+json")
			enc := json.NewEncoder(w)
			err := enc.Encode(ret)
			log.Println(err)
		})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		indexTpl.Execute(w, struct{}{})
	})
	http.Handle("/static/", http.FileServer(http.Dir(assetPath)))
	sandstormhttpbridge.ListenAndServe(context.Background(), nil, ":8000", nil)
}
