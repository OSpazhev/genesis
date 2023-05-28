package handlers

import (
	"fmt"
	rateweb "github.com/ospazhev/genesis/app/web/rate"
	"github.com/ospazhev/genesis/app/web/subscription"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message []byte) {
	_, err := writer.Write(message)
	if err != nil {
		writer.WriteHeader(500)
		log.Fatal(err)
	}
}

func badStatus(writer http.ResponseWriter) {
	if err := recover(); err != nil {
		log.Println(err)
		writer.WriteHeader(400)
	}

}

func RateHandler(writer http.ResponseWriter, _ *http.Request) {
	defer badStatus(writer)
	rate, err := rateweb.GetRate()
	if err != nil {
		writer.WriteHeader(500)
	}

	message := []byte(fmt.Sprintf("%f", rate))
	write(writer, message)
}

func SubscribeHandler(writer http.ResponseWriter, r *http.Request) {
	value := r.FormValue("email")

	err := subscription.Subscribe(value)
	if err != nil {
		writer.WriteHeader(http.StatusConflict)
	}

	writer.Write([]byte{})
}
