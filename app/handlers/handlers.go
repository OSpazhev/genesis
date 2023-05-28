package handlers

import (
	"fmt"
	rateweb "github.com/ospazhev/genesis/app/web/rate"
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
