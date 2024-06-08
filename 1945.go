package main

func getLucky(s string, k int) int {
	var sum int
	for _, c := range s {
		tmp := int(c-'a') + 1
		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}
	}
	for i := 1; i < k; i++ {
		var newsum int
		for sum != 0 {
			newsum += sum % 10
			sum /= 10
		}
		sum = newsum
	}
	return sum
}
