package main

import (
	"strconv"
)

func daysBetweenDates(date1 string, date2 string) int {
	monthDays := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}

	getDay := func(year, month, day int) int {
		d := day
		for i := 1; i < month; i++ {
			d += monthDays[i]
		}
		leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)
		if month > 2 && leap {
			d += 1
		}
		return d
	}
	days := func(d string) int {
		year, _ := strconv.Atoi(d[:4])
		year2 := year - 1
		month, _ := strconv.Atoi(d[5:7])
		day, _ := strconv.Atoi(d[8:])
		ret := (year2 - 1970) * 365
		n := year2/4 - 1970/4
		n -= year2/100 - 1970/100
		n += year2/400 - 1970/400
		ret += n
		ret += getDay(year, month, day)
		return ret
	}
	a := days(date2)
	b := days(date1)
	if a < b {
		a, b = b, a
	}
	return a - b
}
