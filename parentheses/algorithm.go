// Package implements function BalancedString, that verifies if the given string is a balanced sequence of brackets
package parentheses

// BalancedString verifies if the given string is a balanced sequence of brackets
var brackets = map[rune]rune{'{': '}', '(': ')', '[': ']'}

func BalancedString(s string) bool {
	stack := make([]rune, 0, len([]rune(s)))

	for _, v := range s {
		switch v {
		case '{', '(', '[':
			stack = append(stack, v)

		case '}', ')', ']':
			stackLength := len(stack)
			if stackLength == 0 || v != brackets[stack[stackLength-1]] {
				return false
			}

			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}
