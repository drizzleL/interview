package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var q []int
	var ret int
	keyGot := make([]bool, len(status))
	boxFound := make([]bool, len(status))
	used := make([]bool, len(status))
	for _, box := range initialBoxes {
		boxFound[box] = true
		if status[box] == 1 {
			q = append(q, box)
			used[box] = true
		}
	}
	for len(q) > 0 {
		top := q[len(q)-1]
		q = q[:len(q)-1]
		ret += candies[top]
		for _, key := range keys[top] {
			keyGot[key] = true
			if boxFound[key] && !used[key] {
				used[key] = true
				q = append(q, key)
			}
		}
		for _, childBox := range containedBoxes[top] {
			boxFound[childBox] = true
			if !used[childBox] && (status[childBox] == 1 || keyGot[childBox]) {
				q = append(q, childBox)
				used[childBox] = true
			}
		}
	}
	return ret
}
