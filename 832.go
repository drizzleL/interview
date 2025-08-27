package main

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		for j, k := 0, len(image[0])-1; j <= k; j, k = j+1, k-1 {
			image[i][j], image[i][k] = image[i][k]^1, image[i][j]^1
		}
	}
	return image
}
