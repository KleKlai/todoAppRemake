package infra

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockSource struct {
	mock.Mock
}

func (m *mockSource) Open() error {
	args := m.Called()
	return args.Error(0)
}

func TestOpen(t *testing.T) {
	assertions := require.New(t)
	mock := new(mockSource)

	mock.On("Open").Return(nil)

	assertions.NoError(mock.Open())
	mock.AssertExpectations(t)
}
