package test

import (
	"github.com/stretchr/testify/mock"

	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

type MockCalculation struct {
	mock.Mock
}

type MockSnowballCalculation struct {
	mock.Mock
}

type MockSequenceCalculation struct {
	mock.Mock
}

func (m *MockCalculation) Calculate(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSnowballCalculation) Calculate(model *calculator.Model) calculator.DebtSequences {
	args := m.Called(model)
	return args.Get(0).(calculator.DebtSequences)
}

func (m *MockSequenceCalculation) Calculate(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}
