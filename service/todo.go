package service

import (
	"fmt"
	"myapp/graph/model"
	"strconv"
)

var list []*model.Todo
var id = 0

func TodolistGet() []*model.Todo {
	return list
}

func NewTodo(input model.NewTodo) *model.Todo {
	id = id + 1

	statusInInteger := input.Status.Int()
	fmt.Println(statusInInteger)

	newTodo := &model.Todo{
		ID:     strconv.Itoa(id),
		Text:   input.Text,
		Status: input.Status,
	}
	list = append(list, newTodo)

	return newTodo
}
