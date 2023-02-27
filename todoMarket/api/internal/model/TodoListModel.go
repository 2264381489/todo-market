package model

import (
	"time"
	"todo-market/todoMarket/api/internal/constant"

	"github.com/jinzhu/gorm"
)

type TodoListModel struct {
	Conn *gorm.DB
}

// notice 以分数代替

type TodoList struct {
	ID            int64                   `json:"ID" gorm:"column:todoListId"`
	Name          string                  `json:"name"`
	Type          string                  `json:"type"`                             // 类型 每日/每次/每周
	Status        constant.TodoListStatus `json:"status"`                           // 开启 关闭 暂停
	StartTime     int64                   `json:"startTime"`                        // 开始时间
	EndTime       int64                   `json:"endTime"`                          // 结束时间
	Score         int64                   `json:"score"`                            // 总分值
	DailyScore    int64                   `json:"dailyScore"`                       // 每日分值
	CompleteScore int64                   `json:"completeScore"`                    // 奖励分值
	CreateAt      int64                   `json:"createAt" gorm:"column:createAt"`  // 秒级时间戳
	UpdateAt      int64                   `json:"updateAt"  gorm:"column:updateAt"` // 秒级时间戳
}

func (c *TodoList) TableName() string {
	return "TodoList"
}

func NewTodoListModel(db *gorm.DB) *TodoListModel {
	return &TodoListModel{Conn: db.Model(TodoList{})}
}

func (c *TodoListModel) Insert(TodoList *TodoList) error {
	TodoList.CreateAt = time.Now().Unix()
	TodoList.UpdateAt = time.Now().Unix()
	//todo 看看怎么在插入之后,直接获取到 id
	return c.Conn.Create(TodoList).Error
}
func (c *TodoListModel) Update(TodoList *TodoList, ID int64) error {
	TodoList.ID = ID
	TodoList.UpdateAt = time.Now().Unix()
	return c.Conn.Update(TodoList).Error
}

func (c *TodoListModel) FindById(ID int64) (*TodoList, error) {
	TodoList := &TodoList{ID: ID}
	return TodoList, c.Conn.First(TodoList).Error
}

func (c *TodoListModel) FindByName(name string) (*TodoList, error) {
	TodoList := &TodoList{Name: name}
	return TodoList, c.Conn.First(TodoList).Error
}
