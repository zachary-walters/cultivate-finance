package calculator

type TotalInterestRothCalculation Calculation

type TotalInterestRoth struct {
	TotalDisbursementsRothMatchingNetCalculation
	TotalContributionsRothCalculation
}

func NewTotalInterestRoth() TotalInterestRoth {
	return TotalInterestRoth{
		TotalDisbursementsRothMatchingNetCalculation: NewTotalDisbursementsRothMatchingNet(),
		TotalContributionsRothCalculation:            NewTotalContributionsRoth(),
	}
}

func (c TotalInterestRoth) Calculate(model Model) float64 {
	totalDisbursementsRothMatchingNetCalculation := c.TotalDisbursementsRothMatchingNetCalculation.Calculate(model)
	totalContributionsRoth := c.TotalContributionsRothCalculation.Calculate(model)

	return totalDisbursementsRothMatchingNetCalculation - totalContributionsRoth
}

func (c TotalInterestRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
