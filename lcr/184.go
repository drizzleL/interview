package main

type Checkout struct {
	items []int
	maxs  []int
}

func CheckoutConstructor() Checkout {
	return Checkout{}
}

func (this *Checkout) Get_max() int {
	if len(this.maxs) == 0 {
		return -1
	}
	return this.maxs[0]
}

func (this *Checkout) Add(value int) {
	this.items = append(this.items, value)
	for len(this.maxs) != 0 && this.maxs[len(this.maxs)-1] < value {
		this.maxs = this.maxs[:len(this.maxs)-1]
	}
	this.maxs = append(this.maxs, value)
}

func (this *Checkout) Remove() int {
	if len(this.items) == 0 {
		return -1
	}
	ret := this.items[0]
	this.items = this.items[1:]
	if ret == this.maxs[0] {
		this.maxs = this.maxs[1:]
	}
	return ret
}
