package main

import "strings"

func replaceWords(dictionary []string, sentence string) string {
	type trieNode struct {
		flag     bool
		children [26]*trieNode
	}
	root := &trieNode{}
	for _, s := range dictionary {
		n := root
		for _, c := range s {
			if n.children[c-'a'] == nil {
				n.children[c-'a'] = &trieNode{}
			}
			n = n.children[c-'a']
		}
		n.flag = true
	}
	helper := func(word string) string {
		n := root
		for i, c := range word {
			n = n.children[c-'a']
			if n == nil {
				break
			}
			if n.flag {
				return word[:i+1]
			}
		}
		return word
	}
	words := strings.Fields(sentence)
	for i, w := range words {
		words[i] = helper(w)
	}
	return strings.Join(words, " ")
}
