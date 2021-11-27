// Package implements a function, that generates a random sequence of parentheses with length = n, and handler,
// that handle request to /generate?n={length}
package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func HandleRequests() {
	http.HandleFunc("/generate", GenerateHandler)
}

// GenerateHandler handles request to /generate?n={length}
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	length := r.URL.Query().Get("n")
	fmt.Fprintln(w, ParenthesesGeneration(length))
}

// ParenthesesGeneration generates a random sequence of parentheses of n = length
func ParenthesesGeneration(length string) string {
	p := "[]{}()"
	g := ""
	n, _ := strconv.Atoi(length)

	for i := 0; i < n; i++ {
		g += string(p[rand.Intn(len(p))])
	}

	return g
}
