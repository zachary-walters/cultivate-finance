package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var monthlySequencePaymentsTests = []struct {
	name               string
	model              *calculator.Model
	debtPayoffMonth    float64
	snowball           calculator.DebtSequences
	totalBeginningDebt float64
}{
	{
		name:            "Test Case 0",
		model:           &calculator.Model{},
		debtPayoffMonth: 30,
		snowball: calculator.DebtSequences{
			{
				Payments: []float64{1, 2, 3, 4, 5, 6},
			},
			{
				Payments: []float64{1, 2, 3, 4, 5, 6},
			},
			{
				Payments: []float64{0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100},
			},
		},
		totalBeginningDebt: 1337,
	},
}

func TestNewMonthlySequencePaymentsCalculateSnowball(t *testing.T) {
	actual := calculator.NewMonthlySequencePayments()
	expected := &calculator.MonthlySequencePayments{
		DebtPayoffMonthCalculation:   calculator.NewDebtPayoffMonth(),
		SnowballAvalancheCalculation: calculator.NewSnowballAvalanche(),
	}

	assert.Equal(t, expected, actual)
}

func TestMonthlySequencePaymentsCalculateSnowball(t *testing.T) {
	for _, test := range monthlySequencePaymentsTests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoff := new(MockCalculation)
			mockSnowball := new(MockSnowballCalculation)
			mockTotalBeginningDebt := new(MockCalculation)

			mockDebtPayoff.On("CalculateSnowball", test.model).Return(test.debtPayoffMonth)
			mockSnowball.On("CalculateSnowball", test.model).Return(test.snowball)
			mockTotalBeginningDebt.On("CalculateSnowball", test.model).Return(test.totalBeginningDebt)

			c := &calculator.MonthlySequencePayments{
				DebtPayoffMonthCalculation:   mockDebtPayoff,
				SnowballAvalancheCalculation: mockSnowball,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() []float64 {
				balances := []float64{}

				for i := 0; i < int(test.debtPayoffMonth); i++ {
					balance := 0.0
					for _, debtSequence := range test.snowball {
						if len(debtSequence.Payments) > i {
							balance += debtSequence.Payments[i]
						}
					}

					balances = append(balances, balance)
				}

				return balances
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMonthlySequencePaymentsCalculateAvalanche(t *testing.T) {
	for _, test := range monthlySequencePaymentsTests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoff := new(MockCalculation)
			mockSnowball := new(MockSnowballCalculation)
			mockTotalBeginningDebt := new(MockCalculation)

			mockDebtPayoff.On("CalculateAvalanche", test.model).Return(test.debtPayoffMonth)
			mockSnowball.On("CalculateAvalanche", test.model).Return(test.snowball)
			mockTotalBeginningDebt.On("CalculateAvalanche", test.model).Return(test.totalBeginningDebt)

			c := &calculator.MonthlySequencePayments{
				DebtPayoffMonthCalculation:   mockDebtPayoff,
				SnowballAvalancheCalculation: mockSnowball,
			}

			actual := c.CalculateAvalanche(test.model)
			expected := func() []float64 {
				balances := []float64{}

				for i := 0; i < int(test.debtPayoffMonth); i++ {
					balance := 0.0
					for _, debtSequence := range test.snowball {
						if len(debtSequence.Payments) > i {
							balance += debtSequence.Payments[i]
						}
					}

					balances = append(balances, balance)
				}

				return balances
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
