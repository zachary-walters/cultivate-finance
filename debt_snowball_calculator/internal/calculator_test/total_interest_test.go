package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var totalInterestTests = []struct {
	name               string
	model              *calculator.Model
	totalBeginningDebt float64
	totalPayments      float64
}{
	{
		name:  "Test Case 0",
		model: &calculator.Model{},
	},
	{
		name:               "Test Case 1",
		model:              &calculator.Model{},
		totalBeginningDebt: 1293012,
		totalPayments:      12093810984,
	},
	{
		name:               "Test Case 2",
		model:              &calculator.Model{},
		totalBeginningDebt: 1293012,
		totalPayments:      1,
	},
	{
		name:               "Test Case 3",
		model:              &calculator.Model{},
		totalBeginningDebt: 1,
		totalPayments:      12093810984,
	},
}

func TestNewTotalInterest(t *testing.T) {
	actual := calculator.NewTotalInterest()
	expected := &calculator.TotalInterest{
		TotalBeginningDebtCalculation: calculator.NewTotalBeginningDebt(),
		TotalPaymentsCalculation:      calculator.NewTotalPayments(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalInterest(t *testing.T) {
	for _, test := range totalInterestTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalBeginningDebt := new(MockCalculation)
			mockTotalPayments := new(MockCalculation)

			mockTotalBeginningDebt.On("CalculateSnowball", test.model).Return(test.totalBeginningDebt)
			mockTotalPayments.On("CalculateSnowball", test.model).Return(test.totalPayments)

			c := &calculator.TotalInterest{
				TotalBeginningDebtCalculation: mockTotalBeginningDebt,
				TotalPaymentsCalculation:      mockTotalPayments,
			}

			actual := c.CalculateSnowball(test.model)
			expected := c.SanitizeToZero(test.totalPayments - test.totalBeginningDebt)

			if test.totalPayments-test.totalBeginningDebt < 0 {
				assert.Zero(t, actual)
			}

			assert.Equal(t, expected, actual)
		})
	}
}
