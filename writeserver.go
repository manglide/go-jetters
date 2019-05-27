// writeserver
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Heool Dancers</h1></body>
	</html>`
	w.Write([]byte(str))
}

func listDocuments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Service not available")
}

func externalLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.nan.ng")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Adeyemi Salau",
		Threads: []string{"First", "Second", "Third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/home", writeExample)
	http.HandleFunc("/documents", listDocuments)
	http.HandleFunc("/outside", externalLink)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
