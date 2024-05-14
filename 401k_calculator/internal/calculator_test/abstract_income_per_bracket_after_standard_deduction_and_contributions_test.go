package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAbstractIncomePerBracketAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockAbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditional(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateRoth(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeductionAndContributions) CalculateRothRetirement(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func TestIncomePerBracketAfterStandarddeductionAndContributionsCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                         string
		model                        calculator.Model
		taxRates                     []calculator.TaxRate
		incomeAfterStandardDeduction float64
		incomePerBracket             float64
	}{
		{
			name: "Test Case 0",
			taxRates: []calculator.TaxRate{
				{
					Cap:  12.0,
					Rate: 0.123,
				},
				{
					Cap:  4214.0,
					Rate: 0.646546,
				},
				{
					Cap:  4564.0,
					Rate: 0.231,
				},
			},
			incomeAfterStandardDeduction: 2000.0,
			incomePerBracket:             3943.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomeAfterStandardDeductionAndContributions := new(MockIncomeAfterStandardDeductionAndContributions)
			mockIncomeAfterStandardDeductionAndContributions.On("CalculateTraditional", test.model).Return(test.incomeAfterStandardDeduction)

			mockAbstractIncomePerBracket := new(MockAbstractIncomePerBracket)
			mockAbstractIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.AbstractIncomePerBracketAfterStandardDeductionAndContributions{
				IncomeAfterStandardDeductionAndContributionsCalculation: mockIncomeAfterStandardDeductionAndContributions,
				AbstractIncomePerBracketCalculation:                     mockAbstractIncomePerBracket,
			}

			actual := c.CalculateTraditional(test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}
