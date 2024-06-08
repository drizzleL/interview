package main

import "strconv"

func printBin(num float64) string {
	str := "0."
	for num != 0 && len(str) <= 32 {
		digit := int(num * 2)
		num -= float64(digit)
		str += strconv.Itoa(digit)
	}
	if len(str) <= 32 {
		return str
	}
	return "ERROR"
}
