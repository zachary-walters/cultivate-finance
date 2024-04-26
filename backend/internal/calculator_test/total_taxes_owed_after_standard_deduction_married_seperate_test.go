package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionMarriedSeperate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeperate) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeperate) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculate(t *testing.T) {
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
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeperate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeperate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeperate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeperate.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.Calculate(test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeperate.Calculate(test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
