package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}
func TestTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateTraditional(t *testing.T) {
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
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
