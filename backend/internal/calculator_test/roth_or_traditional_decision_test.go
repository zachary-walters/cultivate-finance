package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockRothOrTraditionalDecision struct {
	mock.Mock
}

func (m *MockRothOrTraditionalDecision) Calculate(model calculator.Model) string {
	args := m.Called(model)
	return args.Get(0).(string)
}

func (m *MockRothOrTraditionalDecision) CalculateRetirement(model calculator.Model) string {
	return m.Calculate(model)
}

func TestRothOrTraditionalDecisionCalculate(t *testing.T) {
	tests := []struct {
		name                    string
		taxRateOfSavings        float64
		effectiveTaxRateOnGross float64
	}{
		{
			name:                    "Test Case Higher TaxRateOfSavings",
			taxRateOfSavings:        1,
			effectiveTaxRateOnGross: 0,
		},
		{
			name:                    "Test Case Higher EffectiveTaxRateOnGross",
			taxRateOfSavings:        0,
			effectiveTaxRateOnGross: 1,
		},
		{
			name:                    "Test Case Equal",
			taxRateOfSavings:        0,
			effectiveTaxRateOnGross: 0,
		},
	}

	testModel := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxRateOfSavings := new(MockTaxRateOfSavings)
			mockEffectiveTaxRateOnGross := new(MockEffectiveTaxRateOnGross)

			mockTaxRateOfSavings.On("Calculate", testModel).Return(test.taxRateOfSavings)
			mockEffectiveTaxRateOnGross.On("Calculate", testModel).Return(test.effectiveTaxRateOnGross)

			c := calculator.RothOrTraditionalDecision{
				TaxRateOfSavingsCalculation:        mockTaxRateOfSavings,
				EffectiveTaxRateOnGrossCalculation: mockEffectiveTaxRateOnGross,
			}

			actual := c.Calculate(testModel)
			expected := func() string {
				if test.taxRateOfSavings >= test.effectiveTaxRateOnGross {
					return "Traditional"
				}

				return "Roth"
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
