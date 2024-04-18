package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeduction) Calculate(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeduction) CalculateRetirement(model calculator.Model, taxRates []calculator.TaxRate) []float64 {
	return m.Calculate(model, taxRates)
}

func TestIncomePerBracketAfterStandardDeductionCalculate(t *testing.T) {
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
			mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)
			mockIncomeAfterStandardDeduction.On("Calculate", test.model).Return(test.incomeAfterStandardDeduction)

			mockIncomePerBracket := new(MockIncomePerBracket)
			mockIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.IncomePerBracketAfterStandardDeduction{
				IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
				IncomePerBracketCalculation:             mockIncomePerBracket,
			}

			actual := c.Calculate(test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}
