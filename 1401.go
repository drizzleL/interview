package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	x, y := xCenter, yCenter
	if x1 > xCenter {
		x = x1
	} else if x2 < xCenter {
		x = x2
	}
	if y1 > yCenter {
		y = y1
	} else if y2 < yCenter {
		y = y2
	}
	distX, distY := xCenter-x, yCenter-y
	return distX*distX+distY*distY <= radius*radius
}
