package main

import (
	"math/big"
	"sort"
)

func maxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	var uniqs []int
	for i, r := range rewardValues {
		if i == 0 || r != rewardValues[i-1] {
			uniqs = append(uniqs, r)
		}
	}
	dp := big.NewInt(0)
	dp.SetBit(dp, 0, 1)
	for _, r := range uniqs {
		mask := big.NewInt(0)
		mask.Lsh(big.NewInt(1), uint(r)).Sub(mask, big.NewInt(1))
		toAdd := big.NewInt(0).And(dp, mask)
		toAdd.Lsh(toAdd, uint(r))
		dp.Or(dp, toAdd)
	}
	return dp.BitLen() - 1
}
