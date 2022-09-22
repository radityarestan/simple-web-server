package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
