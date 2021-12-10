// Package implements a function, that generates a random sequence of parentheses with length = n, and handler,
// that handle request to /generate?n={length}
package service

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var ErrIncorrectInput = errors.New("incorrect input")

// HandleRequests implements all handlers
func HandleRequests() {
	http.HandleFunc("/generate", GenerateHandler)
}

// GenerateHandler handles request to /generate?n={length}
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	l, err := strconv.Atoi(n)

	if err != nil || l < 0 {
		log.Println(ErrIncorrectInput)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	s := ParenthesesGeneration(l)

	_, err = w.Write(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ParenthesesGeneration generates a random sequence of parentheses of n = length
func ParenthesesGeneration(n int) []byte {
	p := "[]{}()"
	g := make([]byte, n)

	for i := 0; i < n; i++ {
		g[i] = p[rand.Intn(len(p))]
	}

	return g
}
