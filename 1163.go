package main

func lastSubstring(s string) string {
	type Item struct {
		start int
		now   int
	}
	var maxCh int
	for _, c := range s {
		maxCh = max(maxCh, int(c))
	}
	var items []*Item
	seen := map[int]bool{}
	for i, c := range s {
		if int(c) != maxCh {
			continue
		}
		if i != 0 && int(s[i-1]) == int(c) {
			continue
		}
		items = append(items, &Item{
			start: i,
			now:   i,
		})
		seen[i] = true
	}
	for len(items) > 1 {
		var maxCh int
		var flag bool
		for _, item := range items {
			item.now += 1
			if item.now == len(s) {
				continue
			}
			if seen[item.now] {
				items = append(items[:0], item)
				flag = true
				break
			}
			maxCh = max(maxCh, int(s[item.now]))
		}
		if flag {
			break
		}
		var next []*Item
		for _, item := range items {
			if item.now == len(s) {
				continue
			}
			if int(s[item.now]) == maxCh {
				next = append(next, item)
			}
		}
		items = next
	}
	return s[items[0].start:]
}
