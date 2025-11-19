package main

func generateSchedule(n int) [][]int {
	if n <= 4 {
		return nil
	}
	res := make([][]int, 0, n*(n-1))

	if n%2 != 0 {
		for i := 0; i < 2*n; i += 2 {
			res = append(res, []int{mod(i, n), mod(i+1, n)})
		}
		for i := 0; i < 2*n; i += 2 {
			res = append(res, []int{mod(i+1, n), mod(i, n)})
		}
	} else {
		for i := 0; i < n; i += 2 {
			res = append(res, []int{i, i + 1})
		}
		for i := 0; i < n; i += 2 {
			res = append(res, []int{i + 1, i})
		}
		for i := 1; i < n; i += 2 {
			res = append(res, []int{i, mod(i+1, n)})
		}
		for i := 1; i < n; i += 2 {
			res = append(res, []int{mod(i+1, n), i})
		}
	}

	for diff := 2; diff < (n+1)/2; diff++ {
		start := res[len(res)-1][0] + 1
		for i := start; i < start+n; i++ {
			res = append(res, []int{mod(i, n), mod(i+diff, n)})
		}

		start = res[len(res)-1][1] - 1
		for i := start; i < start+n; i++ {
			res = append(res, []int{mod(i+diff, n), mod(i, n)})
		}
	}

	if n%2 == 0 {
		start := res[len(res)-1][0] - 1
		for i := start; i < start+n; i++ {
			res = append(res, []int{mod(i, n), mod(i+n/2, n)})
		}
	}

	return res
}

func mod(a, n int) int {
	return ((a % n) + n) % n
}
