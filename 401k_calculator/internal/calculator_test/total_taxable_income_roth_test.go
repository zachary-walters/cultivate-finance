package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxableIncome struct {
	mock.Mock
}

func (m *MockTotalTaxableIncome) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncome) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxableIncomeCalculate(t *testing.T) {
	tests := []struct {
		name                        string
		adjustedGrossIncome         float64
		socialSecurityTaxbaleIncome float64
	}{
		{
			name:                        "Test Case 0",
			adjustedGrossIncome:         10000,
			socialSecurityTaxbaleIncome: 500,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockSocialSecurityTaxableIncome := new(MockSocialSecurityTaxableIncome)

			mockAdjustedGrossIncome.On("CalculateTraditional", model).Return(test.adjustedGrossIncome)
			mockSocialSecurityTaxableIncome.On("CalculateTraditional", model).Return(test.socialSecurityTaxbaleIncome)

			c := &calculator.TotalTaxableIncome{
				AdjustedGrossIncomeCalculation:                   mockAdjustedGrossIncome,
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncome,
			}

			actual := c.CalculateTraditional(model)
			expected := test.adjustedGrossIncome + test.socialSecurityTaxbaleIncome

			assert.Equal(t, expected, actual)
		})
	}
}
