package main

type StreamRank struct {
	data [500001]int
}

func Constructor() StreamRank {
	return StreamRank{}

}

func (this *StreamRank) Track(x int) {
	x += 1
	for x <= len(this.data) {
		this.data[x] += 1
		x += lowbit(x)
	}
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	x += 1
	var ret int
	for x != 0 {
		ret += this.data[x]
		x -= lowbit(x)
	}
	return ret
}

func lowbit(x int) int {
	return x & -x
}
