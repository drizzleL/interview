package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i, j}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
