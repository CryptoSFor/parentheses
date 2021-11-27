// Package implements a function, that generates a random sequence of parentheses with length = n, and handler,
// that handle request to /generate?n={length}
package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

// HandleRequests implements all handlers
func HandleRequests() {
	http.HandleFunc("/generate", GenerateHandler)
}

// GenerateHandler handles request to /generate?n={length}
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")
	fmt.Fprintln(w, ParenthesesGeneration(n))
}

// ParenthesesGeneration generates a random sequence of parentheses of n = length
func ParenthesesGeneration(n string) string {
	p := "[]{}()"
	g := ""
	length, _ := strconv.Atoi(n)

	for i := 0; i < length; i++ {
		g += string(p[rand.Intn(len(p))])
	}

	return g
}
