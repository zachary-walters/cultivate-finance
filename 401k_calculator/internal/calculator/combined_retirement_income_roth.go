package calculator

type CombinedRetirementIncomeRothCalculation Calculation

type CombinedRetirementIncomeRoth struct {
	AdjustedGrossIncomeRothCalculation
	HalfOfSocialSecurityCalculation
}

func NewCombinedRetirementIncomeRoth() CombinedRetirementIncomeRoth {
	return CombinedRetirementIncomeRoth{
		AdjustedGrossIncomeRothCalculation: NewAdjustedGrossIncomeRoth(),
		HalfOfSocialSecurityCalculation:    NewHalfOfSocialSecurity(),
	}
}

func (c CombinedRetirementIncomeRoth) Calculate(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeRothCalculation.Calculate(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.Calculate(model)

	return adjustedGrossIncomeRoth + halfOfSocialSecurity
}

func (c CombinedRetirementIncomeRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
