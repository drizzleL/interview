package main

func subarrayGCD(nums []int, k int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		m := nums[i]
		for j := i; j < len(nums); j++ {
			newm := gcd(m, nums[j])
			if newm < k {
				break
			}
			if newm == k {
				ret += 1
			}
			m = newm
		}
	}
	return ret
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
