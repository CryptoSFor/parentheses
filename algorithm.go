package parentheses

func BalancedString(s string) bool {
	list := []string{}
	flag := true

	for i := 0; i < len(s); i++ {
		v := string(s[i])
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

func isOpen(v string) string {
	if v == "{" {
		return v
	}

	if v == "[" {
		return v
	}

	if v == "(" {
		return v
	}

	return ""
}

func isClose(v string) string {
	if v == "}" {
		return v
	}

	if v == "]" {
		return v
	}

	if v == ")" {
		return v
	}

	return ""
}

func isEqual(x1, x2 string) bool {
	if x1 == "}" && x2 == "{" {
		return true
	}

	if x1 == "]" && x2 == "[" {
		return true
	}

	if x1 == ")" && x2 == "(" {
		return true
	}

	return false
}

func pop(r *[]string) string {
	list := *r
	head := list[len(list)-1]
	*r = list[:len(list)-1]

	return head
}
