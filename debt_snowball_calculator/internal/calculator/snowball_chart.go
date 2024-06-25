package calculator

import (
	"fmt"
	"sort"
)

type SnowballCalculation interface{}

type Snowball struct {
	AbstractCalculation
}

func NewSnowball() *Snowball {
	return &Snowball{}
}

func (c *Snowball) Calculate(model *Model) {
	debts := model.Input.Debts
	extraMonthlyPayment := model.Input.ExtraMonthlyPayment
	rolloverPayment := 0.0
	oneTimeImmediatePayment := model.Input.OneTimeImmediatePayment
	compoundMinimumPayments := 0.0
	maxMonth := 0

	sort.Slice(debts, func(i, j int) bool {
		return debts[i].Amount < debts[j].Amount
	})

	debtSequences := DebtSequences{}

	for _, debt := range debts {
		debtBalance := debt.Amount
		
		debtSequence := DebtSequence{
			Debt: debt,
			Months: []int{},
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
			if debtBalance <= 0 && monthIter == 1 {
				debtSequence.Balances = append(debtSequence.Balances, 0)
				debtSequence.Payments = append(debtSequence.Payments, debt.Amount)

				oneTimeImmediatePayment = debtBalance * -1
				compoundMinimumPayments += debt.MinimumPayment
				maxMonth = debtSequence.Months[len(debtSequence.Months)-1]
				break
			} 

			// use other payments
			leftover := (debtBalance - basePayment) * -1
			debtBalance = debtBalance - basePayment
			if debtBalance <= 0 {
				debtSequence.Balances = append(debtSequence.Balances, 0)
				debtSequence.Payments = append(debtSequence.Payments, basePayment - (debtBalance * -1))

				rolloverPayment = leftover

				compoundMinimumPayments += debt.MinimumPayment
				maxMonth = debtSequence.Months[len(debtSequence.Months)-1]
				break
			}
			debtSequence.Payments = append(debtSequence.Payments, basePayment + oneTimeImmediatePayment)

			oneTimeImmediatePayment = 0

			// increase debtBalance by it's annual interest
			if monthIter != 1 {
				debtBalance = debtBalance + (debtBalance * (debt.AnnualInterest / 100 / 12))
			}
			
			debtSequence.Balances = append(debtSequence.Balances, debtBalance)
			monthIter+=1
		}

		debtSequences = append(debtSequences, debtSequence)
	}

	fmt.Println(debtSequences)
}
