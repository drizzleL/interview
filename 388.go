package main

import (
	"strings"
)

func lengthLongestPath(input string) int {
	words := strings.Split(input, "\n")
	var stack []int
	var ret int
	for _, word := range words {
		var i int
		for i < len(word) && word[i:i+1] == "\t" {
			i += 1
		}
		stack = stack[:i]
		if strings.Contains(word, ".") { // file
			var curr int
			for _, v := range stack {
				curr += v
			}
			ret = max(ret, curr+len(word)-i+len(stack))
		} else { // dir
			stack = append(stack, len(word)-i)
		}
	}
	return ret
}
