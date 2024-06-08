package main

func doesValidArrayExist(derived []int) bool {
	var flag int
	for _, num := range derived {
		flag ^= num
	}
	return flag == 0
}
