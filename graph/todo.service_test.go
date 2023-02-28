package graph

import (
	"testing"

	"github.com/KleKlai/todoappremake/graph/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type TodoRepositoryMock struct {
	mock.Mock
}

func (m *TodoRepositoryMock) AddUser(user model.User) (*model.User, error) {
	args := m.Called(user)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *TodoRepositoryMock) GetUser(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestAddUser(t *testing.T) {

	m := &TodoRepositoryMock{}
	s := NewTodoService(m)

	u := model.CreateUserInput{}

	t.Run("success", func(t *testing.T) {

		m.On("AddUser", mock.Anything).Return(&u, nil)

		// z := model.CreateUserInput{
		// 	ID:   "1234567",
		// 	Name: "Maynard Magallen",
		// }

		_, err := s.AddUser(&u)

		require.Error(t, err)
		m.AssertExpectations(t)
	})

}
