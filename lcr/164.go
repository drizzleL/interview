package main

import (
	"sort"
	"strconv"
	"strings"
)

func crackPassword(password []int) string {
	strs := []string{}
	for _, p := range password {
		strs = append(strs, strconv.Itoa(p))
	}
	sort.Slice(strs, func(i, j int) bool {
		s1, s2 := strs[i], strs[j]
		return strings.Compare(s1+s2, s2+s1) < 0
	})
	return strings.Join(strs, "")
}
