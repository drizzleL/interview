package main

func kthCharacter2(k int64, operations []int) byte {
	var diff int
	for i := len(operations) - 1; i >= 0; i-- {
		if (k>>i)&1 != 0 {
			diff += operations[i]
		}
	}
	return 'a' + byte(diff%26)
}

func kthCharacter3(k int64, operations []int) byte {
	var diff byte
	for i := 0; i < len(operations); i, k = i+1, k/2 {
		if k&1 == 0 {
			continue
		}
		if operations[i] == 1 {
			diff += 1
			diff %= 26
		}
	}
	return 'a' + diff
}
