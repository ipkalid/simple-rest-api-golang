package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello-world", helloWorldHandler)

	fmt.Println("server started at: http://localhost:3000")
	error := http.ListenAndServe(":3000", router)
	if error != nil {
		panic(error.Error())
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}
