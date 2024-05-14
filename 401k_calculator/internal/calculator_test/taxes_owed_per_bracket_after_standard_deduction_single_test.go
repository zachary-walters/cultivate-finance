package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionSingle struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditional(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditionalRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRoth(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRothRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func TestTaxesOwedPerBracketAfterStandardDeductionSingleCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                                         string
		model                                        calculator.Model
		incomePerBracketAfterStandardDeductionSingle []float64
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
			incomePerBracketAfterStandardDeductionSingle: []float64{
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
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateTraditional", test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionSingleCalculation: mockIncomePerBracketAfterStandardDeductionSingle,
			}

			actual := c.CalculateTraditional(test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditional(test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}