package calculator

type TotalDisbursementsRothMatchingNetCalculation Calculation

type TotalDisbursementsRothMatchingNet struct {
	BalancesRothMatchingNetContributionsCalculation
}

func NewTotalDisbursementsRothMatchingNet() TotalDisbursementsRothMatchingNet {
	return TotalDisbursementsRothMatchingNet{
		BalancesRothMatchingNetContributionsCalculation: NewBalancesRothMatchingNetContributions(),
	}
}

func (c TotalDisbursementsRothMatchingNet) Calculate(model Model) float64 {
	traditionalBalances := c.BalancesRothMatchingNetContributionsCalculation.Calculate(model)

	var totalDispersementsRothMatchingNetContributions float64

	for _, income := range traditionalBalances.Withdrawal {
		totalDispersementsRothMatchingNetContributions += income
	}

	return totalDispersementsRothMatchingNetContributions
}

func (c TotalDisbursementsRothMatchingNet) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
