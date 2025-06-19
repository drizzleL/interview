package main

func maximumNumber(num string, change []int) string {
	b := []byte(num)
	for i := range b {
		d := int(b[i] - '0')
		if change[d] <= d {
			continue
		}
		for i < len(b) {
			d := int(b[i] - '0')
			if change[d] < d {
				break
			}
			b[i] += byte(change[d] - d)
			i += 1
		}
		break
	}
	return string(b)
}
