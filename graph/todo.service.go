package graph

import (
	"fmt"

	"github.com/KleKlai/todoappremake/graph/model"
)

type TodoService struct {
	repository TodoRepository
}

func NewTodoService(t TodoRepository) *TodoService {
	return &TodoService{t}
}

func (s *TodoService) AddUser(user *model.CreateUserInput) (*model.User, error) {

	// Map the CreateUserInput to User

	u := model.User{
		ID:   user.ID,
		Name: user.Name,
	}

	// Add the user to the database
	res, err := s.repository.AddUser(u)

	if err != nil {
		return nil, fmt.Errorf("Failed to add user: %v", err)
		// return nil, errors.New("Failed to add user")
	}

	return res, nil
}

func (s *TodoService) GetUser(id string) (*model.User, error) {

	// Get the user from the database
	res, err := s.repository.GetUser(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to get user: %v", err)
		// return nil, errors.New("Failed to get user")
	}

	return res, nil
}

func (s *TodoService) DeleteUser(id string) (*model.User, error) {

	res, err := s.repository.DeleteUser(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to delete user: %v", err)
	}

	return res, nil
}

// TODO

func (s *TodoService) AddTodo(todo *model.CreateTodoInput) (*model.Todo, error) {

	// Map the CreateTodoInput to Todo
	t := model.Todo{
		Task:   todo.Task,
		UserID: todo.UserID,
	}

	// Add the todo to the database
	res, err := s.repository.AddTodo(t)

	if err != nil {
		return nil, fmt.Errorf("Failed to add todo: %v", err)
	}

	return res, nil
}

func (s *TodoService) DeleteTodo(id string) (*model.Todo, error) {

	res, err := s.repository.DeleteTodo(id)

	if err != nil {
		return nil, fmt.Errorf("Failed to delete todo: %v", err)
	}

	return res, nil
}

func (s *TodoService) UpdateTodoTask(todo model.UpdateTodoTaskInput) (*model.Todo, error) {

	t := model.Todo{
		ID:   todo.ID,
		Task: todo.Task,
	}
	res, err := s.repository.UpdateTodoTask(t)

	if err != nil {
		return nil, fmt.Errorf("Failed to update todo: %v", err)
	}

	return res, nil
}
