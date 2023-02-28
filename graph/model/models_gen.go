// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTodoInput struct {
	Task   string `json:"task"`
	UserID string `json:"userId"`
}

type CreateUserInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Todo struct {
	ID     string `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
	UserID string `json:"userID"`
}

type User struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Todos []*Todo `json:"todos" gorm:"foreignKey:UserID;references:ID"`
}
