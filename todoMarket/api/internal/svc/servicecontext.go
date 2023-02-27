package svc

import (
	"todo-market/todoMarket/api/internal/config"
	"todo-market/todoMarket/api/internal/model"
)

type ServiceContext struct {
	Config        config.Config
	TodoItemModel *model.TodoItemModel
	TodoListModel *model.TodoListModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := model.InitModel(c.DataSource)
	return &ServiceContext{
		Config:        c,
		TodoListModel: model.NewTodoListModel(db),
		TodoItemModel: model.NewTodoItemModel(db),
	}
}
