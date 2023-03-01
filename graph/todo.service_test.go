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

func (m *TodoRepositoryMock) DeleteUser(id string) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *TodoRepositoryMock) AddTodo(t model.Todo) (*model.Todo, error) {
	args := m.Called(t)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) DeleteTodo(id string) (*model.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) UpdateTodoTask(t model.Todo) (*model.Todo, error) {
	args := m.Called(t)
	return args.Get(0).(*model.Todo), args.Error(1)
}

func (m *TodoRepositoryMock) UpdateTodoStatus(t model.Todo) (*model.Todo, error) {
	args := m.Called(t)
	return args.Get(0).(*model.Todo), args.Error(1)
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

func TestDeleteUser(t *testing.T) {

	m := &TodoRepositoryMock{}

	s := NewTodoService(m)

	t.Run("success", func(t *testing.T) {

		m.On("DeleteUser", mock.Anything).Return(&model.User{}, nil)

		_, err := s.DeleteUser(mock.Anything)

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}

func TestAddTodo(t *testing.T) {

	m := &TodoRepositoryMock{}

	s := NewTodoService(m)

	t.Run("success", func(t *testing.T) {

		m.On("AddTodo", mock.Anything).Return(&model.Todo{}, nil)

		_, err := s.AddTodo(&model.CreateTodoInput{})

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}

func TestDeleteTodo(t *testing.T) {

	m := &TodoRepositoryMock{}

	s := NewTodoService(m)

	t.Run("success", func(t *testing.T) {

		m.On("DeleteTodo", mock.Anything).Return(&model.Todo{}, nil)

		_, err := s.DeleteTodo(mock.Anything)

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}

func TestUpdateTodoTask(t *testing.T) {

	m := &TodoRepositoryMock{}

	s := NewTodoService(m)

	input := model.UpdateTodoTaskInput{}

	t.Run("success", func(t *testing.T) {

		m.On("UpdateTodoTask", mock.Anything).Return(&model.Todo{}, nil)

		_, err := s.UpdateTodoTask(input)

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}

func TestUpdateTodoStatus(t *testing.T) {

	m := &TodoRepositoryMock{}

	s := NewTodoService(m)

	input := model.UpdateTodoStatusInput{}

	t.Run("success", func(t *testing.T) {

		m.On("UpdateTodoStatus", mock.Anything).Return(&model.Todo{}, nil)

		_, err := s.UpdateTodoStatus(input)

		require.Nil(t, err)
		m.AssertExpectations(t)
	})
}
