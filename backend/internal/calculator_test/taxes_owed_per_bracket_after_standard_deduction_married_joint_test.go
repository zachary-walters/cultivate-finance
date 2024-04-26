package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculate(t *testing.T) {
	tests := []struct {
		name                                               string
		model                                              calculator.Model
		incomePerBracketAfterStandardDeductionMarriedJoint []float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				MarriedJointTaxRates: []calculator.TaxRate{
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
			incomePerBracketAfterStandardDeductionMarriedJoint: []float64{
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
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("Calculate", test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionMarriedJoint,
			}

			actual := c.Calculate(test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.Calculate(test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
