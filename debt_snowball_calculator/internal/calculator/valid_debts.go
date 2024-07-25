package calculator

type ValidDebtsCalculation DebtCalculation
type ValidDebts struct {
	SnowballAvalancheCalculation
}

func NewValidDebts() *ValidDebts {
	return &ValidDebts{
		SnowballAvalancheCalculation: NewSnowballAvalanche(),
	}
}

func (c *ValidDebts) CalculateSnowball(model Model) []Debt {
	snowball := c.SnowballAvalancheCalculation.CalculateSnowball(model)
	avalanche := c.SnowballAvalancheCalculation.CalculateAvalanche(model)

	snowballDebts := []Debt{}

	validDebts := []Debt{}

	for _, debtSequence := range snowball {
		snowballDebts = append(snowballDebts, debtSequence.Debt)
	}

	for _, avalancheSequence := range avalanche {
		for _, debt := range snowballDebts {
			if avalancheSequence.Debt == debt {
				validDebts = append(validDebts, debt)
			}
		}
	}

	return validDebts
}

func (c *ValidDebts) CalculateAvalanche(model Model) []Debt {
	return c.CalculateSnowball(model)
}
