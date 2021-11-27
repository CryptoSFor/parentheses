package main

import (
	"log"
	"net/http"

	"github.com/CryptoSFor/Parentheses/parentheses"
)

func main() {
	parentheses.HandleRequests()
	log.Fatal(http.ListenAndServe(":8181", nil))
}
