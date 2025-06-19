package main

func maximumBinaryString(binary string) string {
	ret := make([]byte, len(binary))
	for i := 0; i < len(binary); i++ {
		ret[i] = '1'
	}
	var zeros, ones int
	for i := 0; i < len(binary); i++ {
		switch binary[i] {
		case '0':
			zeros += 1
		case '1':
			if zeros == 0 {
				ones += 1
			}
		}
	}
	if ones < len(binary) {
		ret[ones+zeros-1] = '0'
	}
	return string(ret)
}
