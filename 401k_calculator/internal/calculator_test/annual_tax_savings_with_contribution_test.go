package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAnnualTaxSavingsWithContribution struct {
	mock.Mock
}

func (m *MockAnnualTaxSavingsWithContribution) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualTaxSavingsWithContribution) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestAnnualTaxSavingsWithContributionCalculate(t *testing.T) {
	tests := []struct {
		name                                                 string
		totalTaxesOwedAfterStandardDeduction                 float64
		totalTaxesOwedAfterStandardDeductionAndContributions float64
	}{
		{
			name:                                 "Test Case Basic",
			totalTaxesOwedAfterStandardDeduction: 200.0,
			totalTaxesOwedAfterStandardDeductionAndContributions: 100.0,
		},
		{
			name:                                 "Test Case Zero 0",
			totalTaxesOwedAfterStandardDeduction: 0.0,
			totalTaxesOwedAfterStandardDeductionAndContributions: 100.0,
		},
		{
			name:                                 "Test Case Zero 1",
			totalTaxesOwedAfterStandardDeduction: 342817,
			totalTaxesOwedAfterStandardDeductionAndContributions: 0,
		},
		{
			name:                                 "Test Case Infinity",
			totalTaxesOwedAfterStandardDeduction: math.Inf(1),
			totalTaxesOwedAfterStandardDeductionAndContributions: 985908,
		},
		{
			name:                                 "Test Case Negative Infinity",
			totalTaxesOwedAfterStandardDeduction: math.Inf(-1),
			totalTaxesOwedAfterStandardDeductionAndContributions: 985908,
		},
	}

	for _, test := range tests {
		totalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
		totalTaxesOwedAfterStandardDeductionAndContributions := new(MockTotalTaxesOwedAfterStandardDeductionAndContributions)

		model := calculator.Model{}

		c := &calculator.AnnualTaxSavingsWithContribution{
			TotalTaxesOwedAfterStandardDeductionCalculation:                 totalTaxesOwedAfterStandardDeduction,
			TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: totalTaxesOwedAfterStandardDeductionAndContributions,
		}

		totalTaxesOwedAfterStandardDeduction.On("Calculate", model).Return(test.totalTaxesOwedAfterStandardDeduction)
		totalTaxesOwedAfterStandardDeductionAndContributions.On("Calculate", model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributions)

		actual := c.Calculate(model)
		expected := float64(test.totalTaxesOwedAfterStandardDeduction) - float64(test.totalTaxesOwedAfterStandardDeductionAndContributions)

		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, expected, actual)
		})
	}
}
