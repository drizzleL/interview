package main

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	curr := shorter * k
	ret := []int{curr}
	gap := longer - shorter
	if gap == 0 {
		return ret
	}
	for i := 0; i < k; i++ {
		curr += gap
		ret = append(ret, curr)
	}
	return ret
}
