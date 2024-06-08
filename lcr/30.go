package main

import "math/rand"

type RandomizedSet struct {
	data []int
	dict map[int]int
}

/** Initialize your data structure here. */
func RandomConstructor() RandomizedSet {
	return RandomizedSet{
		dict: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.dict[val]
	if ok {
		return false
	}
	this.dict[val] = len(this.data)
	this.data = append(this.data, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.dict[val]
	if !ok {
		return false
	}
	delete(this.dict, val)
	if idx+1 == len(this.data) { // last one ignore
		this.data = this.data[:len(this.data)-1]
		return true
	}
	top := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	this.dict[top] = idx
	this.data[idx] = top
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[rand.Intn(len(this.data))]
}
