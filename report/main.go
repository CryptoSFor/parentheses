package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/CryptoSFor/Parentheses/parentheses"
)

const url = "http://localhost:8181/generate?n="

func main() {
	N := 1000

	for i := 2; i <= 8; i *= 2 {
		var counter int
		for j := 0; j < N; j++ {
			s, err := GetParentheses(url + strconv.Itoa(i))
			if err != nil {
				log.Println(err)
				continue
			}

			if parentheses.BalancedString(s) {
				counter++
			}
		}

		p := float64(counter) * 100 / float64(N)
		fmt.Printf("n = %d, %.1f%%\n", i, p)
	}
}

func GetParentheses(url string) (string, error) {
	r, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	s, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
