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
