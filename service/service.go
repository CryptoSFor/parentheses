// Package implements a function, that generates a random sequence of parentheses with length = n, and handler,
// that handle request to /generate?n={length}
package service

import (
	"errors"
	"math/rand"
	"net/http"
	"strconv"
)

var ErrNegativeLength = errors.New("negative length")

// HandleRequests implements all handlers
func HandleRequests() {
	http.HandleFunc("/generate", GenerateHandler)
}

// GenerateHandler handles request to /generate?n={length}
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	l, err := strconv.Atoi(n)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	s, err := ParenthesesGeneration(l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	_, err = w.Write([]byte(s))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// ParenthesesGeneration generates a random sequence of parentheses of n = length
func ParenthesesGeneration(n int) (string, error) {
	if n < 0 {
		return "", ErrNegativeLength
	}

	p := "[]{}()"
	g := make([]byte, n)

	for i := 0; i < n; i++ {
		g[i] = p[rand.Intn(len(p))]
	}

	return string(g), nil
}
