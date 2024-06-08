package main

import (
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	vals := strings.Split(equation, "=")
	l, r := vals[0], vals[1]
	parse := func(s string) (num int, xnum int) {
		for i := 0; i < len(s); {
			j := i + 1
			for ; j < len(s) && s[j] != '+' && s[j] != '-'; j++ {
			}
			if s[j-1] == 'x' {
				var val int
				if s[i:j-1] == "-" {
					val = -1
				} else if s[i:j-1] == "" || s[i:j-1] == "+" {
					val = 1
				} else {
					val, _ = strconv.Atoi(s[i : j-1])
				}
				xnum += val
			} else {
				val, _ := strconv.Atoi(s[i:j])
				num += val
			}
			i = j
		}
		return
	}
	lnum, lxnum := parse(l)
	rnum, rxnum := parse(r)
	xnum := lxnum - rxnum
	num := rnum - lnum
	if xnum == 0 && num == 0 {
		return "Infinite solutions"
	}
	if xnum == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(num/xnum)
}
