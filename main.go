package main

import (
	"log"
	"net/http"

	"github.com/CryptoSFor/Parentheses/service"
)

func main() {
	service.HandleRequests()
	log.Fatal(http.ListenAndServe(":8181", nil))
}
