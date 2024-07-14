package calculator

type ValidDebtsCalculation interface {
	Calculate(*Model) []Debt
}

type ValidDebts struct {
	SnowballCalculation
}

func NewValidDebts() *ValidDebts {
	return &ValidDebts{
		SnowballCalculation: NewSnowball(),
	}
}

func (c ValidDebts) Calculate(model *Model) []Debt {
	snowball := c.SnowballCalculation.CalculateSnowball(model)
	avalanche := c.SnowballCalculation.CalculateAvalanche(model)

	validDebts := []Debt{}
	invalidDebts := []Debt{}

	for _, debtSequence := range snowball {
		if debtSequence.Invalid {
			invalidDebts = append(invalidDebts, debtSequence.Debt)
		}
	}

	for _, debtSequence := range avalanche {
		if debtSequence.Invalid {
			invalidDebts = append(invalidDebts, debtSequence.Debt)
		}
	}

	for _, debt := range model.Input.Debts {
		if func(s []Debt, d Debt) bool {
			for _, a := range s {
				if a == d {
					return true
				}
			}
			return false
		}(invalidDebts, debt) {
			continue
		}

		validDebts = append(validDebts, debt)
	}

	return validDebts
}
