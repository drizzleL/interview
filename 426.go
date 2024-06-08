package main

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	type item struct {
		ele string
		cnt int
	}
	var items []*item
	dict := map[string]*item{}
	add := func(ele string, cnt int) {
		if dict[ele] == nil {
			dict[ele] = &item{
				ele: ele,
			}
			items = append(items, dict[ele])
		}
		dict[ele].cnt += cnt
	}
	tmp := []int{1}
	for i := len(formula) - 1; i >= 0; i-- {
		c := formula[i]
		switch {
		case c >= 'a' && c <= 'z',
			c >= 'A' && c <= 'Z':
			j := i
			for ; formula[j] >= 'a' && formula[j] <= 'z'; j-- {
			}
			ele := formula[j : i+1]
			add(ele, tmp[len(tmp)-1])
			i = j
		case c == '(':
			tmp = tmp[:len(tmp)-1]
		case c >= '0' && c <= '9':
			j := i
			for formula[j] >= '0' && formula[j] <= '9' {
				j--
			}
			m, _ := strconv.Atoi(formula[j+1 : i+1])
			mm := tmp[len(tmp)-1] * m
			i = j
			if formula[j] == ')' {
				tmp = append(tmp, mm)
				i = j
				continue
			}
			for ; formula[j] >= 'a' && formula[j] <= 'z'; j-- {
			}
			ele := formula[j : i+1]
			add(ele, mm)
			i = j
		case c == ')':
			tmp = append(tmp, tmp[len(tmp)-1])
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return strings.Compare(items[i].ele, items[j].ele) < 0
	})
	var b []byte
	for _, it := range items {
		b = append(b, []byte(it.ele)...)
		if it.cnt != 1 {
			b = append(b, []byte(strconv.Itoa(it.cnt))...)
		}
	}
	return string(b)
}
