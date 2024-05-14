package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomeAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockIncomeAfterStandardDeduction) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestIncomeAfterStandardDeductionCalculate(t *testing.T) {
	mockStandardDeduction := new(MockStandardDeduction)
	mockStandardDeduction.On("CalculateTraditional", mock.Anything).Return(10000.0)

	calc := &calculator.IncomeAfterStandardDeduction{
		StandardDeductionCalculation: mockStandardDeduction,
	}

	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case Basic",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentAnnualIncome: 60000,
				},
			},
		},
		{
			name: "Test Case Infinity",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentAnnualIncome: math.Inf(1),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calc.CalculateTraditional(test.model)
			expected := test.model.Input.CurrentAnnualIncome - float64(mockStandardDeduction.CalculateTraditional(test.model))
			assert.Equal(t, expected, result)
		})
	}
}
