package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTotalInterestAfterTax struct {
	mock.Mock
}

func (m *MockTotalInterestAfterTax) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestMockTotalInterestAfterTaxCalculate(t *testing.T) {
	tests := []struct {
		name                       string
		totalDisbursementsAfterTax float64
		totalContributions         float64
	}{
		{
			name:                       "Test Case 0",
			totalDisbursementsAfterTax: 100000,
			totalContributions:         1,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursementsAfterTax := new(MockTotalDisbursementsAfterTax)
			mockTotalContributions := new(MockTotalContributions)

			mockTotalDisbursementsAfterTax.On("Calculate", model).Return(test.totalDisbursementsAfterTax)
			mockTotalContributions.On("Calculate", model).Return(test.totalContributions)

			c := &calculator.TotalInterestAfterTax{
				TotalDisbursementsAfterTaxCalculation: mockTotalDisbursementsAfterTax,
				TotalContributionsCalculation:         mockTotalContributions,
			}

			actual := c.Calculate(model)
			expected := test.totalDisbursementsAfterTax - test.totalContributions

			assert.Equal(t, expected, actual)
		})
	}
}
