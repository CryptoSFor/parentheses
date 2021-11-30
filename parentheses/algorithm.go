// Package implements function BalancedString, that verifies if the given string is a balanceds sequence of brackets
package parentheses

// BalancedString verifies if the given string is a balanceds sequence of brackets
var brackets = map[string]string{"{": "}", "(": ")", "[": "]"}

func BalancedString(s string) bool {
	list := []string{}

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		switch v {
		case "{", "(", "[":
			list = append(list, v)

		case "}", ")", "]":
			listLength := len(list)
			if listLength == 0 || v != brackets[list[listLength-1]] {
				return false
			}

			list = list[:listLength-1]
		}
	}

	return len(list) == 0
}

// pairOfParentheses checks if param x1 is an open parenthesis and x2 is a close parenthesis of one type
func pairOfParentheses(x1, x2 string) bool {
	if x1 == "{" && x2 == "}" {
		return true
	}

	if x1 == "[" && x2 == "]" {
		return true
	}

	if x1 == "(" && x2 == ")" {
		return true
	}

	return false
}
