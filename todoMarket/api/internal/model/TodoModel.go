package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type TodoItemModel struct {
	Conn *gorm.DB
}

// notice 以分数代替

type TodoItem struct {
	ID         int64  `json:"ID" gorm:"column:TodoItemId"`
	TodoListId int64  `json:"todeListId"`
	Name       string `json:"name"`
	Type       string `json:"type"`                             // 类型 每日/每次/每周
	StartTime  int64  `json:"startTime"`                        // 开始时间
	EndTime    int64  `json:"endTime"`                          // 结束时间
	GetScore   int64  `json:"GetScore"`                         // 获得的分值
	TotalScore int64  `json:"score"`                            // 总分值
	Status     bool   `json:"status"`                           // 是否w
	CreateAt   int64  `json:"createAt" gorm:"column:createAt"`  // 秒级时间戳
	UpdateAt   int64  `json:"updateAt"  gorm:"column:updateAt"` // 秒级时间戳
}

func (c *TodoItem) TableName() string {
	return "TodoItem"
}

func NewTodoItemModel(db *gorm.DB) *TodoItemModel {
	return &TodoItemModel{Conn: db.Model(TodoItem{})}
}

func (c *TodoItemModel) Insert(TodoItem *TodoItem) error {
	TodoItem.CreateAt = time.Now().Unix()
	TodoItem.UpdateAt = time.Now().Unix()
	return c.Conn.Create(TodoItem).Error
}
func (c *TodoItemModel) Update(TodoItem *TodoItem, ID int64) error {
	TodoItem.ID = ID
	TodoItem.UpdateAt = time.Now().Unix()
	return c.Conn.Update(TodoItem).Error
}

func (c *TodoItemModel) FindById(ID int64) (*TodoItem, error) {
	TodoItem := &TodoItem{ID: ID}
	return TodoItem, c.Conn.First(TodoItem).Error
}
