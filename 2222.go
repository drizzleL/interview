package main

func numberOfWays(s string) int64 {
	var preOne, preZero, suffOne, suffZero int
	for _, c := range s {
		switch c {
		case '0':
			suffZero += 1
		case '1':
			suffOne += 1
		}
	}
	var ret int
	for i := 0; i < len(s)-1; i++ {
		switch s[i] {
		case '0':
			suffZero -= 1
			ret += preOne * suffOne
			preZero += 1
		case '1':
			suffOne -= 1
			ret += preZero * suffZero
			preOne += 1
		}
	}
	return int64(ret)
}
