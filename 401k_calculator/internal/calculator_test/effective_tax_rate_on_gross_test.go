package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockEffectiveTaxRateOnGross struct {
	mock.Mock
}

func (m *MockEffectiveTaxRateOnGross) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestEffectiveTaxRateOnGrossCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                                                 string
		model                                                calculator.Model
		totalTaxesOwedAfterStandardDeductionAndContributions float64
	}{
		{
			name: "Test Case 0",
			totalTaxesOwedAfterStandardDeductionAndContributions: 100,
			model: calculator.Model{
				Input: calculator.Input{
					CurrentAnnualIncome: 41987,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionAndContributions := new(MockTotalTaxesOwedAfterStandardDeductionAndContributions)
			mockTotalTaxesOwedAfterStandardDeductionAndContributions.On("CalculateTraditional", mock.Anything).Return(test.totalTaxesOwedAfterStandardDeductionAndContributions)

			c := &calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: mockTotalTaxesOwedAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditional(test.model)
			expected := func() float64 {
				return test.totalTaxesOwedAfterStandardDeductionAndContributions / test.model.Input.CurrentAnnualIncome
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
