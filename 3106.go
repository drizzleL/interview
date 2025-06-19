package main

func getSmallestString(s string, k int) string {
	b := []byte(s)
	for i := 0; k > 0 && i < len(s); i++ {
		if s[i] == 'a' {
			continue
		}
		tmp := int(s[i] - 'a')
		if tmp <= k || 26-tmp <= k {
			k -= min(26-tmp, tmp)
			b[i] = 'a'
		} else {
			b[i] -= byte(k)
			k = 0
		}

	}
	return string(b)
}
