package main

func reverseParentheses(s string) string {
	var b []byte
	var lefts []int
	for _, c := range s {
		switch c {
		case '(':
			lefts = append(lefts, len(b))
		case ')':
			lidx := lefts[len(lefts)-1]
			lefts = lefts[:len(lefts)-1]
			for i, j := lidx, len(b)-1; i < j; i, j = i+1, j-1 {
				b[i], b[j] = b[j], b[i]
			}
		default:
			b = append(b, byte(c))
		}
	}
	return string(b)
}
