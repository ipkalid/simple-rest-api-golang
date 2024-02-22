package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	server := http.NewServeMux()
	server.HandleFunc("/", basicHandler)
	error := http.ListenAndServe(":3000", server)
	if error != nil {
		panic(error.Error())
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}
