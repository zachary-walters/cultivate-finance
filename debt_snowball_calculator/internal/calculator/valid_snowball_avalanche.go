package calculator

type ValidSnowballAvalancheCalculation SnowballCalculation

type ValidSnowballAvalanche struct {
	SnowballAvalancheCalculation
	ValidDebtsCalculation
}

func NewValidSnowballAvalanche() *ValidSnowballAvalanche {
	return &ValidSnowballAvalanche{
		SnowballAvalancheCalculation: NewSnowballAvalanche(),
		ValidDebtsCalculation:        NewValidDebts(),
	}
}

func (c *ValidSnowballAvalanche) CalculateSnowball(model *Model) DebtSequences {
	snowballAvalanche := c.SnowballAvalancheCalculation.CalculateSnowball(model)
	validDebts := c.ValidDebtsCalculation.CalculateSnowball(model)

	debtSequences := DebtSequences{}

	for _, debtSequence := range snowballAvalanche {
		for _, debt := range validDebts {
			if debt == debtSequence.Debt {
				debtSequences = append(debtSequences, debtSequence)
			}
		}
	}

	return debtSequences
}

func (c *ValidSnowballAvalanche) CalculateAvalanche(model *Model) DebtSequences {
	snowballAvalanche := c.SnowballAvalancheCalculation.CalculateAvalanche(model)
	validDebts := c.ValidDebtsCalculation.CalculateAvalanche(model)

	debtSequences := DebtSequences{}

	for _, debtSequence := range snowballAvalanche {
		for _, debt := range validDebts {
			if debt == debtSequence.Debt {
				debtSequences = append(debtSequences, debtSequence)
			}
		}
	}

	return debtSequences
}
