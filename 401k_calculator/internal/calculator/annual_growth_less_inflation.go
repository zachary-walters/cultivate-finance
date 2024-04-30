package calculator

type AnnualGrowthLessInflationCalculation Calculation

type AnnualGrowthLessInflation struct{}

func NewAnnualGrowthLessInflation() AnnualGrowthLessInflation {
	return AnnualGrowthLessInflation{}
}

func (c AnnualGrowthLessInflation) Calculate(model Model) float64 {
	annualInvestmentGrowth := model.Input.AnnualInvestmentGrowth

	return annualInvestmentGrowth - 0.03
}

func (c AnnualGrowthLessInflation) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
