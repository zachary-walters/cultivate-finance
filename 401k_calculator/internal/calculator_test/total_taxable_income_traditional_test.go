package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxableIncomeTraditional struct {
	mock.Mock
}

func (m *MockTotalTaxableIncomeTraditional) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncomeTraditional) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxableIncomeTraditionalCalculate(t *testing.T) {
	tests := []struct {
		name                                   string
		adjustedGrossIncomeTraditional         float64
		socialSecurityTaxbaleIncomeTraditional float64
	}{
		{
			name:                                   "Test Case 0",
			adjustedGrossIncomeTraditional:         10000,
			socialSecurityTaxbaleIncomeTraditional: 500,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeTraditional := new(MockAdjustedGrossIncomeTraditional)
			mockSocialSecurityTaxableIncomeTraditional := new(MockSocialSecurityTaxableIncomeTraditional)

			mockAdjustedGrossIncomeTraditional.On("Calculate", model).Return(test.adjustedGrossIncomeTraditional)
			mockSocialSecurityTaxableIncomeTraditional.On("Calculate", model).Return(test.socialSecurityTaxbaleIncomeTraditional)

			c := &calculator.TotalTaxableIncomeTraditional{
				AdjustedGrossIncomeTraditionalCalculation:                   mockAdjustedGrossIncomeTraditional,
				SocialSecurityTaxableIncomeIndividualTraditionalCalculation: mockSocialSecurityTaxableIncomeTraditional,
			}

			actual := c.Calculate(model)
			expected := test.adjustedGrossIncomeTraditional + test.socialSecurityTaxbaleIncomeTraditional

			assert.Equal(t, expected, actual)
		})
	}
}
