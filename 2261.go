package main

func countDistinct(nums []int, k int, p int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	dict := map[int][][]int{}
	var ret int
	for i := 0; i < len(nums); i++ {
		cnt := 0
		base := 1
		key := 0
		var curr []int
		for j := i; j < len(nums); j++ {
			curr = append(curr, nums[j])
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			key += base * nums[j]
			key %= 1e9 + 7
			base *= maxVal
			base %= 1e9 + 7
			if len(dict[key]) == 0 {
				ret += 1
				tmp := make([]int, len(curr))
				copy(tmp, curr)
				dict[key] = append(dict[key], tmp)
				continue
			}
			var flag bool
			for _, v := range dict[key] {
				if len(v) != len(curr) {
					continue
				}
				for idx, val := range v {
					if curr[idx] == val {
						continue
					}
					break
				}
				flag = true
			}
			if !flag {
				ret += 1
				tmp := make([]int, len(curr))
				copy(tmp, curr)
				dict[key] = append(dict[key], tmp)
			}
		}
	}
	return ret
}
