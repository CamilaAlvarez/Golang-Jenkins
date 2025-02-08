package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/example/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	log.Print("Listening ...")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
