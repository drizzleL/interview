package main

func canAliceWin(n int) bool {
	var flag bool
	for k := 10; n >= k; k-- {
		n -= k
		flag = !flag
	}
	return flag
}
