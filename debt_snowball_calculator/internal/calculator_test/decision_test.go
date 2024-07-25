package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var decisionTests = []struct {
	name                     string
	model                    calculator.Model
	debtPayoffMonthSnowball  float64
	debtPayoffMonthAvalanche float64
	totalPaymentsSnowball    float64
	totalPaymentsAvalanche   float64
}{
	{
		name:                     "Test Case 0 - Either",
		model:                    calculator.Model{},
		debtPayoffMonthSnowball:  10,
		debtPayoffMonthAvalanche: 10,
		totalPaymentsSnowball:    10,
		totalPaymentsAvalanche:   10,
	},
	{
		name:                     "Test Case 1 - Snowball",
		model:                    calculator.Model{},
		debtPayoffMonthSnowball:  10,
		debtPayoffMonthAvalanche: 111,
		totalPaymentsSnowball:    10,
		totalPaymentsAvalanche:   100,
	},
	{
		name:                     "Test Case 2 - Avalanche",
		model:                    calculator.Model{},
		debtPayoffMonthSnowball:  1443,
		debtPayoffMonthAvalanche: 123,
		totalPaymentsSnowball:    100,
		totalPaymentsAvalanche:   43,
	},
}

func TestNewDecision(t *testing.T) {
	actual := calculator.NewDecision()
	expected := &calculator.Decision{
		DebtPayoffMonthCalculation: calculator.NewDebtPayoffMonth(),
		TotalPaymentsCalculation:   calculator.NewTotalPayments(),
	}

	assert.Equal(t, expected, actual)
}

func TestDecisionSnowball(t *testing.T) {
	for _, test := range decisionTests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoff := new(MockCalculation)
			mockDebtPayoff.On("CalculateSnowball", test.model).Return(test.debtPayoffMonthSnowball)
			mockDebtPayoff.On("CalculateAvalanche", test.model).Return(test.debtPayoffMonthAvalanche)

			mockTotalPayments := new(MockCalculation)
			mockTotalPayments.On("CalculateSnowball", test.model).Return(test.totalPaymentsSnowball)
			mockTotalPayments.On("CalculateAvalanche", test.model).Return(test.totalPaymentsAvalanche)

			c := &calculator.Decision{
				DebtPayoffMonthCalculation: mockDebtPayoff,
				TotalPaymentsCalculation:   mockTotalPayments,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() calculator.FinalDecision {
				choice := "Avalanche"
				if test.totalPaymentsSnowball-test.totalPaymentsAvalanche == 0 {
					choice = "Either"
				} else if test.totalPaymentsSnowball-test.totalPaymentsAvalanche < 0 {
					choice = "Snowball"
				}

				return calculator.FinalDecision{
					Choice:                  choice,
					MonthDifference:         test.debtPayoffMonthSnowball - test.debtPayoffMonthAvalanche,
					TotalPaymentsDifference: test.totalPaymentsSnowball - test.totalPaymentsAvalanche,
				}
			}()

			assert.Equal(t, expected, actual)

			switch test.name {
			case "Test Case 0 - Either":
				assert.Equal(t, "Either", actual.Choice)
			case "Test Case 1 - Snowball":
				assert.Equal(t, "Snowball", actual.Choice)
			case "Test Case 2 - Avalanche":
				assert.Equal(t, "Avalanche", actual.Choice)
			}
		})
	}
}

func TestDecisionAvalanche(t *testing.T) {
	for _, test := range decisionTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.NewDecision()

			actual := c.CalculateAvalanche(test.model)
			expected := c.CalculateSnowball(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
