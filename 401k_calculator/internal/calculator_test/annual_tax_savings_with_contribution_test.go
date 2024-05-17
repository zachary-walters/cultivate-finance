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

func (m *MockAnnualTaxSavingsWithContribution) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualTaxSavingsWithContribution) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualTaxSavingsWithContribution) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualTaxSavingsWithContribution) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var annualTaxSavingsWithContributionTests = []struct {
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

func TestNewAnnualTaxSavingsWithContribution(t *testing.T) {
	actual := calculator.NewAnnualTaxSavingsWithContribution()
	expected := calculator.AnnualTaxSavingsWithContribution{
		TotalTaxesOwedAfterStandardDeductionCalculation:                 calculator.NewTotalTaxesOwedAfterStandardDeduction(),
		TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
	}

	assert.Equal(t, actual, expected)
}

func TestAnnualTaxSavingsWithContributionCalculateTraditional(t *testing.T) {
	for _, test := range annualTaxSavingsWithContributionTests {
		totalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
		totalTaxesOwedAfterStandardDeductionAndContributions := new(MockTotalTaxesOwedAfterStandardDeductionAndContributions)

		model := calculator.Model{}

		c := &calculator.AnnualTaxSavingsWithContribution{
			TotalTaxesOwedAfterStandardDeductionCalculation:                 totalTaxesOwedAfterStandardDeduction,
			TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: totalTaxesOwedAfterStandardDeductionAndContributions,
		}

		totalTaxesOwedAfterStandardDeduction.On("CalculateTraditional", model).Return(test.totalTaxesOwedAfterStandardDeduction)
		totalTaxesOwedAfterStandardDeductionAndContributions.On("CalculateTraditional", model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributions)

		actual := c.CalculateTraditional(model)
		expected := float64(test.totalTaxesOwedAfterStandardDeduction) - float64(test.totalTaxesOwedAfterStandardDeductionAndContributions)

		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualTaxSavingsWithContributionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range annualTaxSavingsWithContributionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualTaxSavingsWithContribution{}

			actual := c.CalculateTraditionalRetirement(calculator.Model{})
			expected := float64(0)

			assert.Equal(t, expected, actual)
			assert.Zero(t, actual)
		})
	}
}

func TestAnnualTaxSavingsWithContributionCalculateRoth(t *testing.T) {
	for _, test := range annualTaxSavingsWithContributionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualTaxSavingsWithContribution{}

			actual := c.CalculateRoth(calculator.Model{})
			expected := float64(0)

			assert.Equal(t, expected, actual)
			assert.Zero(t, actual)
		})
	}
}

func TestAnnualTaxSavingsWithContributionCalculateRothRetirement(t *testing.T) {
	for _, test := range annualTaxSavingsWithContributionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualTaxSavingsWithContribution{}

			actual := c.CalculateRothRetirement(calculator.Model{})
			expected := float64(0)

			assert.Equal(t, expected, actual)
			assert.Zero(t, actual)
		})
	}
}
