package calculator

type SnowballCalculation interface {
	Calculate(model *Model) DebtSequences
}

type Snowball struct {
	AbstractCalculation

	MaxYear float64
}

func NewSnowball() *Snowball {
	return &Snowball{
		MaxYear: 1000,
	}
}

func (c *Snowball) Calculate(model *Model) DebtSequences {
	extraMonthlyPayment := model.Input.ExtraMonthlyPayment
	rolloverPayment := 0.0
	oneTimeImmediatePayment := model.Input.OneTimeImmediatePayment
	compoundMinimumPayments := 0.0
	maxMonth := 0.0
	invalid := false

	// custom insertion sort to avoid importing the sort library
	// handrolling our own saves 8kb in the binary file
	debts := func(arr []Debt) []Debt {
		for i := 0; i < len(arr); i++ {
			for j := i; j > 0 && arr[j-1].Amount > arr[j].Amount; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		return arr
	}(model.Input.Debts)

	debtSequences := DebtSequences{}

	for _, debt := range debts {
		debtBalance := debt.Amount

		debtSequence := DebtSequence{
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
			if monthIter/12 >= c.MaxYear || invalid {
				debtSequence.Invalid = true
				debtSequence.MaxYear = c.MaxYear
				invalid = true

				debtSequences = append(debtSequences, debtSequence)
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
}
