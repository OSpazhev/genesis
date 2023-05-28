package main

import (
	"github.com/ospazhev/genesis/app/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web app...")
	http.HandleFunc("/rate", handlers.RateHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
