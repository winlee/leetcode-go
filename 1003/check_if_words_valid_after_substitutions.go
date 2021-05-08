package _003

func isValid1003(S string) bool {
	if len(S) < 3 {
		return false
	}

	var stack []byte

	for _, ch := range []byte(S) {
		if len(stack) == 0 || ch == 'a' {
			stack = append(stack, ch)
		} else if ch == 'b' {
			if len(stack) > 0 && stack[len(stack)-1] == 'a' {
				stack = append(stack, ch)
			} else {
				return false
			}
		} else {
			if len(stack) > 1 && stack[len(stack)-1] == 'b' && stack[len(stack)-2] == 'a' {
				stack = stack[:len(stack)-2]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
