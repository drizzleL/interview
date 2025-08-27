package main

func processStr(s string) string {
	var ret []byte
	for _, c := range s {
		switch c {
		case '*':
			if len(ret) != 0 {
				ret = ret[:len(ret)-1]
			}
		case '%':
			for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
				ret[i], ret[j] = ret[j], ret[i]
			}
		case '#':
			ret = append(ret, ret...)
		default:
			ret = append(ret, byte(c))
		}
	}
	return string(ret)
}
