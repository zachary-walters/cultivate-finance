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

func TestNewMonthlySequencePaymentsCalculate(t *testing.T) {
	actual := calculator.NewMonthlySequencePayments()
	expected := &calculator.MonthlySequencePayments{
		DebtPayoffMonthCalculation: calculator.NewDebtPayoffMonth(),
		SnowballCalculation:        calculator.NewSnowball(),
	}

	assert.Equal(t, expected, actual)
}

func TestMonthlySequencePaymentsCalculate(t *testing.T) {
	for _, test := range monthlySequencePaymentsTests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoff := new(MockCalculation)
			mockSnowball := new(MockSnowballCalculation)
			mockTotalBeginningDebt := new(MockCalculation)

			mockDebtPayoff.On("Calculate", test.model).Return(test.debtPayoffMonth)
			mockSnowball.On("Calculate", test.model).Return(test.snowball)
			mockTotalBeginningDebt.On("Calculate", test.model).Return(test.totalBeginningDebt)

			c := &calculator.MonthlySequencePayments{
				DebtPayoffMonthCalculation: mockDebtPayoff,
				SnowballCalculation:        mockSnowball,
			}

			actual := c.Calculate(test.model)
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
