package calculator

type TotalDisbursementsRothMatchingGrossCalculation Calculation

type TotalDisbursementsRothMatchingGross struct {
	BalancesRothMatchingGrossContributionsCalculation
}

func NewTotalDisbursementsRothMatchingGross() TotalDisbursementsRothMatchingGross {
	return TotalDisbursementsRothMatchingGross{
		BalancesRothMatchingGrossContributionsCalculation: NewBalancesRothMatchingGrossContributions(),
	}
}

func (c TotalDisbursementsRothMatchingGross) Calculate(model Model) float64 {
	traditionalBalances := c.BalancesRothMatchingGrossContributionsCalculation.Calculate(model)

	var totalDispersementsRothMatchingNetContributions float64

	for _, income := range traditionalBalances.Withdrawal {
		totalDispersementsRothMatchingNetContributions += income
	}

	return totalDispersementsRothMatchingNetContributions
}

func (c TotalDisbursementsRothMatchingGross) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
