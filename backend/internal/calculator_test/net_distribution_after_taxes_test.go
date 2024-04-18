package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockNetDistributionAfterTaxes struct {
	mock.Mock
}

func (m *MockNetDistributionAfterTaxes) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockNetDistributionAfterTaxes) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestNetDistributionAfterTaxesCalculate(t *testing.T) {
	tests := []struct {
		name     string
		expected float64
	}{
		{
			name:     "Test Case 0",
			expected: 12398,
		},
		{
			name:     "Test Case 1",
			expected: 3294809840,
		},
		{
			name:     "Test Case 2",
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRetirement", mock.Anything).Return(test.expected)

			c := calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.Calculate(calculator.Model{})
			expected := c.CalculateRetirement(calculator.Model{})

			assert.Equal(t, expected, actual)
		})
	}
}

func TestNetDistributionAfterTaxesCalculateRetirement(t *testing.T) {
	tests := []struct {
		name                                string
		model                               calculator.Model
		totalTaxesOwedAfterStandadDeduction float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					YearlyWithdrawal: 1000,
				},
			},
			totalTaxesOwedAfterStandadDeduction: 1337,
		},
		{
			name: "Test Case 1",
			model: calculator.Model{
				Input: calculator.Input{
					YearlyWithdrawal: 0,
				},
			},
			totalTaxesOwedAfterStandadDeduction: 1337,
		},
		{
			name: "Test Case 2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRetirement", test.model).Return(test.totalTaxesOwedAfterStandadDeduction)

			c := &calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateRetirement(test.model)
			expected := test.model.Input.YearlyWithdrawal - test.totalTaxesOwedAfterStandadDeduction

			assert.Equal(t, expected, actual)
		})
	}
}
