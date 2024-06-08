package main

type SortedStack struct {
	mins []int
	tmp  []int
}

func ConstructorStack() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	if len(this.mins) != 0 && val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
		return
	}
	if len(this.mins) != 0 && this.mins[len(this.mins)-1] < val {
		this.tmp = append(this.tmp, this.mins[len(this.mins)-1])
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.mins = append(this.mins, val)
	for len(this.tmp) != 0 {
		this.mins = append(this.mins, this.tmp[len(this.tmp)-1])
		this.tmp = this.tmp[:len(this.tmp)-1]
	}

}

func (this *SortedStack) Pop() {
	if len(this.mins) == 0 {
		return
	}
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *SortedStack) Peek() int {
	if len(this.mins) == 0 {
		return -1
	}
	return this.mins[len(this.mins)-1]
}
