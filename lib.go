package main

type Bitmap struct {
	arr    []byte
	maxVal int
}

func NewBitmap(n int) *Bitmap {
	return &Bitmap{
		arr:    make([]byte, n/8+1),
		maxVal: n,
	}
}

func (b *Bitmap) mark(v int) {
	idx := v / 8
	offset := v % 8
	b.arr[idx] |= 1 << offset
}

func (b *Bitmap) check(v int) bool {
	idx := v / 8
	offset := v % 8
	return b.arr[idx]&(1<<offset) != 0
}

func NewBiTree(size int) biTree {
	return biTree(make([]int, size+1))
}

type biTree []int

func (bit biTree) Query(idx int) int {
	idx += 1
	var ret int
	for idx > 0 {
		ret += bit[idx]
		idx -= lowbit(idx)
	}
	return ret
}

func (bit biTree) Update(idx int, val int) {
	idx += 1
	for idx < len(bit) {
		bit[idx] += val
		idx += lowbit(idx)
	}
}

func lowbit(x int) int {
	return x & (-x)
}
