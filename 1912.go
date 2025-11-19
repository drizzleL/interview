package main

import "github.com/emirpasic/gods/trees/redblacktree"

type MovieRentingSystem struct {
	unrented map[int]*redblacktree.Tree
	rented   *redblacktree.Tree
	prices   map[[2]int]int
}

func MovieConstructor(n int, entries [][]int) MovieRentingSystem {
	comparator := func(a, b interface{}) int {
		va, vb := a.([]int), b.([]int)
		if va[2] != vb[2] {
			return va[2] - vb[2]
		}
		if va[0] != vb[0] {
			return va[0] - vb[0]
		}
		return va[1] - vb[1]
	}
	dict := map[int]*redblacktree.Tree{}
	prices := make(map[[2]int]int)
	for _, ent := range entries {
		if _, ok := dict[ent[1]]; !ok {
			dict[ent[1]] = redblacktree.NewWith(comparator)
		}
		dict[ent[1]].Put(ent, nil)
		prices[[2]int{ent[0], ent[1]}] = ent[2]
	}
	return MovieRentingSystem{
		rented:   redblacktree.NewWith(comparator),
		prices:   prices,
		unrented: dict,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	tr := this.unrented[movie]
	nodes := this.top5(tr)
	var ret []int
	for _, node := range nodes {
		ret = append(ret, node[0])
	}
	return ret
}

func (this *MovieRentingSystem) top5(tr *redblacktree.Tree) [][]int {
	if tr == nil {
		return nil
	}
	var ret [][]int
	node := tr.Root
	var q []*redblacktree.Node
	for len(ret) < 5 && (node != nil || len(q) != 0) {
		for node != nil {
			q = append(q, node)
			node = node.Left
		}
		node = q[len(q)-1].Right
		ret = append(ret, q[len(q)-1].Key.([]int))
		q = q[:len(q)-1]
	}
	return ret
}

func (this *MovieRentingSystem) getPrice(shop int, movie int) int {
	return this.prices[[2]int{shop, movie}]
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	price := this.getPrice(shop, movie)
	node := []int{shop, movie, price}
	tr := this.unrented[movie]
	tr.Remove(node)
	this.rented.Put(node, nil)
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	price := this.getPrice(shop, movie)
	node := []int{shop, movie, price}
	this.rented.Remove(node)
	tr := this.unrented[movie]
	tr.Put(node, nil)
}

func (this *MovieRentingSystem) Report() [][]int {
	nodes := this.top5(this.rented)
	var ret [][]int
	for _, node := range nodes {
		ret = append(ret, []int{node[0], node[1]})
	}
	return ret
}
