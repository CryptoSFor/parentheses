// Package implements function BalancedString, that verifies if the given string is a balanced sequence of brackets
package parentheses

// BalancedString verifies if the given string is a balanced sequence of brackets
var brackets = map[string]string{"{": "}", "(": ")", "[": "]"}

func BalancedString(s string) bool {
	stack := make([]string, 0, 16)

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		switch v {
		case "{", "(", "[":
			stack = append(stack, v)

		case "}", ")", "]":
			stackLength := len(stack)
			if stackLength == 0 || v != brackets[stack[stackLength-1]] {
				return false
			}

			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}
