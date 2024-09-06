package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}

func dealWith(ff []string) [][]int {
	var ret [][]int
	for _, f := range ff {
		strs := strings.Split(f, " ")
		a, b := strs[0], strs[1]
		av, _ := strconv.Atoi(a)
		bv, _ := strconv.Atoi(b)
		if av < bv {
			av, bv = bv, av
		}
		ret = append(ret, []int{av, bv})
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i][0] == ret[j][0] {
			return ret[i][1] < ret[j][1]
		}
		return ret[i][0] < ret[j][0]
	})
	return ret

}

func cmpFile(f1, f2 string) {
	ff1, ff2 := readFile(f1), readFile(f2)
	l1, l2 := dealWith(ff1), dealWith(ff2)
	var i, j int
	for {
		if l1[i][0] != l2[j][0] || l1[i][1] != l2[j][1] {
			log.Println(l1[i], l2[j])
			j++
			continue
		}
		i++
		j++
	}
}

func main() {
	// cmpFile("result2.txt", "result3.txt")

	start := time.Now()
	defer func() {
		log.Printf("takes %v\n", time.Since(start))
	}()
	// p(modifiedGraphEdges(5, toIntInt("[[4,1,-1],[2,0,-1],[0,3,-1],[4,3,-1]]"), 0, 1, 5))
	// p(modifiedGraphEdges(5, toIntInt("[[1,4,1],[2,4,-1],[3,0,2],[0,4,-1],[1,3,10],[1,0,10]]"), 0, 2, 15))
	p(deleteAndEarn([]int{3, 4, 2}))

}

func toMap(data map[string]interface{}, obj interface{}, prefix string) {
	rv := reflect.ValueOf(obj)
	rv = reflect.Indirect(rv)
	rt := rv.Type()
	if rv.Kind() != reflect.Struct {
		data[prefix] = obj
		return
	}

	for i := 0; i < rt.NumField(); i++ {
		fieldName := rt.Field(i).Tag.Get("json")
		fieldName, _, _ = strings.Cut(fieldName, ",")
		if len(fieldName) == 0 {
			continue
		}
		field := rv.Field(i)
		if field.Kind() == reflect.Struct {
			toMap(data, field.Interface(), prefix+"."+fieldName)
			continue
		}
		if field.Kind() == reflect.Slice {
			continue
		}
		v := rv.Field(i).Interface()
		switch {
		case rv.Field(i).CanInt():
			v = rv.Field(i).Int()
		case rv.Field(i).CanFloat():
			v = rv.Field(i).Float()
		}
		data[prefix+"."+fieldName] = v
	}
}

var maskPhoneReg = regexp.MustCompile(`["\\](\d{3})(\d{4})(\d{4})["\\]`)

type A struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func toInts(s string) []int {
	var v []int
	json.Unmarshal([]byte(s), &v)
	return v
}

func toStrings(s string) []string {
	var v []string
	json.Unmarshal([]byte(s), &v)
	return v
}

func toIntInt(s string) [][]int {
	var v [][]int
	json.Unmarshal([]byte(s), &v)
	return v
}

func printMatrix(m [][]int) {
	fmt.Println()
	for i := range m {
		for _, v := range m[i] {
			fmt.Printf("%3d", v)
		}
		fmt.Println()
	}
	fmt.Println()
}

func ToListNode(s string) *ListNode {
	head := &ListNode{}
	node := head
	for _, n := range strings.Split(s, ",") {
		d, _ := strconv.Atoi(n)
		node.Next = &ListNode{Val: d}
		node = node.Next
	}

	return head.Next
}

func ToTreeNode(data string) *TreeNode {
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

// [10,5,15,1,8,null,7]

func printList(l *ListNode) {
	var v []int
	for l != nil {
		v = append(v, l.Val)
		l = l.Next
	}
	fmt.Println(v)
}
func genList(nums []int) *ListNode {
	pre := &ListNode{}
	node := pre
	for _, num := range nums {
		node.Next = &ListNode{
			Val: num,
		}
		node = node.Next
	}
	return pre.Next
}

var s = `order_cost
order_cnt_cancel
order_cnt_closed
order_cost_closed
order_cnt_refund
order_cost_refund
order_cnt_unpaid
order_cost_unpaid
order_cnt_lowprice
order_cnt_change
order_cost_change
order_cnt_samedp
order_cost_samedp
order_cnt_abnormal
order_cost_abnormal`

func p(k interface{}) {
	b, _ := json.Marshal(k)
	fmt.Println(string(b))
}
