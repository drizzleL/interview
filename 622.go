package main

type MyCircularQueue struct {
	read, insert int
	data         []int
}

func QConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		read:   0,
		insert: 1,
		data:   make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.insert] = value
	this.insert += 1
	this.insert %= len(this.data)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.read += 1
	this.read %= len(this.data)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	idx := (this.read + 1) % len(this.data)
	return this.data[idx]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	idx := (this.insert - 1 + len(this.data)) % len(this.data)
	return this.data[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return ((this.read + 1) % len(this.data)) == this.insert
}

func (this *MyCircularQueue) IsFull() bool {
	return this.read == this.insert
}
