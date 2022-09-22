package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Println("Server running on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
