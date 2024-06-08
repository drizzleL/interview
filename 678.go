package main

func checkValidString(s string) bool {
	type item struct {
		val rune
		cnt int
	}
	var stack []item
	for _, c := range s {
		switch c {
		case '(', '*':
			if len(stack) == 0 || stack[len(stack)-1].val != c {
				stack = append(stack, item{
					val: c,
					cnt: 1,
				})
				continue
			}
			stack[len(stack)-1].cnt += 1
		case ')':
			if len(stack) >= 1 && stack[len(stack)-1].val == '(' {
				stack[len(stack)-1].cnt -= 1
				if stack[len(stack)-1].cnt == 0 {
					stack = stack[:len(stack)-1]
				}
				continue
			}
			if len(stack) >= 2 && stack[len(stack)-2].val == '(' {
				stack[len(stack)-2].cnt -= 1
				if stack[len(stack)-2].cnt == 0 {
					if len(stack) >= 3 {
						stack[len(stack)-3].cnt += stack[len(stack)-1].cnt
						stack = stack[:len(stack)-2]
						continue
					}
					stack = append(stack[:len(stack)-2], stack[len(stack)-1])
				}
				continue
			}
			if len(stack) == 0 {
				return false
			}
			stack[len(stack)-1].cnt -= 1
			if stack[len(stack)-1].cnt == 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}
	var left int
	for _, v := range stack {
		switch v.val {
		case '*':
			left = max(0, left-v.cnt)
		case '(':
			left += v.cnt
		}
	}
	return left == 0
}
