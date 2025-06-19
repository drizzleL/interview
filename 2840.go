package main

func checkStrings(s1 string, s2 string) bool {
	helper := func(s string) (even, odd [26]int) {
		for i := 0; i < len(s1); i++ {
			c := int(s[i] - 'a')
			if i%2 == 0 {
				even[c] += 1
			} else {
				odd[c] += 1
			}
		}
		return even, odd
	}
	even1, odd1 := helper(s1)
	even2, odd2 := helper(s2)
	return even1 == even2 && odd1 == odd2
}
