package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := "Hello World!"
	vec := []byte(message)
	// Add message to the response
	_, err := writer.Write(vec)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// register hander
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
