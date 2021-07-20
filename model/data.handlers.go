package model

func New() TodoList {
	return TodoList{
		TodoList: []Todo{},
	}
}

func (thisList *TodoList) Add(todo Todo) {
	thisList.TodoList = append(thisList.TodoList, todo)
}

func (thisList *TodoList) GetAll() []Todo {
	return thisList.TodoList
}

func (thisList *TodoList) Delete(todo Todo) {
	newList := New().TodoList
	for i := 0; i < len(thisList.TodoList); i++ {
		if thisList.TodoList[i].ID != todo.ID {
			newList = append(newList, thisList.TodoList[i])
		}
	}
	thisList.TodoList = newList
}

func (thisList *TodoList) Update(todo Todo) {
	for i := 0; i < len(thisList.TodoList); i++ {
		if thisList.TodoList[i].ID == todo.ID {
			thisList.TodoList[i] = todo
		}
	}
}
