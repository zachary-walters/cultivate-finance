package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockBalancesTraditional struct {
	mock.Mock
}

func (m *MockBalancesTraditional) Calculate(model calculator.Model) calculator.ChartData {
	args := m.Called(model)
	return args.Get(0).(calculator.ChartData)
}
