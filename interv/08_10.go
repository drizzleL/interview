package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	nodes := [][2]int{{sr, sc}}
	for len(nodes) != 0 {
		var next [][2]int
		for _, n := range nodes {
			image[n[0]][n[1]] = newColor
			for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				r, c := n[0]+dir[0], n[1]+dir[1]
				if r >= len(image) || r < 0 {
					continue
				}
				if c >= len(image[0]) || c < 0 {
					continue
				}
				if image[r][c] != oldColor {
					continue
				}
				next = append(next, [2]int{r, c})
			}
		}
		nodes = next
	}
	return image
}
