package model

type Todo struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Message   string `json:"message"`
}

type TodoList struct {
	TodoList []Todo
}
