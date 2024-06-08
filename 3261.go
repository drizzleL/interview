package main

func countKConstraintSubstrings2(s string, k int, queries [][]int) []int64 {
	var cnt [2]int
	var l, r int
	presum := make([]int, 1, len(s)+1)
	for r := 0; r < len(s); r++ {
		cnt[s[r]-'0'] += 1
		for cnt[0] > k && cnt[1] > k {
			cnt[s[l]-'0'] -= 1
			l++
		}
		presum = append(presum, presum[len(presum)-1]+r-l+1)
	}
	cnt = [2]int{}
	right := make([]int, len(s))
	r = len(s) - 1
	for l := len(s) - 1; l >= 0; l-- {
		cnt[s[l]-'0'] += 1
		for cnt[0] > k && cnt[1] > k {
			cnt[s[r]-'0'] -= 1
			r--
		}
		right[l] = r
	}
	ret := make([]int64, len(queries))
	for i, q := range queries {
		idx := min(q[1], right[q[0]])
		length := idx - q[0] + 1
		tmp := (length+1)*length/2 + presum[q[1]+1] - presum[idx+1]
		ret[i] = int64(tmp)
	}
	return ret
}
