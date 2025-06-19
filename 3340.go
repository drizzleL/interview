package main

func isBalanced(num string) bool {
	var evenSum, oddSum int
	for i, c := range num {
		switch i % 2 {
		case 0:
			evenSum += int(c - '0')
		case 1:
			oddSum += int(c - '0')
		}
	}
	return evenSum == oddSum
}
