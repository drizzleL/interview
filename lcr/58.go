package main

import (
	"sort"
)

type MyCalendar struct {
	data [][2]int
}

func CalConstructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.data) == 0 {
		this.data = append(this.data, [2]int{start, end})
		return true
	}
	idx := sort.Search(len(this.data), func(i int) bool {
		return this.data[i][0] > start
	})
	if idx >= 1 && this.data[idx-1][1] > start {
		return false
	}
	if idx < len(this.data) && this.data[idx][0] < end {
		return false
	}
	tail := append([][2]int{{start, end}}, this.data[idx:]...)
	this.data = append(this.data[:idx], tail...)
	return true
}
