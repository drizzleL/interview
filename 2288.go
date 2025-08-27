package main

import (
	"fmt"
	"strings"
)

func discountPrices(sentence string, discount int) string {
	helper := func(word string) string {
		if len(word) <= 1 {
			return word
		}
		if word[0] != '$' {
			return word
		}
		var m int
		for i := 1; i < len(word); i++ {
			if word[i] < '0' || word[i] > '9' {
				return word
			}
			m = m*10 + int(word[i]-'0')
		}
		fm := float64(m) * (1 - float64(discount)/100)
		return fmt.Sprintf("$%.2f", fm)
	}
	words := strings.Fields(sentence)
	var b []byte
	for i := 0; i < len(words); i++ {
		m := helper(words[i])
		b = append(b, m...)
		if i == len(words)-1 {
			break
		}
		b = append(b, ' ')
	}
	return string(b)
}
