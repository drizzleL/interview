package main

func minNumberOperations(target []int) int {
	var ret int
	var q []int
	q = append(q, 0)
	for i := 0; i < len(target); i++ {
		t := target[i]
		ret += max(0, t-q[len(q)-1])
		for q[len(q)-1] >= t {
			q = q[:len(q)-1]
		}
		q = append(q, t)
	}
	return ret
}
