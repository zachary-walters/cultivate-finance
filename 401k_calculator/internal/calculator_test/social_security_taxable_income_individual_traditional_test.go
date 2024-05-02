package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeIndividualTraditional struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeIndividualTraditional) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeIndividualTraditional) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestSocialSecurityTaxableIncomeIndividualTraditionalCalculate(t *testing.T) {
	tests := []struct {
		name                                                  string
		model                                                 calculator.Model
		adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					SocialSecurity: 100000,
				},
				SocialSecurityTaxRatesIndividual: []calculator.TaxRate{
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
			adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity := new(MockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity)
			mockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity.On("Calculate", test.model).Return(test.adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity)

			c := &calculator.SocialSecurityTaxableIncomeIndividualTraditional{
				AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation: mockAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				for _, taxRate := range test.model.SocialSecurityTaxRatesIndividual {
					if taxRate.Cap > test.adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesIndividual[len(test.model.SocialSecurityTaxRatesIndividual)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
