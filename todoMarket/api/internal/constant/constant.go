package constant

type TodoListStatus string

// ONLINE todoList 状态
const (
	ONLINE TodoListStatus = "online"
)

func (t TodoListStatus) String() string {
	return string(t)
}

type TodoListType string

const (
	ONCE  TodoListType = "once"
	DAY   TodoListType = "day"
	WEEK  TodoListType = "weak"
	MONTH TodoListType = "month"
)

func (t TodoListType) String() string {
	return string(t)
}

const ONE_DAY_SECOND float64 = 3600 * 24
