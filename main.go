package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color: red'>Hello World</h1>"))
}

func main() {
	http.HandleFunc("/hello-world", Hello)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
