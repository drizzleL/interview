package main

func calculateScore2(instructions []string, values []int) int64 {
	var ret int
	seen := make([]bool, len(instructions))
	for i := 0; i > 0 && i < len(instructions) && !seen[i]; {
		switch instructions[i] {
		case "add":
			ret += values[i]
			i += 1
		case "jump":
			i += values[i]
		}
	}
	return int64(ret)
}
