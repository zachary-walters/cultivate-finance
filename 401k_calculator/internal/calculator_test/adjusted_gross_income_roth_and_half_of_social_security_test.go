package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAdjustedGrossIncomeRothAndHalfOfSocialSecurity struct {
	mock.Mock
}

func (m *MockAdjustedGrossIncomeRothAndHalfOfSocialSecurity) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncomeRothAndHalfOfSocialSecurity) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestAdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculate(t *testing.T) {
	tests := []struct {
		name                    string
		adjustedGrossIncomeRoth float64
		halfOfSocialSecurity    float64
	}{
		{
			name:                    "Test Case 0",
			adjustedGrossIncomeRoth: 10000,
			halfOfSocialSecurity:    500,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeRoth := new(MockAdjustedGrossIncomeRoth)
			mockHalfOfSocialSecurity := new(MockHalfOfSocialSecurity)

			mockAdjustedGrossIncomeRoth.On("Calculate", model).Return(test.adjustedGrossIncomeRoth)
			mockHalfOfSocialSecurity.On("Calculate", model).Return(test.halfOfSocialSecurity)

			c := &calculator.AdjustedGrossIncomeRothAndHalfOfSocialSecurity{
				AdjustedGrossIncomeRothCalculation: mockAdjustedGrossIncomeRoth,
				HalfOfSocialSecurityCalculation:    mockHalfOfSocialSecurity,
			}

			actual := c.Calculate(model)
			expected := test.adjustedGrossIncomeRoth + test.halfOfSocialSecurity

			assert.Equal(t, expected, actual)
		})
	}
}
