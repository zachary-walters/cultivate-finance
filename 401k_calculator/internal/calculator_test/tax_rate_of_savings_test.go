package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxRateOfSavings struct {
	mock.Mock
}

func (m *MockTaxRateOfSavings) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxRateOfSavings) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxRateOfSavings) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxRateOfSavings) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var taxRateOfSavingsTest = []struct {
	name                             string
	model                            calculator.Model
	annualTaxSavingsWithContribution float64
}{
	{
		name: "Test Case Basic",
		model: calculator.Model{
			Input: calculator.Input{
				AnnualContributionsPreTax: 1000,
			},
		},
		annualTaxSavingsWithContribution: 1000,
	},
	{
		name: "Test Case Divide By 0",
		model: calculator.Model{
			Input: calculator.Input{
				AnnualContributionsPreTax: 0,
			},
		},
		annualTaxSavingsWithContribution: 1000,
	},
	{
		name: "Test Case Divide 0 By",
		model: calculator.Model{
			Input: calculator.Input{
				AnnualContributionsPreTax: 1000,
			},
		},
		annualTaxSavingsWithContribution: 0,
	},
}

func TestNewTaxRateOfSavings(t *testing.T) {
	actual := calculator.NewTaxRateOfSavings()
	expected := calculator.TaxRateOfSavings{
		AnnualTaxSavingsWithContributionCalculation: calculator.NewAnnualTaxSavingsWithContribution(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxRateOfSavingsCalculateTraditional(t *testing.T) {
	for _, test := range taxRateOfSavingsTest {
		t.Run(test.name, func(t *testing.T) {
			mockAnnualTaxSavingsWithContribution := new(MockAnnualTaxSavingsWithContribution)
			mockAnnualTaxSavingsWithContribution.On("CalculateTraditional", &test.model).Return(test.annualTaxSavingsWithContribution)

			c := &calculator.TaxRateOfSavings{
				AnnualTaxSavingsWithContributionCalculation: mockAnnualTaxSavingsWithContribution,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := test.annualTaxSavingsWithContribution / test.model.Input.AnnualContributionsPreTax

			assert.Equal(t, actual, expected)
		})
	}
}

func TestTaxRateOfSavingsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxRateOfSavingsTest {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TaxRateOfSavings{}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := 0.0

			assert.Equal(t, actual, expected)
		})
	}
}

func TestTaxRateOfSavingsCalculateRoth(t *testing.T) {
	for _, test := range taxRateOfSavingsTest {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TaxRateOfSavings{}

			actual := c.CalculateRoth(&test.model)
			expected := 0.0

			assert.Equal(t, actual, expected)
		})
	}
}

func TestTaxRateOfSavingsCalculateRothRetirement(t *testing.T) {
	for _, test := range taxRateOfSavingsTest {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TaxRateOfSavings{}

			actual := c.CalculateRothRetirement(&test.model)
			expected := 0.0

			assert.Equal(t, actual, expected)
		})
	}
}
