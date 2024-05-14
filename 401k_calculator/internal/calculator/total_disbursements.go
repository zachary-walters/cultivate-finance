package calculator

type TotalDisbursementsCalculation Calculation

type TotalDisbursements struct {
	BalancesTraditionalCalculation
	BalancesRothMatchingNetContributionsCalculation
}

func NewTotalDisbursements() TotalDisbursements {
	return TotalDisbursements{
		BalancesTraditionalCalculation:                  NewBalancesTraditional(),
		BalancesRothMatchingNetContributionsCalculation: NewBalancesRothMatchingNetContributions(),
	}
}

func (c TotalDisbursements) CalculateTraditional(model Model) float64 {
	return 0
}

func (c TotalDisbursements) CalculateTraditionalRetirement(model Model) float64 {
	traditionalBalances := c.BalancesTraditionalCalculation.Calculate(model)

	var totalDisbursementsTraditional float64

	for _, income := range traditionalBalances.AfterTaxIncome {
		totalDisbursementsTraditional += income
	}

	return totalDisbursementsTraditional
}

func (c TotalDisbursements) CalculateRoth(model Model) float64 {
	return c.CalculateRothRetirement(model)
}

func (c TotalDisbursements) CalculateRothRetirement(model Model) float64 {
	rothBalances := c.BalancesRothMatchingNetContributionsCalculation.Calculate(model)

	var totalDisbursementsRoth float64

	for _, income := range rothBalances.AfterTaxIncome {
		totalDisbursementsRoth += income
	}

	return totalDisbursementsRoth
}
