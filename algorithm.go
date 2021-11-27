package parentheses

func BalancedString(s string) bool {
	list := []string{}

	for i := 0; i < len(s); i++ {
		v := string(s[i])
		switch v {
		case "{", "(", "[":
			list = append(list, v)

		case "}", ")", "]":
			listLength := len(list)
			if listLength == 0 {
				return false
			}

			head := list[listLength-1]

			if !pairOfParentheses(head, v) {
				return false
			}

			list = list[:listLength-1]
		}
	}

	return len(list) == 0
}

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
