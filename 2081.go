package main

func kMirror(k int, n int) int64 {
	check := func(x int) bool {
		arr := []int{}
		for x != 0 {
			arr = append(arr, x%k)
			x /= k
		}
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			if arr[i] != arr[j] {
				return false
			}
		}
		return true
	}
	getRev := func(x int) int {
		var ret int
		for x != 0 {
			ret *= 10
			ret += x % 10
			x /= 10
		}
		return ret
	}
	var ret int
	var cnt int
	base := 1
	for i := base; cnt < n && i <= base*10-1; i++ {
		if check(i) {
			ret += i
			cnt += 1
		}
	}
	for cnt < n {
		for i := base; cnt < n && i <= base*10-1; i++ {
			v := i*base*10 + getRev(i)
			if check(v) {
				ret += v
				cnt += 1
			}
		}
		for i := base; cnt < n && i <= base*10-1; i++ {
			for j := 0; j <= 9 && cnt < n; j++ {
				v := i*base*100 + j*base*10 + getRev(i)
				if check(v) {
					ret += v
					cnt += 1
				}
			}
		}
		base *= 10
	}
	return int64(ret)
}
