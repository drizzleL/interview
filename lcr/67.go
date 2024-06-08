package main

type TrieNode struct {
	nodes [2]*TrieNode
}

func add(root *TrieNode, val int) {
	node := root
	for i := 31; i >= 0; i-- {
		flag := ((1 << i) & val) != 0
		if flag {
			if node.nodes[1] == nil {
				node.nodes[1] = &TrieNode{}
			}
			node = node.nodes[1]
		} else {
			if node.nodes[0] == nil {
				node.nodes[0] = &TrieNode{}
			}
			node = node.nodes[0]
		}
	}
}

func check(root *TrieNode, val int) int {
	node := root
	var ret int
	for i := 31; i >= 0; i-- {
		flag := ((1 << i) & val) != 0
		if flag {
			if node.nodes[0] != nil {
				node = node.nodes[0]
				ret += 1 << i
			} else {
				node = node.nodes[1]
			}
		} else {
			if node.nodes[1] != nil {
				node = node.nodes[1]
				ret += 1 << i
			} else {
				node = node.nodes[0]
			}
		}
	}
	return ret
}

func findMaximumXOR(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	root := &TrieNode{}
	var ret int
	add(root, nums[0])
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		ret = max(ret, check(root, num))
		add(root, num)
	}
	return ret
}
