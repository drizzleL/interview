package main

func minOperations3(logs []string) int {
	var level int
	for _, l := range logs {
		switch l {
		case "../":
			if level != 0 {
				level -= 1
			}
		case "./":
		default:
			level += 1
		}
	}
	return level
}
