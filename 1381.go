package main

type CustomStack struct {
	maxSize int
	data    []int
}

func CustomConstructor(maxSize int) CustomStack {
	return CustomStack{
		data:    make([]int, 0, maxSize),
		maxSize: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) == this.maxSize {
		return
	}
	this.data = append(this.data, x)
}

func (this *CustomStack) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	v := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return v
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.data); i++ {
		this.data[i] += val
	}
}
