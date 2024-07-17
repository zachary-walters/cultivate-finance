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

type MockValidDebtsCalculation struct {
	mock.Mock
}

func (m *MockCalculation) CalculateSnowball(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCalculation) CalculateAvalanche(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSnowballCalculation) CalculateSnowball(model *calculator.Model) calculator.DebtSequences {
	args := m.Called(model)
	return args.Get(0).(calculator.DebtSequences)
}

func (m *MockSnowballCalculation) CalculateAvalanche(model *calculator.Model) calculator.DebtSequences {
	args := m.Called(model)
	return args.Get(0).(calculator.DebtSequences)
}

func (m *MockSequenceCalculation) CalculateSnowball(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockSequenceCalculation) CalculateAvalanche(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockValidDebtsCalculation) CalculateSnowball(model *calculator.Model) []calculator.Debt {
	args := m.Called(model)
	return args.Get(0).([]calculator.Debt)
}

func (m *MockValidDebtsCalculation) CalculateAvalanche(model *calculator.Model) []calculator.Debt {
	args := m.Called(model)
	return args.Get(0).([]calculator.Debt)
}
