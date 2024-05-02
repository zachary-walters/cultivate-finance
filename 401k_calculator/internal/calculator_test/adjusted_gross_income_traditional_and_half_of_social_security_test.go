package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity struct {
	mock.Mock
}

func (m *MockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculate(t *testing.T) {
	tests := []struct {
		name                           string
		adjustedGrossIncomeTraditional float64
		halfOfSocialSecurity           float64
	}{
		{
			name:                           "Test Case 0",
			adjustedGrossIncomeTraditional: 10000,
			halfOfSocialSecurity:           500,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeTraditional := new(MockAdjustedGrossIncomeTraditional)
			mockHalfOfSocialSecurity := new(MockHalfOfSocialSecurity)

			mockAdjustedGrossIncomeTraditional.On("Calculate", model).Return(test.adjustedGrossIncomeTraditional)
			mockHalfOfSocialSecurity.On("Calculate", model).Return(test.halfOfSocialSecurity)

			c := &calculator.AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity{
				AdjustedGrossIncomeTraditionalCalculation: mockAdjustedGrossIncomeTraditional,
				HalfOfSocialSecurityCalculation:           mockHalfOfSocialSecurity,
			}

			actual := c.Calculate(model)
			expected := test.adjustedGrossIncomeTraditional + test.halfOfSocialSecurity

			assert.Equal(t, expected, actual)
		})
	}
}
