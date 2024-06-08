package main

import (
	"log"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	type trieNode struct {
		flag     bool
		children [26]*trieNode
	}
	rootTrie := &trieNode{}
	var insert func(node *trieNode, w string, i int)
	insert = func(node *trieNode, w string, i int) {
		if i == len(w) {
			node.flag = true
			return
		}
		child := node.children[w[i]-'a']
		if child == nil {
			node.children[w[i]-'a'] = &trieNode{}
		}
		insert(node.children[w[i]-'a'], w, i+1)
	}
	for _, w := range dictionary {
		insert(rootTrie, w, 0)
	}
	log.Println(rootTrie)
	var find func(node *trieNode, w string, i int) string
	find = func(node *trieNode, w string, i int) string {
		if node.flag {
			return w[:i]
		}
		if i == len(w) {
			return w
		}
		child := node.children[w[i]-'a']
		if child == nil {
			return w
		}
		return find(child, w, i+1)
	}

	words := strings.Split(sentence, " ")
	for i, w := range words {
		words[i] = find(rootTrie, w, 0)
	}
	return strings.Join(words, " ")
}
