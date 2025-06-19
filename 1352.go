package main

type ProductOfNumbers struct {
	arr []int
}

func ProductConstructor() ProductOfNumbers {
	return ProductOfNumbers{}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 || len(this.arr) == 0 || this.arr[len(this.arr)-1] == 0 {
		this.arr = this.arr[:0]
		this.arr = append(this.arr, num)
		return
	}
	this.arr = append(this.arr, num*this.arr[len(this.arr)-1])
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > len(this.arr) {
		return 0
	}
	last := this.arr[len(this.arr)-1]
	if k == len(this.arr) {
		return last
	}
	first := this.arr[len(this.arr)-k-1]
	return last / first
}
