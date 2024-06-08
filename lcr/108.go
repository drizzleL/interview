package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// preprocess
	dict := map[string][]int{}
	var flag bool
	for idx, word := range wordList {
		if word == endWord {
			flag = true
		}
		if word == beginWord {
			continue
		}
		b := []byte(word)
		for i := 0; i < len(word); i++ {
			tmp := b[i]
			b[i] = 'A'
			dict[string(b)] = append(dict[string(b)], idx)
			b[i] = tmp
		}
	}
	if !flag {
		return 0
	}
	getNext := func(s string) []string {
		b := []byte(s)
		var ret []string
		for i := 0; i < len(s); i++ {
			tmp := b[i]
			b[i] = 'A'
			ret = append(ret, string(b))
			b[i] = tmp
		}
		return ret
	}
	var viscnt int
	vis := make([]bool, len(wordList))
	nodes := []string{beginWord}
	times := 1
	for len(nodes) != 0 {
		var nextnodes []string
		for _, n := range nodes {
			for _, child := range getNext(n) {
				nextwords := dict[child]
				for _, wIdx := range nextwords {
					if vis[wIdx] {
						continue
					}
					if wordList[wIdx] == endWord {
						return times + 1
					}
					vis[wIdx] = true
					viscnt++
					nextnodes = append(nextnodes, wordList[wIdx])
				}
			}
		}
		if viscnt == len(wordList) { // all used
			break
		}
		nodes = nextnodes
		times++
	}
	return 0
}
