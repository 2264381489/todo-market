package logic

import (
	"context"
	"math"
	"time"
	"todo-market/todoMarket/api/internal/constant"
	"todo-market/todoMarket/api/internal/model"

	"todo-market/todoMarket/api/internal/svc"
	"todo-market/todoMarket/api/internal/types"

	"github.com/jinzhu/now"
	"github.com/zeromicro/go-zero/core/logx"
)

type AddTodoLisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTodoLisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTodoLisLogic {
	return &AddTodoLisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func throwErr(err error) (resp *types.AddTodoListResp, errors error) {
	if err != nil {
		// todo 测试阶段使用panic,上线之后记得删掉
		panic(err)
	}
	logx.Infof(err.Error())
	return nil, err
}

func (l *AddTodoLisLogic) AddTodoLis(req *types.AddTodoListReq) (resp *types.AddTodoListResp, err error) {

	list := &model.TodoList{Name: req.Name, Type: req.Type, Status: constant.ONLINE, StartTime: req.StartTime, EndTime: req.EndTime, Score: req.Score, DailyScore: req.DailyScore, CompleteScore: req.CompleteScore}
	err = l.svcCtx.TodoListModel.Insert(list)
	if err != nil {
		return throwErr(err)
	}

	switch constant.TodoListType(req.Type) {

	case constant.ONCE:
		todoList, err := l.svcCtx.TodoListModel.FindByName(req.Name)
		if err != nil {
			return throwErr(err)
		}
		if todoList != nil {
			//todo 报错,插入todoList已经存在重名函数
			//item := &model.TodoItem{TodoListId: todoList.ID, Name: todoList.Name, Type: todoList.Type, StartTime: todoList.StartTime, EndTime: todoList.EndTime}
			//err = l.svcCtx.TodoItemModel.Insert(item)
			//if err != nil {
			//	return throwErr(err)
			//}
		}
		err = l.svcCtx.TodoListModel.Insert(&model.TodoList{
			Name:          req.Name,
			Type:          req.Type,
			Status:        constant.ONLINE,
			StartTime:     req.StartTime,
			EndTime:       req.EndTime,
			CompleteScore: req.CompleteScore})
		if err != nil {
			return throwErr(err)
		}
	case constant.DAY:
		err = l.svcCtx.TodoListModel.Insert(&model.TodoList{
			Name:          req.Name,
			Type:          req.Type,
			Status:        constant.ONLINE,
			StartTime:     req.StartTime,
			EndTime:       req.EndTime,
			Score:         req.Score,
			CompleteScore: req.CompleteScore})
		if err != nil {
			return throwErr(err)
		}
		todoList, err := l.svcCtx.TodoListModel.FindByName(req.Name)
		if err != nil {
			return throwErr(err)
		}
		startTime := now.New(time.Unix(todoList.StartTime, 0)).BeginningOfDay()
		endTime := now.New(time.Unix(todoList.StartTime, 0)).EndOfDay()
		days := endTime.Sub(startTime).Seconds() / constant.ONE_DAY_SECOND
		totalScore := todoList.Score
		for ; startTime.Before(endTime); startTime = startTime.AddDate(0, 0, 1) {
			score := int64(math.Ceil(float64(todoList.Score-todoList.CompleteScore) / days))
			if totalScore-score <= 0 {
				score = totalScore - score
			}
			err := l.svcCtx.TodoItemModel.Insert(&model.TodoItem{
				TodoListId: todoList.ID,
				Name:       startTime.Format("2006-01-02") + " " + req.Name,
				Type:       todoList.Type,
				StartTime:  todoList.StartTime,
				EndTime:    todoList.EndTime,
				GetScore:   score,
			})
			if err != nil {
				return throwErr(err)
			}
		}
	case constant.WEEK:

	}

	return
}
