package main

import (
	"sort"
	"strings"
)

func largestWordCount(messages []string, senders []string) string {
	dict := map[string]int{}
	var topCnt int
	for i, msg := range messages {
		dict[senders[i]] += len(strings.Fields(msg))
		topCnt = max(topCnt, dict[senders[i]])
	}
	var names []string
	for k, v := range dict {
		if v == topCnt {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	return names[len(names)-1]
}
