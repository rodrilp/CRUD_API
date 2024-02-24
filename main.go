package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", healthHandler)

	log.Println("Server starting at port -> :4444")

	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatal("server error ", err)
	}

	log.Println("server is stopped")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server is running..."))
}