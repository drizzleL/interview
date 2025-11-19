package main

import (
	"fmt"
	"sort"
)

type Router struct {
	dict     map[string]bool
	destDict map[int][]int
	limit    int
	cnt      int
	packets  [][]int
}

func RouterConstructor(memoryLimit int) Router {
	return Router{
		dict:     make(map[string]bool),
		destDict: make(map[int][]int),
		limit:    memoryLimit,
	}
}

func (this *Router) getKey(source int, destination int, timestamp int) string {
	return fmt.Sprintf("%d_%d_%d", source, destination, timestamp)
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := this.getKey(source, destination, timestamp)
	if this.dict[key] {
		return false
	}
	if this.cnt == this.limit {
		this.ForwardPacket()
	}
	this.dict[key] = true
	this.packets = append(this.packets, []int{source, destination, timestamp})
	this.cnt += 1
	if this.destDict[destination] == nil {
		this.destDict[destination] = []int{}
	}
	this.destDict[destination] = append(this.destDict[destination], timestamp)
	return true
}

func (this *Router) ForwardPacket() []int {
	if this.cnt == 0 {
		return nil
	}
	top := this.packets[0]
	this.packets = this.packets[1:]
	this.cnt -= 1
	delete(this.dict, this.getKey(top[0], top[1], top[2]))
	this.destDict[top[1]] = this.destDict[top[1]][1:]
	return top
}

func (this *Router) getCount(destination int, endTime int) int {
	if this.destDict[destination] == nil {
		return 0
	}
	return sort.Search(len(this.destDict[destination]), func(i int) bool {
		return this.destDict[destination][i] > endTime
	})
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	return this.getCount(destination, endTime) - this.getCount(destination, startTime-1)
}
