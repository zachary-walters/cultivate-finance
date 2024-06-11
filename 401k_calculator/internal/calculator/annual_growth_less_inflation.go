package calculator

type AnnualGrowthLessInflationCalculation Calculation

type AnnualGrowthLessInflation struct{}

func NewAnnualGrowthLessInflation() AnnualGrowthLessInflation {
	return AnnualGrowthLessInflation{}
}

func (c AnnualGrowthLessInflation) CalculateTraditional(model *Model) float64 {
	annualInvestmentGrowth := model.Input.AnnualInvestmentGrowth

	return annualInvestmentGrowth - model.InflationRate
}

func (c AnnualGrowthLessInflation) CalculateTraditionalRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c AnnualGrowthLessInflation) CalculateRoth(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c AnnualGrowthLessInflation) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}
