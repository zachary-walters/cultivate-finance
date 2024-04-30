package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperateCalculate(t *testing.T) {
	tests := []struct {
		name                                                      string
		model                                                     calculator.Model
		taxesOwedPerBracketAfterStandardDeductionAndContributions []float64
	}{
		{
			name:  "Test Case 0",
			model: calculator.Model{},
			taxesOwedPerBracketAfterStandardDeductionAndContributions: []float64{
				1.0,
				2.0,
				3.0,
				4.0,
				5.0,
				6.0,
			},
		},
	}

	for _, test := range tests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.Calculate(test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudectionAndContributions := mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate.Calculate(test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudectionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
