package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var snowballTests = []struct {
	name  string
	model *calculator.Model
}{
	{
		name: "Test Case 0",
		model: &calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Name:           "debt0",
						Amount:         100,
						MinimumPayment: 1,
						AnnualInterest: 10,
					},
					{
						Name:           "debt1",
						Amount:         10,
						MinimumPayment: 1,
						AnnualInterest: 1,
					},
				},
			},
		},
	},
}

func TestNewSnowball(t *testing.T) {
	actual := calculator.NewSnowball()
	expected := &calculator.Snowball{}

	assert.Equal(t, expected, actual)
}

func TestSnowballCalculate(t *testing.T) {
	for _, test := range snowballTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.Snowball{}

			actual := c.Calculate(test.model)
			expected := func() calculator.DebtSequences {
				debtSequences := calculator.DebtSequences{}

				debts := test.model.Input.Debts
				extraMonthlyPayment := test.model.Input.ExtraMonthlyPayment
				rolloverPayment := 0.0
				oneTimeImmediatePayment := test.model.Input.OneTimeImmediatePayment
				compoundMinimumPayments := 0.0
				maxMonth := 0

				sort.Slice(debts, func(i, j int) bool {
					return debts[i].Amount < debts[j].Amount
				})

				for _, debt := range debts {
					debtBalance := debt.Amount

					debtSequence := calculator.DebtSequence{
						Debt:     debt,
						Months:   []int{},
						Payments: []float64{},
						Balances: []float64{},
					}

					monthIter := 1
					for {
						basePayment := debt.MinimumPayment

						if monthIter == maxMonth {
							basePayment = debt.MinimumPayment + rolloverPayment
						} else if monthIter > maxMonth {
							basePayment = debt.MinimumPayment + extraMonthlyPayment + compoundMinimumPayments
							rolloverPayment = 0
						}

						debtSequence.Months = append(debtSequence.Months, monthIter)

						// use oneTimeImmediatePayment
						debtBalance = debtBalance - oneTimeImmediatePayment
						leftover := (debtBalance - basePayment) * -1
						if debtBalance <= 0 && monthIter == 1 {
							debtSequence.Balances = append(debtSequence.Balances, 0)
							debtSequence.Payments = append(debtSequence.Payments, debt.Amount)
							rolloverPayment = leftover
							oneTimeImmediatePayment = debtBalance * -1
							compoundMinimumPayments += debt.MinimumPayment
							maxMonth = debtSequence.Months[len(debtSequence.Months)-1]
							break
						}

						// use other payments
						debtBalance = debtBalance - basePayment
						if debtBalance <= 0 {
							debtSequence.Balances = append(debtSequence.Balances, 0)
							debtSequence.Payments = append(debtSequence.Payments, basePayment-(debtBalance*-1))
							rolloverPayment = leftover
							compoundMinimumPayments += debt.MinimumPayment
							maxMonth = debtSequence.Months[len(debtSequence.Months)-1]
							break
						}
						debtSequence.Payments = append(debtSequence.Payments, basePayment+oneTimeImmediatePayment)

						oneTimeImmediatePayment = 0

						// increase debtBalance by it's annual interest
						if monthIter != 1 {
							debtBalance = debtBalance + (debtBalance * (debt.AnnualInterest / 100 / 12))
						}

						debtSequence.Balances = append(debtSequence.Balances, debtBalance)
						monthIter += 1
					}

					debtSequences = append(debtSequences, debtSequence)
				}

				return debtSequences
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
