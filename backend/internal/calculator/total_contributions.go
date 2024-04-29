package calculator

type TotalContributionsCalculation Calculation

type TotalContributions struct {
	BalancesTraditionalCalculation
}

func NewTotalContributions() TotalContributions {
	return TotalContributions{
		BalancesTraditionalCalculation: NewBalancesTraditional(),
	}
}

func (c TotalContributions) Calculate(model Model) float64 {
	balancesTraditional := c.BalancesTraditionalCalculation.Calculate(model)

	totalContributions := 0.0
	for _, contribution := range balancesTraditional.Contribution {
		totalContributions += contribution
	}

	return totalContributions
}

func (c TotalContributions) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
