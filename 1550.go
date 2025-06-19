package main

func threeConsecutiveOdds(arr []int) bool {
	for i, cnt := 0, 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			cnt = 0
			continue
		}
		cnt += 1
		if cnt == 3 {
			return true
		}
	}
	return false
}
