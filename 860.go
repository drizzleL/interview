package main

func lemonadeChange(bills []int) bool {
	var cnt5, cnt10 int
	for _, b := range bills {
		switch b {
		case 5:
			cnt5 += 1
		case 10:
			if cnt5 == 0 {
				return false
			}
			cnt5 -= 1
			cnt10 += 1
		case 20:
			if cnt10 != 0 && cnt5 != 0 {
				cnt10 -= 1
				cnt5 -= 1
			} else if cnt5 >= 3 {
				cnt5 -= 3
			} else {
				return false
			}
		}
	}
	return true
}
