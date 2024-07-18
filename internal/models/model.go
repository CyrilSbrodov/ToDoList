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
	UserID        string `json:"user_id"`
	Id            string `json:"id_list"`
	GroupID       string `json:"group_id"`
	AnotherUserID string `json:"another_user_id"`
	Task          string `json:"post"`
	Group         string `json:"group"`
	Status        bool   `json:"status"`
}
