package main

func diffWaysToCompute(expression string) []int {
	type item struct {
		t   int
		num int
	}
	isDigit := func(x byte) bool {
		return x >= '0' && x <= '9'
	}
	var items []item
	for i := 0; i < len(expression); i++ {
		var num int
		for i < len(expression) && isDigit(expression[i]) {
			num *= 10
			num += int(expression[i] - '0')
			i++
		}
		items = append(items, item{
			num: num,
		})
		if i < len(expression) {
			var t int
			switch expression[i] {
			case '+':
				t = 1
			case '-':
				t = 2
			case '*':
				t = 3
			}
			items = append(items, item{
				t: t,
			})
		}
	}
	dp := make([][][]int, len(items))
	for i := range dp {
		dp[i] = make([][]int, len(items))
	}
	var helper func(i, j int) []int
	helper = func(i, j int) (ret []int) {
		if len(dp[i][j]) != 0 {
			return dp[i][j]
		}
		defer func() {
			dp[i][j] = ret
		}()
		if i == j {
			return []int{items[i].num}
		}
		for k := i + 1; k < j; k += 2 {
			pre, suff := helper(i, k-1), helper(k+1, j)
			for _, v1 := range pre {
				for _, v2 := range suff {
					var v int
					switch items[k].t {
					case 1:
						v = v1 + v2
					case 2:
						v = v1 - v2
					case 3:
						v = v1 * v2
					}
					ret = append(ret, v)
				}
			}
		}
		return
	}
	return helper(0, len(items)-1)
}
