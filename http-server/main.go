package main

import (
	"log"
	"net/http"
)

func main() {
    server := NewPlayerServer()
	log.Fatal(http.ListenAndServe(":5000", server))
}
