package main

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	dict := map[int]string{
		0:  "",
		1:  "One",
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
		20: "Twenty",
		30: "Thirty",
		40: "Forty",
		50: "Fifty",
		60: "Sixty",
		70: "Seventy",
		80: "Eighty",
		90: "Ninety",
	}
	add := func(a, b []byte) []byte {
		if len(a) != 0 {
			a = append(a, ' ')
		}
		return append(a, b...)
	}
	helper := func(x int) []byte {
		if x == 0 {
			return nil
		}
		var b []byte
		if x >= 100 {
			b = []byte(dict[x/100] + " Hundred")
			x %= 100
		}
		if x == 0 {
			return b
		}
		if dict[x] != "" {
			return add(b, []byte(dict[x]))
		}
		b = add(b, []byte(dict[x/10*10]))
		x %= 10
		b = add(b, []byte(dict[x]))
		return b
	}
	unit := []string{"", "Thousand", "Million", "Billion"}
	var ret []byte
	for i := 0; num != 0; i++ {
		b := helper(num % 1000)
		if len(b) != 0 {
			if i != 0 {
				b = append(b, ' ')
				b = append(b, unit[i]...)
				if len(ret) != 0 {
					b = append(b, ' ')
				}
			}
			ret = append(b, ret...)
		}
		num /= 1000
	}
	return string(ret)
}
