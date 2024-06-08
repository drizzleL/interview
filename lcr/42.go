package main

type RecentCounter struct {
	data []int
}

func RecentConstructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.data) != 0 && t-3000 > this.data[0] {
		this.data = this.data[1:]
	}
	this.data = append(this.data, t)
	return len(this.data)
}
