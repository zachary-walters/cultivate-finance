package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculate(t *testing.T) {
	tests := []struct {
		name                                                  string
		model                                                 calculator.Model
		incomePerBracketAfterStandardDeductionMarriedSeperate []float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				MarriedSeperateTaxRates: []calculator.TaxRate{
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
			incomePerBracketAfterStandardDeductionMarriedSeperate: []float64{
				1.0,
				1.0,
				1.0,
				1.0,
				1.0,
				1.0,
				1.0,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate.On("Calculate", test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeperate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeperate{
				IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeperate,
			}

			actual := c.Calculate(test.model)

			incomePerBracketAfterStandarddeductionMarriedSeperate := c.IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation.Calculate(test.model)
			expected := make([]float64, len(test.model.MarriedSeperateTaxRates))

			for idx, taxRate := range test.model.MarriedSeperateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeperate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
