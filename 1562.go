package main

import (
	"sort"
)

func findLatestStep(arr []int, m int) int {
	if len(arr) == m {
		return m
	}
	data := [][2]int{{1, len(arr)}}
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		idx := sort.Search(len(data), func(i int) bool {
			return data[i][1] >= num
		})
		if data[idx][0] == data[idx][1] {
			data = append(data[:idx], data[idx+1:]...)
			continue
		}
		if (num-data[idx][0] == m) || (data[idx][1]-num == m) {
			return i
		}
		if data[idx][0] == num {
			data[idx][0] = num + 1
			continue
		}
		if data[idx][1] == num {
			data[idx][1] = num - 1
			continue
		}
		data = append(data, [2]int{})
		copy(data[idx+1:], data[idx:])
		v := data[idx]
		data[idx] = [2]int{v[0], num - 1}
		data[idx+1] = [2]int{num + 1, v[1]}
	}
	return -1
}
