package main

func compareVersion(version1 string, version2 string) int {
	var i, j int
	getNum := func(s string, idx int) (int, int) {
		var num int
		for idx < len(s) && s[idx] != '.' {
			num *= 10
			num += int(s[idx] - '0')
			idx += 1
		}
		return num, idx + 1
	}
	for i < len(version1) || j < len(version2) {
		var num1, num2 int
		num1, i = getNum(version1, i)
		num2, j = getNum(version2, j)
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}
