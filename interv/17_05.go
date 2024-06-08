package main

func findLongestSubarray(array []string) []string {
	dict := map[int]int{}
	check := func(s string) int {
		if s[0] >= '0' && s[0] <= '9' {
			return 1
		}
		return -1
	}
	var sum int
	var ret []string
	for i := 0; i < len(array); i++ {
		sum += check(array[i])
		if sum == 0 { //  perfectly match, just update
			ret = array[:i+1]
			continue
		}
		if old, ok := dict[sum]; ok {
			if i-old > len(ret) {
				ret = array[old+1 : i+1]
			}
		} else { // not eixts, set idx
			dict[sum] = i
		}
	}
	return ret
}

func findLongestSubarray2(array []string) []string {
	ret := []string{}
	check := func(s string) int {
		if s[0] >= '0' && s[0] <= '9' {
			return 1
		}
		return -1
	}
	for i := 0; i < len(array); i += 1 {
		sum := check(array[i])
		for j := i + 1; j < len(array); j++ {
			sum += check(array[j])
			if sum == 0 && j-i+1 > len(ret) {
				ret = array[i : j+1]
			}
		}
	}
	return ret
}
