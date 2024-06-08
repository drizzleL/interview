package main

func sumGame(num string) bool {
	var res float64
	for i := 0; i < len(num)/2; i++ {
		if num[i] == '?' {
			res += 4.5
		} else {
			res += float64(num[i] - '0')
		}
	}
	for i := len(num) / 2; i < len(num); i++ {
		if num[i] == '?' {
			res -= 4.5
		} else {
			res -= float64(num[i] - '0')
		}
	}
	return res != 0
}
