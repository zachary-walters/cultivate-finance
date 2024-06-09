package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}
func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculate(t *testing.T) {
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
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudectionAndContributions := mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudectionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
