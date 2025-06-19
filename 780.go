package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx > sx && ty > sy {
		if tx > ty {
			tx %= ty
		} else {
			ty %= tx
		}
	}
	return (tx == sx && ty >= sy && (ty-sy)%sx == 0) || (ty == sy && tx >= sx && (tx-sx)%sy == 0)
}
