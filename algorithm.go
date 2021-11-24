package parentheses

import "errors"

var ErrEmptySlice = errors.New("length of slice is 0")

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

			head, _ := pop(&list)

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

func pop(r *[]string) (string, error) {
	if len(*r) == 0 {
		return "the slice is empty", ErrEmptySlice
	}

	list := *r
	head := list[len(list)-1]
	*r = list[:len(list)-1]

	return head, nil
}
