package graph

import (
	"github.com/KleKlai/todoappremake/graph/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	AddUser(model.User) (*model.User, error)
	GetUser(string) (*model.User, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db}
}

func (t *todoRepository) AddUser(user model.User) (*model.User, error) {

	if err := t.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *todoRepository) GetUser(id string) (*model.User, error) {

	var user model.User

	if err := t.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
