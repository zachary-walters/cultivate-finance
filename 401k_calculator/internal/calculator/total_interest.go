package calculator

type TotalInterestCalculation Calculation

type TotalInterest struct {
	TotalDisbursementsCalculation
	TotalContributionsCalculation
}

func NewTotalInterest() TotalInterest {
	return TotalInterest{
		TotalDisbursementsCalculation: NewTotalDisbursements(),
		TotalContributionsCalculation: NewTotalContributions(),
	}
}

func (c TotalInterest) CalculateTraditional(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}

func (c TotalInterest) CalculateTraditionalRetirement(model *Model) float64 {
	totalDisbursements := c.TotalDisbursementsCalculation.CalculateTraditionalRetirement(model)
	totalContributions := c.TotalContributionsCalculation.CalculateTraditionalRetirement(model)

	return totalDisbursements - totalContributions
}

func (c TotalInterest) CalculateRoth(model *Model) float64 {
	return c.CalculateRothRetirement(model)
}

func (c TotalInterest) CalculateRothRetirement(model *Model) float64 {
	totalDisbursements := c.TotalDisbursementsCalculation.CalculateRothRetirement(model)
	totalContributions := c.TotalContributionsCalculation.CalculateRothRetirement(model)

	return totalDisbursements - totalContributions
}
