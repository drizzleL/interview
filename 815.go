package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stopRoutes := map[int][]int{}
	for i, r := range routes {
		for _, s := range r {
			stopRoutes[s] = append(stopRoutes[s], i)
		}
	}
	stops := []int{source}
	seenStop := map[int]bool{
		source: true,
	}
	usedRoutes := make([]bool, len(routes))
	for i := 1; len(stops) != 0; i++ {
		var nextStops []int
		for _, s := range stops {
			for _, r := range stopRoutes[s] {
				if usedRoutes[r] {
					continue
				}
				usedRoutes[r] = true
				for _, next := range routes[r] {
					if next == target { // found it
						return i
					}
					if seenStop[next] {
						continue
					}
					seenStop[next] = true
					nextStops = append(nextStops, next)
				}
			}
		}
		stops = nextStops
	}
	return -1
}
