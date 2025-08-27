package main

import "sort"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	busDict := map[string]int{
		"electronics": 1,
		"grocery":     2,
		"pharmacy":    3,
		"restaurant":  4,
	}
	var ret []string
	type item struct {
		code string
		bus  int
	}
	var items []item
	for i := range code {
		if !isActive[i] {
			continue
		}
		if busDict[businessLine[i]] == 0 {
			continue
		}
		if len(code[i]) == 0 {
			continue
		}
		var codeFlag bool
		for _, c := range code[i] {
			if c >= 'A' && c <= 'Z' {
				continue
			}
			if c >= 'a' && c <= 'z' {
				continue
			}
			if c >= '0' && c <= '9' {
				continue
			}
			if c == '_' {
				continue
			}
			codeFlag = true
			break
		}
		if codeFlag {
			continue
		}
		items = append(items, item{
			code: code[i],
			bus:  busDict[businessLine[i]],
		})
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].bus == items[j].bus {
			return items[i].code < items[j].code
		}
		return items[i].bus < items[j].bus
	})
	for _, it := range items {
		ret = append(ret, it.code)
	}
	return ret
}
