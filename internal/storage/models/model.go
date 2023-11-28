package models

type User struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Group    []string   `json:"groups"`
	ToDoList []TodoList `json:"to_do_list"`
}

type TodoList struct {
	Id     string `json:"id_list"`
	Theme  string `json:"theme"`
	Post   string `json:"post"`
	Group  string `json:"group"`
	Status bool   `json:"status"`
}
