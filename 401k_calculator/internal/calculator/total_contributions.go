package calculator

type TotalContributionsCalculation Calculation

type TotalContributions struct {
	BalancesTraditionalCalculation
	BalancesRothMatchingNetContributionsCalculation
}

func NewTotalContributions() TotalContributions {
	return TotalContributions{
		BalancesTraditionalCalculation:                  NewBalancesTraditional(),
		BalancesRothMatchingNetContributionsCalculation: NewBalancesRothMatchingNetContributions(),
	}
}

func (c TotalContributions) CalculateTraditional(model *Model) float64 {
	balancesTraditional := c.BalancesTraditionalCalculation.Calculate(model)

	totalContributions := 0.0
	for _, contribution := range balancesTraditional.Contribution {
		totalContributions += contribution
	}

	return totalContributions
}

func (c TotalContributions) CalculateTraditionalRetirement(model *Model) float64 {
	return 0.0
}

func (c TotalContributions) CalculateRoth(model *Model) float64 {
	balancesRothMatchingNetContributions := c.BalancesRothMatchingNetContributionsCalculation.Calculate(model)

	totalContributions := 0.0
	for _, contribution := range balancesRothMatchingNetContributions.Contribution {
		totalContributions += contribution
	}

	return totalContributions
}

func (c TotalContributions) CalculateRothRetirement(model *Model) float64 {
	return 0.0
}
