package main

func sumOfThree(num int64) []int64 {
	if num%3 != 0 {
		return nil
	}
	v := num / 3
	return []int64{v - 1, v, v + 1}
}
