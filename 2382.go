package main

func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	ret := make([]int64, len(nums))
	var val int64
	dict := map[int][3]int{}
	for i := len(removeQueries) - 1; i >= 0; i-- {
		idx := removeQueries[i]
		ret[i] = val
		prev, ok1 := dict[idx-1]
		next, ok2 := dict[idx+1]
		var v int64
		switch {
		case ok1 && ok2:
			delete(dict, prev[0])
			delete(dict, prev[1])
			delete(dict, next[0])
			delete(dict, next[1])
			prev[1] = next[1]
			prev[2] += nums[idx] + next[2]
			dict[prev[0]] = prev
			dict[prev[1]] = prev
			v = int64(prev[2])
		case ok1:
			delete(dict, prev[0])
			delete(dict, prev[1])
			prev[1] = idx
			prev[2] += nums[idx]
			dict[prev[0]] = prev
			dict[prev[1]] = prev
			v = int64(prev[2])
		case ok2:
			delete(dict, next[0])
			delete(dict, next[1])
			next[0] = idx
			next[2] += nums[idx]
			dict[next[0]] = next
			dict[next[1]] = next
			v = int64(next[2])
		default:
			dict[idx] = [3]int{idx, idx, nums[idx]}
			v = int64(nums[idx])
		}
		if v > val {
			val = v
		}
	}
	return ret
}
