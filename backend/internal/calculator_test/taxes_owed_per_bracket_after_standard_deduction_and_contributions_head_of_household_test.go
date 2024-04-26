package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(t *testing.T) {
	tests := []struct {
		name                                                  string
		model                                                 calculator.Model
		incomePerBracketAfterStandardDeductionHeadOfHousehold []float64
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
			incomePerBracketAfterStandardDeductionHeadOfHousehold: []float64{
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
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("Calculate", test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.Calculate(test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation.Calculate(test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
