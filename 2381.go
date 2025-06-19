package main

func shiftingLetters(s string, shifts [][]int) string {
	diff := make([]int, len(s)+1)
	for _, sh := range shifts {
		if sh[2] == 1 {
			diff[sh[0]] += 1
			diff[sh[1]+1] -= 1
		} else {
			diff[sh[0]] -= 1
			diff[sh[1]+1] += 1
		}
	}
	b := []byte(s)
	var sum int
	for i, c := range b {
		sum += diff[i]
		tmp := int(c - 'a')
		tmp += sum
		tmp %= 26
		if tmp < 0 {
			tmp += 26
		}
		b[i] = byte(tmp) + 'a'
	}
	return string(b)
}
