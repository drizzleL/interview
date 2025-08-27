package main

var chByte = []byte("a")

func kthCharacter(k int) byte {
	for len(chByte) < k {
		size := len(chByte)
		for i := 0; i < size; i++ {
			c := chByte[i]
			if c == 'z' {
				c = 'a'
			} else {
				c += 1
			}
			chByte = append(chByte, c)
		}
	}
	return chByte[k-1]
}
