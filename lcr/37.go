package main

func asteroidCollision(asteroids []int) []int {
	var ret []int
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] >= 0 {
			ret = append(ret, asteroids[i])
			continue
		}
		for len(ret) != 0 && ret[len(ret)-1] > 0 && ret[len(ret)-1] < -asteroids[i] {
			ret = ret[:len(ret)-1]
		}
		if len(ret) != 0 && ret[len(ret)-1] == -asteroids[i] {
			ret = ret[:len(ret)-1]
			continue
		}
		if len(ret) != 0 && ret[len(ret)-1] > -asteroids[i] {
			continue
		}
		ret = append(ret, asteroids[i])
	}
	return ret
}
