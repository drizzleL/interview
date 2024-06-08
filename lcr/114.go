package main

func alienOrder(words []string) string {
	allCh := map[byte]bool{}
	for _, w := range words {
		for _, c := range w {
			allCh[byte(c)] = true
		}
	}
	if len(allCh) == 1 {
		return words[0]
	}
	var seqs [][]byte
	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		var flag bool
		for j := 0; j < len(w1) && j < len(w2); j++ {
			if w1[j] == w2[j] {
				continue
			}
			flag = true
			seqs = append(seqs, []byte{w1[j], w2[j]})
			break
		}
		if !flag && len(w1) > len(w2) {
			return ""
		}
	}
	ins := map[byte]int{}
	dict := map[byte][]byte{}
	seen := map[byte]bool{}
	for _, seq := range seqs {
		dict[seq[0]] = append(dict[seq[0]], seq[1])
		ins[seq[1]]++
		seen[seq[0]] = true
	}
	nodes := []byte{}
	for v := range seen {
		if ins[v] == 0 {
			nodes = append(nodes, v)
		}
	}
	ret := []byte{}
	for len(nodes) != 0 {
		var next []byte
		for _, n := range nodes {
			ret = append(ret, n)
			for _, n := range dict[n] {
				ins[n]--
				if ins[n] == 0 {
					next = append(next, n)
				}
			}
		}
		nodes = next
	}
	for _, v := range ins {
		if v != 0 {
			return ""
		}
	}
	for _, c := range ret {
		allCh[c] = false
	}
	for c, ok := range allCh {
		if !ok {
			continue
		}
		ret = append(ret, c)
	}
	return string(ret)
}
