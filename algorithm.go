package parentheses

func BalancedString(s string) bool {
	list := []rune{}
	flag := true

	for _, v := range s {
		if v == isOpen(v) {
			list = append(list, v)
		}

		if v == isClose(v) {
			if len(list) == 0 {
				flag = false
				break
			}

			head := pop(&list)

			if isEqual(v, head) {
				continue
			}

			flag = false

			break
		}
	}

	if flag && len(list) == 0 {
		return true
	}

	return false
}

func isOpen(v rune) rune {
	if string(v) == "{" {
		return v
	}

	if string(v) == "[" {
		return v
	}

	if string(v) == "(" {
		return v
	}

	return 0
}

func isClose(v rune) rune {
	if string(v) == "}" {
		return v
	}

	if string(v) == "]" {
		return v
	}

	if string(v) == ")" {
		return v
	}

	return 0
}

func isEqual(x1, x2 rune) bool {
	if string(x1) == "}" && string(x2) == "{" {
		return true
	}

	if string(x1) == "]" && string(x2) == "[" {
		return true
	}

	if string(x1) == ")" && string(x2) == "(" {
		return true
	}

	return false
}

func pop(r *[]rune) rune {
	list := *r
	head := list[len(list)-1]
	*r = list[:len(list)-1]

	return head
}
