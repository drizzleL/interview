package main

import (
	"fmt"
)

func fractionAddition(expression string) string {
	var i int
	type item struct {
		numerator   int
		denominator int
	}
	getNum := func() *item {
		ret := &item{}
		var flag bool
		switch expression[i] {
		case '-':
			flag = true
			i++
		case '+':
			i++
		}
		for ; expression[i] != '/'; i++ {
			ret.numerator *= 10
			ret.numerator += int(expression[i] - '0')
		}
		if flag {
			ret.numerator *= -1
		}
		i++
		for ; i < len(expression) && expression[i] >= '0' && expression[i] <= '9'; i++ {
			ret.denominator *= 10
			ret.denominator += int(expression[i] - '0')
		}
		return ret
	}
	doMath := func(a, b *item) *item {
		if a.numerator == 0 {
			return b
		}
		if b.numerator == 0 {
			return b
		}
		denoGcd := gcd(a.denominator, b.denominator)
		denominator := a.denominator * b.denominator / denoGcd
		a.numerator *= denominator / a.denominator
		b.numerator *= denominator / b.denominator
		a.numerator += b.numerator
		a.denominator = denominator
		return a
	}
	data := &item{
		numerator:   0,
		denominator: 0,
	}
	for i < len(expression) {
		next := getNum()
		data = doMath(data, next)
	}
	if data.numerator == 0 {
		return "0/1"
	}
	lastGcd := gcd(abs(data.numerator), data.denominator)
	data.denominator /= lastGcd
	data.numerator /= lastGcd
	return fmt.Sprintf("%d/%d", data.numerator, data.denominator)
}
