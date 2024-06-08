package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	var sum int
	var ret int
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum >= k*threshold {
		ret += 1
	}
	for i := k; i < len(arr); i++ {
		sum -= arr[i-k]
		sum += arr[i]
		if sum >= k*threshold {
			ret += 1
		}
	}
	return ret
}
