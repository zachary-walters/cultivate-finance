package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxableIncomeRoth struct {
	mock.Mock
}

func (m *MockTotalTaxableIncomeRoth) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncomeRoth) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxableIncomeRothCalculate(t *testing.T) {
	tests := []struct {
		name                            string
		adjustedGrossIncomeRoth         float64
		socialSecurityTaxbaleIncomeRoth float64
	}{
		{
			name:                            "Test Case 0",
			adjustedGrossIncomeRoth:         10000,
			socialSecurityTaxbaleIncomeRoth: 500,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeRoth := new(MockAdjustedGrossIncomeRoth)
			mockSocialSecurityTaxableIncomeRoth := new(MockSocialSecurityTaxableIncomeRoth)

			mockAdjustedGrossIncomeRoth.On("Calculate", model).Return(test.adjustedGrossIncomeRoth)
			mockSocialSecurityTaxableIncomeRoth.On("Calculate", model).Return(test.socialSecurityTaxbaleIncomeRoth)

			c := &calculator.TotalTaxableIncomeRoth{
				AdjustedGrossIncomeRothCalculation:                   mockAdjustedGrossIncomeRoth,
				SocialSecurityTaxableIncomeIndividualRothCalculation: mockSocialSecurityTaxableIncomeRoth,
			}

			actual := c.Calculate(model)
			expected := test.adjustedGrossIncomeRoth + test.socialSecurityTaxbaleIncomeRoth

			assert.Equal(t, expected, actual)
		})
	}
}
