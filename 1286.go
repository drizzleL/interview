package main

type CombinationIterator struct {
	data    string
	pos     []int
	hasNext bool
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	pos := make([]int, combinationLength)
	for i := 0; i < combinationLength; i++ {
		pos[i] = i
	}
	return CombinationIterator{
		data:    characters,
		pos:     pos,
		hasNext: true,
	}
}

func (this *CombinationIterator) Next() string {
	var ret []byte
	for i := 0; i < len(this.pos); i++ {
		ret = append(ret, this.data[this.pos[i]])
	}
	if this.pos[0] == len(this.data)-len(this.pos) {
		this.hasNext = false
	} else {
		idx := len(this.pos) - 1
		dataIdx := len(this.data) - 1
		for this.pos[idx] == dataIdx {
			idx -= 1
			dataIdx -= 1
		}
		val := this.pos[idx] + 1
		for i := idx; i < len(this.pos); i++ {
			this.pos[i] = val
			val += 1
		}
	}
	return string(ret)
}

func (this *CombinationIterator) HasNext() bool {
	return this.hasNext
}
