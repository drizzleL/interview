package main

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	vis := map[string]bool{
		"0000": true,
	}
	for _, d := range deadends {
		if d == "0000" {
			return -1
		}
		vis[d] = true
	}
	times := 1
	nodes := []string{"0000"}
	getChild := func(str string) []string {
		var ret []string
		b := []byte(str)
		for i := range str {
			tmp := b[i]
			if tmp == '0' {
				b[i] = '9'
			} else {
				b[i]--
			}
			ret = append(ret, string(b))
			b[i] = tmp
			if tmp == '9' {
				b[i] = '0'
			} else {
				b[i]++
			}
			ret = append(ret, string(b))
			b[i] = tmp
		}
		return ret
	}
	for len(nodes) > 0 {
		var nextnodes []string
		for _, n := range nodes {
			for _, child := range getChild(n) {
				if vis[child] {
					continue
				}
				if target == child {
					return times
				}
				vis[child] = true
				nextnodes = append(nextnodes, child)
			}
		}
		nodes = nextnodes
		times++
	}
	return -1
}
