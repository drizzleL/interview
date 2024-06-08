package main

func brokenCalc(startValue int, target int) int {
	if startValue > target {
		return startValue - target
	}
	var ret int
	for target != startValue {
		if target%2 == 1 { //  odd just add
			target += 1
			ret++
			continue
		}
		if target/2 >= startValue {
			target /= 2
			ret++
			continue
		}
		v1 := startValue*2 - target + 1
		v2 := startValue - target/2 + 1
		ret += min(v1, v2)
		break
	}
	return ret
}
