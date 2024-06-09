package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculate(t *testing.T) {
	tests := []struct {
		name                                       string
		model                                      calculator.Model
		taxesOwedPerBracketAfterStandardDeductions []float64
	}{
		{
			name:  "Test Case 0",
			model: calculator.Model{},
			taxesOwedPerBracketAfterStandardDeductions: []float64{
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
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
