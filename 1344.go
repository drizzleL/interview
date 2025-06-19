package main

func angleClock(hour int, minutes int) float64 {
	a := float64(minutes * 6)
	b := float64(hour)*30 + float64(minutes)*0.5
	if a > b {
		a, b = b, a
	}
	ret := b - a
	if ret > 180 {
		ret = 360 - ret
	}
	return ret
}
