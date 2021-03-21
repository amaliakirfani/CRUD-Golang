package model

type ToDoList struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	Status    string  `json:"status"`
}

type AddToDoList struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
