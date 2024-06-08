package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func trulyMostPopular(names []string, synonyms []string) []string {
	parent := make([]int, len(names))
	for i := range parent {
		parent[i] = i
	}
	dict := map[string]int{}
	cntDict := make([]int, len(names))
	nameArr := make([]string, len(names))
	for i, str := range names {
		var k int
		for ; str[k] != '('; k++ {
		}
		name := str[:k]
		cntStr := str[k+1 : len(str)-1]
		d, _ := strconv.Atoi(cntStr)
		dict[name] = i
		cntDict[i] = d
		nameArr[i] = name
	}
	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}
	union := func(i, j int) {
		if i > j {
			i, j = j, i
		}
		parent[find(j)] = find(i)
	}
	for _, v := range synonyms {
		names := strings.Split(v[1:len(v)-1], ",")
		i, j := dict[names[0]], dict[names[1]]
		union(i, j)
	}
	type ele struct {
		keys []string
		cnt  int
	}
	eleDict := map[int]*ele{}
	for i := range parent {
		pi := find(i)
		e, ok := eleDict[pi]
		if !ok {
			e = &ele{}
			eleDict[pi] = e
		}
		e.keys = append(e.keys, nameArr[i])
		e.cnt += cntDict[i]
	}
	var ret []string
	for _, e := range eleDict {
		sort.Strings(e.keys)
		ret = append(ret, fmt.Sprintf("%s(%d)", e.keys[0], e.cnt))
	}
	return ret
}
