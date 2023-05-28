package main

import (
	"github.com/ospazhev/genesis/app/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting web app...")
	http.HandleFunc("/rate", handlers.RateHandler)
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
