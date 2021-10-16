package model

var todoStatus = map[string]int{
	TodoStatusTodo.String():    0,
	TodoStatusProcess.String(): 1,
	TodoStatusDone.String():    2,
}

func (t TodoStatus) Int() int {
	return todoStatus[t.String()]
}
