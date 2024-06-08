package main

func smallestString(s string) string {
	var b []byte
	var i int
	var flag bool
	for i < len(s) && s[i] == 'a' {
		b = append(b, s[i])
		i++
	}
	for i < len(s) && s[i] != 'a' {
		flag = true
		b = append(b, s[i]-1)
		i++
	}
	for i < len(s) {
		b = append(b, s[i])
		i++
	}
	if !flag { // all a
		b[len(b)-1] = 'z'
	}
	return string(b)
}
