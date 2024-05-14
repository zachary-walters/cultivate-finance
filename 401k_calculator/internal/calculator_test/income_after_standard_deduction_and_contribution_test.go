package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomeAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockIncomeAfterStandardDeductionAndContributions) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeductionAndContributions) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeductionAndContributions) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestIncomeAfterstandardDeductionAndContributionsCalculate(t *testing.T) {
	tests := []struct {
		name                         string
		model                        calculator.Model
		incomeAfterStandardDeduction float64
	}{
		{
			name: "Test Case Basic",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: 1337,
				},
			},
			incomeAfterStandardDeduction: 42,
		},
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: 0,
				},
			},
			incomeAfterStandardDeduction: 42,
		},
		{
			name:  "Test Case Empty",
			model: calculator.Model{},
		},
		{
			name: "Test Case Infinity",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: math.MaxFloat64,
				},
			},
			incomeAfterStandardDeduction: 42,
		},
		{
			name: "Test Case Negative Infinity",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: -math.MaxFloat64,
				},
			},
			incomeAfterStandardDeduction: 42,
		},
	}

	mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)

	c := &calculator.IncomeAfterStandardDeductionAndContributions{
		IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
	}

	for _, test := range tests {
		mockIncomeAfterStandardDeduction.On("CalculateTraditional", mock.Anything).Return(test.incomeAfterStandardDeduction)
		t.Run(test.name, func(t *testing.T) {
			actual := c.CalculateTraditional(test.model)
			expected := func() float64 {
				return float64(c.IncomeAfterStandardDeductionCalculation.CalculateTraditional(test.model)) - test.model.Input.AnnualContributionsPreTax
			}()
			assert.Equal(t, expected, actual)
		})
	}
}
