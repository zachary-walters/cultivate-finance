package calculator

type TotalContributionsRothCalculation Calculation

type TotalContributionsRoth struct {
	BalancesRothMatchingNetContributionsCalculation
}

func NewTotalContributionsRoth() TotalContributionsRoth {
	return TotalContributionsRoth{
		BalancesRothMatchingNetContributionsCalculation: NewBalancesRothMatchingNetContributions(),
	}
}

func (c TotalContributionsRoth) Calculate(model Model) float64 {
	balancesRothMatchingNetContributions := c.BalancesRothMatchingNetContributionsCalculation.Calculate(model)

	totalContributions := 0.0
	for _, contribution := range balancesRothMatchingNetContributions.Contribution {
		totalContributions += contribution
	}

	return totalContributions
}

func (c TotalContributionsRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
