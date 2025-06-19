package main

type MinStack struct {
	vals []int
	mins []int
}

func MSConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.vals = this.vals[:len(this.vals)-1]
	if top == this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
