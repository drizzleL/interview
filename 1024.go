package main

import (
	"math"
	"sort"
)

func videoStitching(clips [][]int, time int) int {
	if len(clips) == 0 {
		return -1
	}
	s := make([]int, time+1)
	for i := range s {
		s[i] = math.MaxInt32
	}
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	if clips[0][0] != 0 {
		return -1
	}
	for i := 0; i <= min(time, clips[0][1]); i++ {
		s[i] = 1
	}
	for i := 1; i < len(clips); i++ {
		clip := clips[i]
		if clip[0] > time {
			break
		}
		v := s[clip[0]]
		if v == math.MaxInt32 {
			return -1
		}
		for j := clip[0]; j <= min(time, clip[1]); j++ {
			s[j] = min(s[j], v+1)
		}
	}
	if s[time] == math.MaxInt32 {
		return -1
	}
	return s[time]
}
