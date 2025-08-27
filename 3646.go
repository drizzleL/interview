package main

import (
	"sort"
	"strconv"
)

var specialPali [17][]int64

func init() {
	s := [17][][]int{}
	var helper func(i int, curr []int, sum int)
	helper = func(i int, curr []int, sum int) {
		if i > 8 {
			return
		}
		helper(i+2, curr, sum)
		curr = append(curr, i)
		sum += i
		if sum >= 17 {
			return
		}
		s[sum] = append(s[sum], append([]int(nil), curr...))
		helper(i+2, curr, sum)
	}
	helper(2, nil, 0)
	for pos := 1; pos < 17; pos += 2 {
		for i := 1; i <= min(pos, 9); i += 2 {
			if i == pos {
				s[pos] = append(s[pos], []int{i})
				continue
			}
			for _, comb := range s[pos-i] {
				s[pos] = append(s[pos], append([]int{i}, comb...))
			}
		}
	}
	var f func(nums []int, seen []bool, cnt int, b []byte, ret *[]int64)
	f = func(nums []int, seen []bool, cnt int, b []byte, ret *[]int64) {
		if cnt == len(nums) {
			num, _ := strconv.Atoi(string(b))
			*ret = append(*ret, int64(num))
			return
		}
		for i := range nums {
			if seen[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !seen[i-1] {
				continue
			}
			seen[i] = true
			b[cnt] = '0' + byte(nums[i])
			b[len(b)-1-cnt] = b[cnt]
			f(nums, seen, cnt+1, b, ret)
			seen[i] = false
		}
	}
	helper2 := func(sum int, cand []int) {
		b := make([]byte, sum)
		var nums []int
		if sum%2 == 1 {
			b[sum/2] = '0' + byte(cand[0])
		}
		for i, v := range cand {
			limit := v
			if sum%2 == 1 && i == 0 {
				limit--
			}
			limit /= 2
			for j := 0; j < limit; j++ {
				nums = append(nums, v)
			}
		}
		f(nums, make([]bool, len(nums)), 0, b, &specialPali[sum])
	}
	for i := 1; i < 17; i += 1 {
		for _, v := range s[i] {
			helper2(i, v)
		}
		sort.Slice(specialPali[i], func(i2, j2 int) bool {
			return specialPali[i][i2] < specialPali[i][j2]
		})
	}

}

func specialPalindrome(n int64) int64 {
	l := len(strconv.Itoa(int(n)))
	for _, num := range specialPali[l] {
		if num > n {
			return num
		}
	}
	return specialPali[l+1][0]
}
