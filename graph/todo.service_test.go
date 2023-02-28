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

func (m *TodoRepositoryMock) AddUser(u model.User) (*model.User, error) {
	args := m.Called(u)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *TodoRepositoryMock) GetUser(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func TestGetUser(t *testing.T) {

	mock := &TodoRepositoryMock{}

	s := NewTodoService(mock)

	t.Run("success", func(t *testing.T) {
		mock.On("GetUser", "1234567").Return(&model.User{}, nil)

		_, err := s.GetUser("1234567")

		require.NoError(t, err)
		mock.AssertExpectations(t)
	})
}

func TestAddUser(t *testing.T) {

	m := &TodoRepositoryMock{}

	u := model.CreateUserInput{}

	s := NewTodoService(m)

	t.Run("success", func(t *testing.T) {

		m.On("AddUser", mock.Anything).Return(&model.User{}, nil)

		_, err := s.AddUser(&u)

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}
