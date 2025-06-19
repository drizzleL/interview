package main

func canChange(start string, target string) bool {
	var j int
	findL := func(i int) bool {
		for ; j < i; j++ {
			if start[j] != '_' {
				return false
			}
		}
		for ; j < len(start); j++ {
			switch start[j] {
			case 'R':
				return false
			case 'L':
				j += 1
				return true
			}
		}
		return false
	}
	findR := func(i int) bool {
		for ; j <= i; j++ {
			switch start[j] {
			case 'R':
				j += 1
				return true
			case 'L':
				return false
			}
		}
		return false
	}
	for i := 0; i < len(start); i++ {
		switch target[i] {
		case 'L':
			if !findL(i) {
				return false
			}
		case 'R':
			if !findR(i) {
				return false
			}
		}
	}
	for ; j < len(start) && start[j] == '_'; j++ {
	}
	return j == len(start)
}
