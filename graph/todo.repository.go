package graph

import (
	"github.com/KleKlai/todoappremake/graph/model"
	"gorm.io/gorm"
)

type TodoRepository interface {
	AddUser(model.User) (*model.User, error)
	GetUser(string) (*model.User, error)
	DeleteUser(string) (*model.User, error)
	AddTodo(model.Todo) (*model.Todo, error)
	UpdateTodoTask(model.Todo) (*model.Todo, error)
	DeleteTodo(string) (*model.Todo, error)
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

func (t *todoRepository) DeleteUser(id string) (*model.User, error) {

	var user model.User

	if err := t.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// TODOS

func (t *todoRepository) AddTodo(todo model.Todo) (*model.Todo, error) {

	if err := t.db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *todoRepository) DeleteTodo(id string) (*model.Todo, error) {

	var todo model.Todo

	if err := t.db.Where("id = ?", id).Delete(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t *todoRepository) UpdateTodoTask(todo model.Todo) (*model.Todo, error) {

	if err := t.db.Model(&todo).Where("id = ?", todo.ID).Update("task", todo.Task).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
