package main

type MovingAverage struct {
	data []int
	size int
	sum  float64
	idx  int
}

/** Initialize your data structure here. */
func MovingConstructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.data) != this.size { // not full
		this.data = append(this.data, val)
		this.sum += float64(val)
		return this.sum / float64(len(this.data))
	}
	this.sum += float64(val - this.data[this.idx])
	this.data[this.idx] = val
	this.idx = (this.idx + 1) % this.size
	return this.sum / float64(this.size)
}
