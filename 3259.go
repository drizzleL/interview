package main

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	n := len(energyDrinkA)
	dpA, dpB := make([]int, n), make([]int, n)
	dpA[0] = energyDrinkA[0]
	dpB[0] = energyDrinkB[0]
	for i := 1; i < n; i++ {
		dpA[i] = dpA[i-1] + energyDrinkA[i]
		dpB[i] = dpB[i-1] + energyDrinkB[i]
		if i >= 2 {
			dpA[i] = max(dpA[i], dpB[i-2]+energyDrinkA[i])
			dpB[i] = max(dpB[i], dpA[i-2]+energyDrinkB[i])
		}
	}
	return int64(max(dpA[n-1], dpB[n-1]))
}
