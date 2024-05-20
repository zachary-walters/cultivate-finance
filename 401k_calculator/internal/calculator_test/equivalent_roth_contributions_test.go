package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockEquivalentRothContributions struct {
	mock.Mock
}

func (m *MockEquivalentRothContributions) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEquivalentRothContributions) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEquivalentRothContributions) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEquivalentRothContributions) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestEquivalentRothContributionsCalculate(t *testing.T) {
	tests := []struct {
		name                             string
		model                            calculator.Model
		annualTaxSavingsWithContribution float64
	}{
		{
			name: "Test Case Basic 0",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: 100,
				},
			},
			annualTaxSavingsWithContribution: 10,
		},
		{
			name: "Test Case Basic 1",
			model: calculator.Model{
				Input: calculator.Input{
					AnnualContributionsPreTax: 457395,
				},
			},
			annualTaxSavingsWithContribution: 2345987587,
		},
		{
			name: "Test Case Empty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockAnnualTaxSavingsWithContribution := new(MockAnnualTaxSavingsWithContribution)
			mockAnnualTaxSavingsWithContribution.On("CalculateTraditional", &test.model).Return(test.annualTaxSavingsWithContribution)

			c := &calculator.EquivalentRothContributions{
				AnnualTaxSavingsWithContributionCalculation: mockAnnualTaxSavingsWithContribution,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := test.model.Input.AnnualContributionsPreTax - test.annualTaxSavingsWithContribution

			assert.Equal(t, actual, expected)
		})
	}
}
