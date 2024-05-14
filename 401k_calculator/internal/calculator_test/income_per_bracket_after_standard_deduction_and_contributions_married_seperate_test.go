package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate) CalculateTraditional(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate) CalculateTraditionalRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate) CalculateRoth(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate) CalculateRothRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperateCalculate(t *testing.T) {
	tests := []struct {
		name                                                   string
		model                                                  calculator.Model
		incomePerBracketAfterStandardDeductionAndContributions []float64
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
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateTraditional", test.model, test.model.MarriedSeperateTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeperate{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditional(test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(test.model, test.model.MarriedSeperateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}