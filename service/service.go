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

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	length := r.URL.Query().Get("length")
	fmt.Fprintln(w, parenthesesGeneration(length))
}

func parenthesesGeneration(length string) string {
	p := "[]{}()"
	g := ""
	n, _ := strconv.Atoi(length)

	for i := 0; i < n; i++ {
		g += string(p[rand.Intn(len(p))])
	}

	return g
}
