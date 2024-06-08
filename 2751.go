package main

import (
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	type Robot struct {
		Id     int
		Pos    int
		Health int
		Dir    byte
	}
	var robots []*Robot
	for i := range positions {
		robots = append(robots, &Robot{
			Id:     i,
			Pos:    positions[i],
			Health: healths[i],
			Dir:    directions[i],
		})
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].Pos < robots[j].Pos
	})
	var stack []*Robot
	for _, robot := range robots {
		if robot.Dir == 'R' {
			stack = append(stack, robot)
			continue
		}
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top.Health == robot.Health { // both crashed
				top.Health = 0
				robot.Health = 0
				break
			}
			if top.Health < robot.Health {
				top.Health = 0
				robot.Health -= 1
				continue
			}
			top.Health -= 1
			robot.Health = 0
			stack = append(stack, top)
			break
		}
	}
	var ret []int
	for _, robot := range robots {
		healths[robot.Id] = robot.Health
	}
	for _, h := range healths {
		if h == 0 {
			continue
		}
		ret = append(ret, h)
	}
	return ret
}
