package main

type BookMyShow struct {
	vals             []int
	segTree          SegTree
	n, m             int
	currRow, currCol int
}

func BookConstructor(n int, m int) BookMyShow {
	vals := make([]int, n)
	for i := range vals {
		vals[i] = m
	}
	sg := NewSegTree(vals)
	return BookMyShow{
		vals:    vals,
		segTree: sg,
		n:       n,
		m:       m,
	}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	idx := this.segTree.QueryMax(1, 0, len(this.vals)-1, k)
	if idx == -1 || idx > maxRow {
		return nil
	}
	ret := []int{idx, this.m - this.vals[idx]}
	this.vals[idx] -= k
	this.segTree.Update(1, 0, len(this.vals)-1, idx, this.vals[idx])
	return ret
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	sum := this.segTree.QuerySum(1, 0, len(this.vals)-1, maxRow)
	if sum < k {
		return false
	}
	for k > 0 {
		for this.vals[this.currRow] == 0 {
			this.currRow += 1
		}
		taken := min(k, this.vals[this.currRow])
		k -= taken
		this.vals[this.currRow] -= taken
		this.segTree.Update(1, 0, len(this.vals)-1, this.currRow, this.vals[this.currRow])
	}
	return true
}

type SegNode struct {
	Sum int
	Max int
}

type SegTree []*SegNode

func NewSegTree(vals []int) SegTree {
	segTree := make([]*SegNode, len(vals)*4)
	var build func(i int, l, r int)
	build = func(i int, l, r int) {
		segTree[i] = &SegNode{}
		if l == r {
			segTree[i].Sum = vals[l]
			segTree[i].Max = vals[l]
			return
		}
		mid := l + (r-l)/2
		build(i*2, l, mid)
		build(i*2+1, mid+1, r)
		segTree[i].Max = max(segTree[i*2].Max, segTree[i*2+1].Max)
		segTree[i].Sum = segTree[i*2].Sum + segTree[i*2+1].Sum
	}
	build(1, 0, len(vals)-1)
	return SegTree(segTree)
}

func (sg SegTree) Update(i, l, r int, idx int, val int) {
	if l == r {
		sg[i].Sum = val
		sg[i].Max = val
		return
	}
	mid := l + (r-l)/2
	if idx <= mid {
		sg.Update(i*2, l, mid, idx, val)
	} else {
		sg.Update(i*2+1, mid+1, r, idx, val)
	}
	sg[i].Max = max(sg[i*2].Max, sg[i*2+1].Max)
	sg[i].Sum = sg[i*2].Sum + sg[i*2+1].Sum
}

func (sg SegTree) QuerySum(i, l, r int, idx int) int {
	if l == r {
		return sg[i].Sum
	}
	mid := l + (r-l)/2
	if idx <= mid {
		return sg.QuerySum(i*2, l, mid, idx)
	}
	return sg[i*2].Sum + sg.QuerySum(i*2+1, mid+1, r, idx)
}

func (sg SegTree) QueryMax(i, l, r int, val int) int {
	if sg[i].Max < val {
		return -1
	}
	if l == r {
		return l
	}
	mid := l + (r-l)/2
	ret := sg.QueryMax(i*2, l, mid, val)
	if ret != -1 {
		return ret
	}
	return sg.QueryMax(i*2+1, mid+1, r, val)
}
