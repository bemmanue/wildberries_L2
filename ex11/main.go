package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		var b bytes.Buffer
		req.Body.Read(b.Bytes())

		io.WriteString(w, "Hello")
	}
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
