package generator

import (
	"github.com/stretchr/testify/mock"
)

type MockGenerator struct {
	mock.Mock
}

func (m *MockGenerator) Next() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func NewMockGenerator() *MockGenerator {
	return &MockGenerator{}
}
