package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockBalancesRothMatchingGross struct {
	mock.Mock
}

func (m *MockBalancesRothMatchingGross) Calculate(model calculator.Model) calculator.ChartData {
	args := m.Called(model)
	return args.Get(0).(calculator.ChartData)
}
