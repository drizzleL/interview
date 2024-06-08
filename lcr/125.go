package main

type CQueue struct {
	in, out []int
}

func CQConstructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) != 0 {
		v := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return v
	}
	if len(this.in) == 0 {
		return -1
	}
	for i := len(this.in) - 1; i >= 0; i-- {
		this.out = append(this.out, this.in[i])
	}
	this.in = this.in[:0]
	return this.DeleteHead()
}
