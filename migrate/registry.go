package main

import "github.com/engtiago/go-boilerplate/src/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
	}
}
