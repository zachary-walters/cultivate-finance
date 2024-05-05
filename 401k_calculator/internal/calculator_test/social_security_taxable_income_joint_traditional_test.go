package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeJointTraditional struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeJointTraditional) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeJointTraditional) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestSocialSecurityTaxableIncomeJointTraditionalCalculate(t *testing.T) {
	tests := []struct {
		name                           string
		model                          calculator.Model
		adjustedGrossIncomeTraditional float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					SocialSecurity: 100000,
				},
				SocialSecurityTaxRatesJoint: []calculator.TaxRate{
					{
						Cap:  25000,
						Rate: 0.0,
					},
					{
						Cap:  34000,
						Rate: 0.5,
					},
					{
						Cap:  -1,
						Rate: 0.85,
					},
				},
			},
			adjustedGrossIncomeTraditional: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeTraditional := new(MockAdjustedGrossIncomeTraditional)
			mockAdjustedGrossIncomeTraditional.On("Calculate", test.model).Return(test.adjustedGrossIncomeTraditional)

			c := &calculator.SocialSecurityTaxableIncomeJointTraditional{
				AdjustedGrossIncomeTraditionalCalculation: mockAdjustedGrossIncomeTraditional,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				for _, taxRate := range test.model.SocialSecurityTaxRatesJoint {
					if taxRate.Cap > test.adjustedGrossIncomeTraditional {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesJoint[len(test.model.SocialSecurityTaxRatesJoint)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
