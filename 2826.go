package main

func minimumOperations5(nums []int) int {
	var a, b, c int
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 1:
			b += 1
			c += 1
		case 2:
			b = min(b, a)
			a += 1
			c += 1
		case 3:
			c = min(c, a)
			c = min(c, b)
			a += 1
			b += 1
		}
	}
	return min(min(a, b), c)
}
