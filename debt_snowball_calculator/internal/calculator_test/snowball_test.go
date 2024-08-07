package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var snowballTests = []struct {
	name  string
	model calculator.Model
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
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
	{
		name: "Test Case 1 - break out early",
		model: calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Name:           "debt0",
						Amount:         10000000000000000,
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
	{
		name: "Test Case 2",
		model: calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Name:           "debt1",
						Amount:         10,
						MinimumPayment: 1,
						AnnualInterest: 1,
					},
					{
						Name:           "debt0",
						Amount:         100,
						MinimumPayment: 1,
						AnnualInterest: 10,
					},
				},
			},
		},
	},
	{
		name: "Test Case 3 - break out early",
		model: calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Name:           "debt1",
						Amount:         10000000000000000,
						MinimumPayment: 1,
						AnnualInterest: 10,
					},
					{
						Name:           "debt0",
						Amount:         10,
						MinimumPayment: 1,
						AnnualInterest: 11,
					},
				},
			},
		},
	},
}

func TestNewSnowball(t *testing.T) {
	actual := calculator.NewSnowballAvalanche()
	expected := &calculator.SnowballAvalanche{
		MaxYear: 200,
	}

	assert.Equal(t, expected, actual)
}

func TestSnowballCalculateSnowball(t *testing.T) {
	for _, test := range snowballTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.SnowballAvalanche{
				MaxYear: 200,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() calculator.DebtSequences {
				extraMonthlyPayment := test.model.Input.ExtraMonthlyPayment
				rolloverPayment := 0.0
				oneTimeImmediatePayment := test.model.Input.OneTimeImmediatePayment
				compoundMinimumPayments := 0.0
				maxMonth := 0.0
				recalculate := false
				finishedDebts := []calculator.Debt{}

				// custom insertion sort to avoid importing the sort library
				// handrolling our own saves 8kb in the binary file
				debts := func(arr []calculator.Debt) []calculator.Debt {
					for i := 0; i < len(arr); i++ {
						for j := i; j > 0 && arr[j-1].Amount > arr[j].Amount; j-- {
							arr[j], arr[j-1] = arr[j-1], arr[j]
						}
					}
					return arr
				}(test.model.Input.Debts)

				debtSequences := calculator.DebtSequences{}

				for _, debt := range debts {
					debtBalance := debt.Amount

					debtSequence := calculator.DebtSequence{
						Debt:     debt,
						Months:   []float64{},
						Payments: []float64{},
						Balances: []float64{},
					}

					monthIter := 1.0
					for {
						/*
							This triggers when an arbitrarily high debt payoff happens.
							Without this, the calculator will hang trying to calculate millions of months.
							Example: amount == 100000000 and min payoff == 10
						*/
						if monthIter/12 >= c.MaxYear {
							recalculate = true
							break
						}

						basePayment := debt.MinimumPayment

						if monthIter == maxMonth {
							basePayment = debt.MinimumPayment + rolloverPayment
							rolloverPayment = 0
						} else if monthIter > maxMonth {
							basePayment = debt.MinimumPayment + extraMonthlyPayment + compoundMinimumPayments + rolloverPayment + oneTimeImmediatePayment
							rolloverPayment = 0
						}

						debtSequence.Months = append(debtSequence.Months, monthIter)

						leftover := (debtBalance - basePayment) * -1
						oneTimeImmediatePayment = 0

						debtBalance = debtBalance - basePayment
						if debtBalance <= 0 {
							debtSequence.Balances = append(debtSequence.Balances, 0)
							debtSequence.Payments = append(debtSequence.Payments, basePayment-(debtBalance*-1))
							rolloverPayment = leftover
							compoundMinimumPayments += debt.MinimumPayment
							maxMonth = debtSequence.Months[len(debtSequence.Months)-1]

							finishedDebts = append(finishedDebts, debt)
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

				if recalculate {
					test.model.Input.Debts = finishedDebts
					return c.CalculateSnowball(test.model)
				}

				return debtSequences
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSnowballCalculateAvalanche(t *testing.T) {
	for _, test := range snowballTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.SnowballAvalanche{
				MaxYear: 200,
			}

			// actual := c.CalculateAvalanche(test.model)
			c.CalculateAvalanche(test.model)
			// expected := func() calculator.DebtSequences {
			func() calculator.DebtSequences {
				debtSequences := calculator.DebtSequences{}

				debts := test.model.Input.Debts
				extraMonthlyPayment := test.model.Input.ExtraMonthlyPayment
				rolloverPayment := 0.0
				oneTimeImmediatePayment := test.model.Input.OneTimeImmediatePayment
				compoundMinimumPayments := 0.0
				maxMonth := 0.0

				sort.Slice(debts, func(i, j int) bool {
					return debts[i].AnnualInterest > debts[j].AnnualInterest
				})

				for _, debt := range debts {
					debtBalance := debt.Amount

					debtSequence := calculator.DebtSequence{
						Debt:     debt,
						Months:   []float64{},
						Payments: []float64{},
						Balances: []float64{},
					}

					monthIter := 1.0
					for {
						if monthIter/12 >= c.MaxYear {
							debtSequences = append(debtSequences, debtSequence)
							break
						}

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

			// assert.Equal(t, expected, actual)
		})
	}
}
