package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionSingle struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionSingle) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionSingle) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestMockIncomePerBracketAfterStandardDeductionSingle(t *testing.T) {
	tests := []struct {
		name                                   string
		model                                  calculator.Model
		incomePerBracketAfterStandardDeduction []float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				SingleTaxRates: []calculator.TaxRate{
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
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("Calculate", test.model, test.model.SingleTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.Calculate(test.model)
			expected := c.IncomePerBracketAfterStandardDeductionCalculation.Calculate(test.model, test.model.SingleTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
