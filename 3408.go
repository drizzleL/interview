package main

import "container/heap"

type Task struct {
	TaskId   int
	Priority int
	Idx      int
	User     int
}

type Tasks []*Task

func (ts Tasks) Len() int {
	return len(ts)
}

func (ts Tasks) Less(i, j int) bool {
	if ts[i].Priority != ts[j].Priority {
		return ts[i].Priority > ts[j].Priority
	}
	return ts[i].TaskId > ts[j].TaskId
}

func (ts Tasks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
	ts[i].Idx, ts[j].Idx = ts[j].Idx, ts[i].Idx
}

func (ts *Tasks) Push(x any) {
	*ts = append(*ts, x.(*Task))
}

func (ts *Tasks) Pop() any {
	old := *ts
	n := len(old)
	x := old[n-1]
	*ts = old[0 : n-1]
	return x
}

type TaskManager struct {
	taskDict map[int]*Task
	h        *Tasks
}

func TaskConstructor(tasks [][]int) TaskManager {
	h := &Tasks{}
	taskDict := map[int]*Task{}
	for _, t := range tasks {
		user, task, priority := t[0], t[1], t[2]
		t := &Task{
			TaskId:   task,
			Idx:      h.Len(),
			User:     user,
			Priority: priority,
		}
		taskDict[task] = t
		*h = append(*h, t)
	}
	heap.Init(h)
	return TaskManager{
		taskDict: taskDict,
		h:        h,
	}
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	t := &Task{
		TaskId:   taskId,
		Idx:      this.h.Len(),
		User:     userId,
		Priority: priority,
	}
	this.taskDict[taskId] = t
	heap.Push(this.h, t)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	t := this.taskDict[taskId]
	t.Priority = newPriority
	heap.Fix(this.h, t.Idx)
}

func (this *TaskManager) Rmv(taskId int) {
	t := this.taskDict[taskId]
	heap.Remove(this.h, t.Idx)
}

func (this *TaskManager) ExecTop() int {
	if this.h.Len() == 0 {
		return -1
	}
	t := heap.Pop(this.h).(*Task)
	return t.User
}
