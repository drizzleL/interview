package main

// func mergeStones(stones []int, k int) int {
// 	sums := make([]int, len(stones)+1)
// 	for i, v := range stones {
// 		sums[i+1] += sums[i] + v
// 	}
// 	rangeSum := func(i, j int) int {
// 		return sums[j+1] - sums[i]
// 	}
// 	dict := map[string]int{}
// 	getKey := func(i, j int, m int) string {
// 		return fmt.Sprintf("%d_%d_%d", i, j, m)
// 	}
// 	var helper func(i, j int, m int) int
// 	helper = func(i, j int, m int) (ret int) {
// 		key := getKey(i, j, m)
// 		if v, ok := dict[key]; ok {
// 			return v
// 		}
// 		defer func() {
// 			dict[key] = ret
// 		}()
// 		if (j-i+1-m)%(k-1) != 0 {
// 			return math.MaxInt32
// 		}
// 		if i == j {
// 			if m == 1 {
// 				return 0
// 			} else {
// 				return math.MaxInt32
// 			}
// 		}
// 		if m == 1 {
// 			return helper(i, j, k) + rangeSum(i, j)
// 		}
// 		ret = math.MaxInt32
// 		for mid := i; mid < j; mid += k - 1 {
// 			ret = min(ret, helper(i, mid, 1)+helper(mid+1, j, m-1))
// 		}
// 		return ret
// 	}
// 	return helper(0, len(stones), 1)
// }
