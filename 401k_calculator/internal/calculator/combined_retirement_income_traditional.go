package calculator

type CombinedRetirementIncomeTraditionalCalculation Calculation

type CombinedRetirementIncomeTraditional struct {
	AdjustedGrossIncomeTraditionalCalculation
	HalfOfSocialSecurityCalculation
}

func NewCombinedRetirementIncomeTraditional() CombinedRetirementIncomeTraditional {
	return CombinedRetirementIncomeTraditional{
		AdjustedGrossIncomeTraditionalCalculation: NewAdjustedGrossIncomeTraditional(),
		HalfOfSocialSecurityCalculation:           NewHalfOfSocialSecurity(),
	}
}

func (c CombinedRetirementIncomeTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditional := c.AdjustedGrossIncomeTraditionalCalculation.Calculate(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.Calculate(model)

	return adjustedGrossIncomeTraditional + halfOfSocialSecurity
}

func (c CombinedRetirementIncomeTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
