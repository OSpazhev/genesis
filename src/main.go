package main

import (
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, _ *http.Request) {
	message := []byte("Hello, my friend")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Starting web app...")
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
