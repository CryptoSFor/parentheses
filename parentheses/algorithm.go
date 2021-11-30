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
