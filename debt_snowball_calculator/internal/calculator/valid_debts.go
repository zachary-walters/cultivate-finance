package calculator

type ValidDebtsCalculation interface {
	CalculateSnowball(*Model) []Debt
	CalculateAvalanche(*Model) []Debt
}

type ValidDebts struct {
	SnowballAvalancheCalculation
}

func NewValidDebts() *ValidDebts {
	return &ValidDebts{
		SnowballAvalancheCalculation: NewSnowballAvalanche(),
	}
}

func (c *ValidDebts) CalculateSnowball(model *Model) []Debt {
	snowball := c.SnowballAvalancheCalculation.CalculateSnowball(model)
	avalanche := c.SnowballAvalancheCalculation.CalculateAvalanche(model)

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

func (c *ValidDebts) CalculateAvalanche(model *Model) []Debt {
	return c.CalculateSnowball(model)
}
