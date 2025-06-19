package main

func canVisitAllRooms(rooms [][]int) bool {
	var cnt int
	q := []int{0}
	seen := make([]bool, len(rooms))
	for len(q) > 0 {
		now := q[len(q)-1]
		q = q[:len(q)-1]
		seen[now] = true
		cnt += 1
		for _, k := range rooms[now] {
			if seen[k] {
				continue
			}
			seen[k] = true
			q = append(q, k)
		}
	}
	return cnt == len(rooms)
}
