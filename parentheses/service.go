package parentheses

import (
	"math/rand"
	"strconv"
)

func parenthesesGeneration(length string) string {
	p := "[]{}()"
	g := ""
	n, _ := strconv.Atoi(length)

	for i := 0; i < n; i++ {
		g += string(p[rand.Intn(len(p))])
	}

	return g
}
