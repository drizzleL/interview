package main

func processStr2(s string, k int64) byte {
	var size int64
	var beforeSize []int64
	for _, c := range s {
		beforeSize = append(beforeSize, size)
		switch c {
		case '*':
			if size != 0 {
				size -= 1
			}
		case '%':
		case '#':
			size *= 2
		default:
			size += 1
		}
	}
	if k >= size {
		return '.'
	}
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '*':
		case '%':
			k = size - k - 1
		case '#':
			if k >= size/2 {
				k -= size / 2
			}
		default:
			if k == size-1 {
				return s[i]
			}
		}
		size = beforeSize[i]
	}
	return 0
}
