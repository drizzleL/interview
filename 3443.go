package main

func maxDistance3(s string, k int) int {
	var ret int
	for _, dir := range []string{"NE", "NW", "SE", "SW"} {
		kk := k
		var dist int
		for _, c := range s {
			var flag bool
			for _, c2 := range dir {
				if c == c2 {
					flag = true
					break
				}
			}
			if flag {
				dist += 1
			} else if kk != 0 {
				dist += 1
				kk -= 1
			} else {
				dist -= 1
			}
			ret = max(ret, dist)
		}
	}
	return ret
}
