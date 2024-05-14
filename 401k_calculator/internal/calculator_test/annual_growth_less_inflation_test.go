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

func TestAnnualGrowthLessInflationCalculateTraditional(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.AnnualGrowthLessInflation{}

			expected := test.model.Input.AnnualInvestmentGrowth - 0.03
			actual := c.CalculateTraditional(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}