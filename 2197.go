package main

func replaceNonCoprimes(nums []int) []int {
	var ret []int
	check := func(x int) int {
		if len(ret) == 0 {
			return 1
		}
		return gcd(ret[len(ret)-1], x)
	}
	for _, num := range nums {
		for {
			g := check(num)
			if g == 1 {
				break
			}
			last := ret[len(ret)-1]
			ret = ret[:len(ret)-1]
			num = last * num / g
		}
		ret = append(ret, num)
	}
	return ret
}
