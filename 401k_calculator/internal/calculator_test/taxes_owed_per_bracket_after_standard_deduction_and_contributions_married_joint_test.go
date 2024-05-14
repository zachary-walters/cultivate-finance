package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditional(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditionalRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRoth(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint) CalculateRothRetirement(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}
func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointTraditional(t *testing.T) {
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
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint{
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint,
			}

			actual := c.CalculateTraditional(test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation.CalculateTraditional(test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
