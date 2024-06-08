package main

func wonderfulSubstrings(word string) int64 {
	dict := map[int]int{0: 1}
	var mask int
	var ret int
	for _, c := range word {
		d := int(c - 'a')
		mask ^= 1 << d
		ret += dict[mask]
		for i := 0; i < 10; i++ {
			ret += dict[1<<i^mask]
		}
		dict[mask] += 1
	}
	return int64(ret)
}
