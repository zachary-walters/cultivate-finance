package calculator

type SnowballAvalancheCalculation SnowballCalculation

type SnowballAvalanche struct {
	AbstractCalculation

	MaxYear float64
}

func NewSnowballAvalanche() *SnowballAvalanche {
	return &SnowballAvalanche{
		MaxYear: 300,
	}
}

func (c *SnowballAvalanche) CalculateSnowball(model *Model) DebtSequences {
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
			if monthIter/12 >= c.MaxYear && !invalid {
				debtSequence.Invalid = true
				// debtSequences = append(debtSequences, debtSequence)
				invalid = true
				break
			} else if monthIter/12 >= c.MaxYear {
				// debtSequences = append(debtSequences, debtSequence)
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

// func (c *SnowballAvalanche) CalculateAvalanche(model *Model) DebtSequences {
// 	extraMonthlyPayment := model.Input.ExtraMonthlyPayment
// 	rolloverPayment := 0.0
// 	oneTimeImmediatePayment := model.Input.OneTimeImmediatePayment
// 	compoundMinimumPayments := 0.0
// 	maxMonth := 0.0
// 	invalid := false

// 	// custom insertion sort to avoid importing the sort library
// 	// handrolling our own saves 8kb in the binary file
// 	debts := func(arr []Debt) []Debt {
// 		for i := 0; i < len(arr); i++ {
// 			for j := i; j > 0 && arr[j-1].AnnualInterest < arr[j].AnnualInterest; j-- {
// 				arr[j], arr[j-1] = arr[j-1], arr[j]
// 			}
// 		}
// 		return arr
// 	}(model.Input.Debts)

// 	debtSequences := DebtSequences{}

// 	for _, debt := range debts {
// 		debtBalance := debt.Amount

// 		debtSequence := DebtSequence{
// 			Debt:     debt,
// 			Months:   []float64{},
// 			Payments: []float64{},
// 			Balances: []float64{},
// 		}

// 		monthIter := 1.0
// 		for {
// 			/*
// 				This triggers when an arbitrarily high debt payoff happens.
// 				Without this, the calculator will hang trying to calculate millions of months.
// 				Example: amount == 100000000 and min payoff == 10
// 			*/
// 			if monthIter/12 > c.MaxYear && !invalid {
// 				debtSequence.Invalid = true
// 				// debtSequences = append(debtSequences, debtSequence)
// 				invalid = true
// 				break
// 			} else if monthIter/12 > c.MaxYear {
// 				// debtSequences = append(debtSequences, debtSequence)
// 				break
// 			}

// 			basePayment := debt.MinimumPayment

// 			if monthIter == maxMonth {
// 				basePayment = debt.MinimumPayment + rolloverPayment
// 				rolloverPayment = 0
// 			} else if monthIter > maxMonth {
// 				basePayment = debt.MinimumPayment + extraMonthlyPayment + compoundMinimumPayments + rolloverPayment + oneTimeImmediatePayment
// 				rolloverPayment = 0
// 			}

// 			debtSequence.Months = append(debtSequence.Months, monthIter)

// 			leftover := (debtBalance - basePayment) * -1
// 			oneTimeImmediatePayment = 0

// 			debtBalance = debtBalance - basePayment
// 			if debtBalance <= 0 {
// 				debtSequence.Balances = append(debtSequence.Balances, 0)
// 				debtSequence.Payments = append(debtSequence.Payments, basePayment-(debtBalance*-1))
// 				rolloverPayment = leftover
// 				compoundMinimumPayments += debt.MinimumPayment
// 				maxMonth = debtSequence.Months[len(debtSequence.Months)-1]
// 				break
// 			}
// 			debtSequence.Payments = append(debtSequence.Payments, basePayment+oneTimeImmediatePayment)

// 			oneTimeImmediatePayment = 0

// 			// increase debtBalance by it's annual interest
// 			if monthIter != 1 {
// 				debtBalance = debtBalance + (debtBalance * (debt.AnnualInterest / 100 / 12))
// 			}

// 			debtSequence.Balances = append(debtSequence.Balances, debtBalance)
// 			monthIter += 1
// 		}

// 		debtSequences = append(debtSequences, debtSequence)
// 	}

// 	return debtSequences
// }

func (c *SnowballAvalanche) CalculateAvalanche(model *Model) DebtSequences {
	debtSequences := []DebtSequence{}

	// custom insertion sort to avoid importing the sort library
	// handrolling our own saves 8kb in the binary file
	debts := func(arr []Debt) []Debt {
		for i := 0; i < len(arr); i++ {
			for j := i; j > 0 && arr[j-1].AnnualInterest < arr[j].AnnualInterest; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		return arr
	}(model.Input.Debts)

	allDebtPayments := make([][]float64, len(debts))
	allDebtBalances := make([][]float64, len(debts))
	allDebtMonths := make([][]float64, len(debts))

	debtAmounts := make([]float64, len(debts))
	for idx, debt := range debts {
		debtAmounts[idx] = debt.Amount
	}

	oneTimePayment := model.Input.OneTimeImmediatePayment
	month := 1.0
	excessPayment := 0.0
	workingDebtIndex := 0

	finishedDebts := []Debt{}

	recalculate := false

	for {
		if month/12 >= c.MaxYear {
			recalculate = true
			break
		}

		if workingDebtIndex > len(debts) {
			break
		}

		rolloverPayment := 0.0
		next := false
		for idx, debt := range debts {
			if func(debt Debt, debts []Debt) bool {
				for _, d := range debts {
					if d == debt {
						return true
					}
				}

				return false
			}(debt, finishedDebts) {
				continue
			}

			debtPayments := allDebtPayments[idx]
			debtBalances := allDebtBalances[idx]
			debtMonths := allDebtMonths[idx]

			currentBalance := debtAmounts[idx]

			payment := 0.0

			if next {
				payment = debt.MinimumPayment + oneTimePayment + rolloverPayment
				next = false
			} else if workingDebtIndex == idx {
				payment = debt.MinimumPayment + rolloverPayment + excessPayment + model.Input.ExtraMonthlyPayment + oneTimePayment
			} else {
				payment = debt.MinimumPayment
			}

			oneTimePayment = 0.0
			rolloverPayment = 0.0

			if payment >= currentBalance && workingDebtIndex == idx {
				rolloverPayment = (currentBalance - payment) * -1
				debtAmounts[idx] = 0.0
				excessPayment = excessPayment + debt.MinimumPayment

				payment = currentBalance

				workingDebtIndex++
				next = true
				finishedDebts = append(finishedDebts, debt)
			} else if payment >= currentBalance && currentBalance > 0.0 {
				rolloverPayment = (currentBalance - payment) * -1
				debtAmounts[idx] = 0.0
				excessPayment = excessPayment + debt.MinimumPayment

				payment = currentBalance
				finishedDebts = append(finishedDebts, debt)
			} else {
				debtAmounts[idx] = currentBalance - payment
			}

			currentBalance = currentBalance - payment

			if month > 1 {
				currentBalance = currentBalance + (currentBalance * (debt.AnnualInterest / 100 / 12))
			}

			debtAmounts[idx] = currentBalance

			debtPayments = append(debtPayments, payment)
			debtBalances = append(debtBalances, currentBalance)
			debtMonths = append(debtMonths, month)

			allDebtPayments[idx] = debtPayments
			allDebtBalances[idx] = debtBalances
			allDebtMonths[idx] = debtMonths
		}
		month++
	}

	if recalculate {
		// model.Input.Debts = finishedDebts
		// return c.CalculateAvalanche(model)
	}

	for idx := range finishedDebts {
		debtSequences = append(debtSequences, DebtSequence{
			Debt:     debts[idx],
			Months:   allDebtMonths[idx],
			Payments: allDebtPayments[idx],
			Balances: allDebtBalances[idx],
			Invalid:  float64(len(allDebtMonths[idx])+1) >= c.MaxYear*12,
		})
	}

	return debtSequences
}
