package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAnnualGrowthLessInflation struct {
	mock.Mock
}

func (m *MockAnnualGrowthLessInflation) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualGrowthLessInflation) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualGrowthLessInflation) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualGrowthLessInflation) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var annualGrowthLessInflationTests = []struct {
	name  string
	model calculator.Model
}{
	{
		name: "Test Case Basic",
		model: calculator.Model{
			Input: calculator.Input{
				AnnualInvestmentGrowth: 10,
			},
		},
	},
}

func TestNewAnnualGrowthLessInflation(t *testing.T) {
	actual := calculator.NewAnnualGrowthLessInflation()
	expected := calculator.AnnualGrowthLessInflation{}

	assert.Equal(t, expected, actual)
}

func TestAnnualGrowthLessInflationCalculateTraditional(t *testing.T) {
	for _, test := range annualGrowthLessInflationTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualGrowthLessInflation{}

			expected := test.model.Input.AnnualInvestmentGrowth - 0.03
			actual := c.CalculateTraditional(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualGrowthLessInflationCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range annualGrowthLessInflationTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualGrowthLessInflation{}

			expected := c.CalculateTraditional(test.model)
			actual := c.CalculateTraditionalRetirement(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualGrowthLessInflationCalculateRoth(t *testing.T) {
	for _, test := range annualGrowthLessInflationTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualGrowthLessInflation{}

			expected := c.CalculateTraditional(test.model)
			actual := c.CalculateRoth(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualGrowthLessInflationCalculateRothRetirement(t *testing.T) {
	for _, test := range annualGrowthLessInflationTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualGrowthLessInflation{}

			expected := c.CalculateTraditional(test.model)
			actual := c.CalculateRothRetirement(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
