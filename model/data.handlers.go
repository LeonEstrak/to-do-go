package model

func New() *TodoList {
	return &TodoList{
		TodoList: []Todo{},
	}
}

func (thisList *TodoList) Add(todo Todo) {
	thisList.TodoList = append(thisList.TodoList, todo)
}

func (thisList *TodoList) GetAll() []Todo {
	return thisList.TodoList
}
