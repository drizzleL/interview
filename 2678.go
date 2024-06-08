package main

func countSeniors(details []string) int {
	var ret int
	for _, det := range details {
		a, b := det[11], det[12]
		if a > '6' || (a == '6' && b != '0') {
			ret += 1
		}
	}
	return ret
}
