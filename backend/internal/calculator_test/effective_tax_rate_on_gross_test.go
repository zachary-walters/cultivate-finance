package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockEffectiveTaxRateOnGross struct {
	mock.Mock
}

func (m *MockEffectiveTaxRateOnGross) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestEffectiveTaxRateOnGrossCalculate(t *testing.T) {
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
			mockTotalTaxesOwedAfterStandardDeduction.On("Calculate", mock.Anything).Return(test.expected)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRetirement", mock.Anything).Return(test.expected)

			c := calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.Calculate(calculator.Model{})
			expected := c.CalculateRetirement(calculator.Model{})

			assert.Equal(t, expected, actual)
		})
	}
}

func TestEffectiveTaxRateOnGrossCalculateRetirement(t *testing.T) {
	tests := []struct {
		name                                 string
		totalTaxesOwedAfterStandardDeduction float64
		model                                calculator.Model
	}{
		{
			name:                                 "Test Case 0",
			totalTaxesOwedAfterStandardDeduction: 12398,
			model: calculator.Model{
				Input: calculator.Input{
					YearlyWithdrawal: 1000.0,
				},
			},
		},
		{
			name:                                 "Test Case 1",
			totalTaxesOwedAfterStandardDeduction: 3294809840,
			model: calculator.Model{
				Input: calculator.Input{
					YearlyWithdrawal: -29382.0,
				},
			},
		},
		{
			name:                                 "Test Case 2",
			totalTaxesOwedAfterStandardDeduction: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRetirement", mock.Anything).Return(test.totalTaxesOwedAfterStandardDeduction)

			c := calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateRetirement(test.model)
			expected := func() float64 {
				if test.model.Input.YearlyWithdrawal == 0.0 {
					return 0
				}

				return test.totalTaxesOwedAfterStandardDeduction / test.model.Input.YearlyWithdrawal
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
