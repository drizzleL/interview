package main

func flowerGame(n int, m int) int64 {
	odd1, odd2 := (n+1)/2, (m-1)/2
	even1, even2 := n/2, m/2
	return int64(odd1*even2 + odd2*even1)
}
