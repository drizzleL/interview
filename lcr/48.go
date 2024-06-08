package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	nums := []string{strconv.Itoa(root.Val)}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var next []*TreeNode
		var tmp []string
		for _, n := range nodes {
			if n.Left == nil {
				tmp = append(tmp, "null")
			} else {
				tmp = append(tmp, strconv.Itoa(n.Left.Val))
				next = append(next, n.Left)
			}
			if n.Right == nil {
				tmp = append(tmp, "null")
			} else {
				tmp = append(tmp, strconv.Itoa(n.Right.Val))
				next = append(next, n.Right)
			}
		}
		nodes = next
		if len(nodes) != 0 {
			nums = append(nums, tmp...)
		}
	}
	return strings.Join(nums, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data, ",")
	num, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: num}
	lastnodes := []*TreeNode{root}
	nextnodes := []*TreeNode{}
	var cnt int
	for i := 1; i < len(vals); i++ {
		val := vals[i]
		if val != "null" {
			num, _ := strconv.Atoi(val)
			n := &TreeNode{Val: num}
			nextnodes = append(nextnodes, n)
			lastnode := lastnodes[cnt/2]
			if cnt%2 == 0 {
				lastnode.Left = n
			} else {
				lastnode.Right = n
			}
		}
		cnt++
		if cnt == len(lastnodes)*2 {
			// check if nextnodes all empty
			lastnodes = nextnodes
			nextnodes = nil
			cnt = 0
		}
	}
	return root
}
