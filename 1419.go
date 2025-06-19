package main

func minNumberOfFrogs(croakOfFrogs string) int {
	dict := map[byte]int{}
	for i, c := range "croak" {
		dict[byte(c)] = i
	}
	var calls [5]int
	var ret, rest int
	for _, c := range croakOfFrogs {
		i := dict[byte(c)]
		if i == 0 {
			calls[0] += 1
			if rest != 0 {
				rest -= 1
			} else {
				ret += 1
			}
			continue
		}
		if calls[i-1] == 0 {
			return -1
		}
		calls[i-1] -= 1
		calls[i] += 1
		if i == 4 {
			calls[i] -= 1
			rest += 1
		}
	}
	for _, v := range calls {
		if v != 0 {
			return -1
		}
	}
	return ret
}
