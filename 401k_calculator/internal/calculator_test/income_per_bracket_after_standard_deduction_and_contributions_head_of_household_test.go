package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculate(t *testing.T) {
	tests := []struct {
		name                                                   string
		model                                                  calculator.Model
		incomePerBracketAfterStandardDeductionAndContributions []float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				HeadOfHouseholdTaxRates: []calculator.TaxRate{
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
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("Calculate", test.model, test.model.HeadOfHouseholdTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.Calculate(test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.Calculate(test.model, test.model.HeadOfHouseholdTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
