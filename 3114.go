package main

func findLatestTime(s string) string {
	b := []byte(s)
	if b[3] == '?' {
		b[3] = '5'
	}
	if b[4] == '?' {
		b[4] = '9'
	}
	switch {
	case b[0] == '?' && b[1] == '?':
		b[0] = '1'
		b[1] = '1'
	case b[0] == '0' && b[1] == '?':
		b[1] = '9'
	case b[0] == '1' && b[1] == '?':
		b[1] = '1'
	case b[0] == '?' && b[1] >= '2':
		b[0] = '0'
	case b[0] == '?' && b[1] < '2':
		b[0] = '1'
	}
	return string(b)
}
