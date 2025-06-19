package main

func beautifulSplits(nums []int) int {
	getZ := func(nums []int) []int {
		z := make([]int, len(nums))
		for i, l, r := 1, 0, 0; i < len(nums); i++ {
			if i <= r && z[i-l] < r-i+1 {
				z[i] = z[i-l]
			} else {
				z[i] = max(0, r-i+1)
				for i+z[i] < len(nums) && nums[z[i]] == nums[i+z[i]] {
					z[i] += 1
				}
			}
			if i+z[i]-1 > r {
				l = i
				r = i + z[i] - 1
			}
		}
		return z
	}
	seen := make([][]bool, len(nums))
	for i := range seen {
		seen[i] = make([]bool, len(nums))
	}
	var ret int
	z1 := getZ(nums)
	for i := 1; i < len(z1)-1; i++ {
		if z1[i] >= i {
			for j := i * 2; j < len(z1); j++ {
				seen[i][j] = true
			}
		}
		z2 := getZ(nums[i:])
		for j := i + 1; j < len(nums); j++ {
			if z2[j-i] >= j-i {
				seen[i][j] = true
			}
		}
	}
	for i := range seen {
		for j := range seen[i] {
			if seen[i][j] {
				ret += 1
			}
		}
	}
	return ret
}
