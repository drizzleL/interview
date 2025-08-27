package main

import "math"

func minMoves5(sx int, sy int, tx int, ty int) int {
	var ret int
	for tx != ty && tx > sx && ty > sy {
		if ty > tx {
			if tx*2 >= ty {
				ty -= tx
				ret += 1
			} else {
				if ty%2 != 0 {
					return -1
				}
				ty /= 2
				ret += 1
			}
		} else {
			if ty*2 >= tx {
				tx -= ty
				ret += 1
			} else {
				if tx%2 != 0 {
					return -1
				}
				tx /= 2
				ret += 1
			}
		}
	}
	helper := func(a, b int, v int) int {
		var ret int
		for b < v {
			b += max(a, b)
			ret += 1
		}
		if b == v {
			return ret
		}
		return -1
	}
	if sx == tx {
		v := helper(sx, sy, ty)
		if v == -1 {
			return -1
		}
		return ret + v
	}
	if sy == ty {
		v := helper(sy, sx, tx)
		if v == -1 {
			return -1
		}
		return ret + v
	}
	if tx == ty {
		if min(sx, sy) != 0 {
			return -1
		}
		if max(sx, sy) == 0 {
			return -1
		}
		if tx%max(sx, sy) != 0 {
			return -1
		}
		return ret + int(math.Log2(float64(tx/max(sx, sy)))) + 1
	}
	return -1
}
