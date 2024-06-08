package main

type LUPrefix struct {
	left  []int
	right []int
	flag  []bool
}

func ConstructorLU(n int) LUPrefix {
	return LUPrefix{
		flag:  make([]bool, n+1),
		left:  make([]int, n+1),
		right: make([]int, n+1),
	}
}

func (this *LUPrefix) Upload(video int) {
	this.flag[video] = true
	this.left[video] = video
	this.right[video] = video
	realleft, realright := video, video
	if video != 0 && this.flag[video-1] {
		realleft = this.left[video-1]
	}
	if video != len(this.flag)-1 && this.flag[video+1] {
		realright = this.right[video+1]
	}
	this.right[realleft] = realright
	this.left[realright] = realleft
}

func (this *LUPrefix) Longest() int {
	if !this.flag[1] {
		return 0
	}
	return this.right[1]
}
