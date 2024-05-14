package calculator

type CombinedRetirementIncomeCalculation Calculation

type CombinedRetirementIncome struct {
	AdjustedGrossIncomeCalculation
	HalfOfSocialSecurityCalculation
}

func NewCombinedRetirementIncome() CombinedRetirementIncome {
	return CombinedRetirementIncome{
		AdjustedGrossIncomeCalculation:  NewAdjustedGrossIncome(),
		HalfOfSocialSecurityCalculation: NewHalfOfSocialSecurity(),
	}
}

func (c CombinedRetirementIncome) CalculateTraditional(model Model) float64 {
	return 0
}

func (c CombinedRetirementIncome) CalculateTraditionalRetirement(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateTraditionalRetirement(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.CalculateTraditionalRetirement(model)

	return adjustedGrossIncomeRoth + halfOfSocialSecurity
}

func (c CombinedRetirementIncome) CalculateRoth(model Model) float64 {
	return 0
}

func (c CombinedRetirementIncome) CalculateRothRetirement(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateRothRetirement(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.CalculateRothRetirement(model)

	return adjustedGrossIncomeRoth + halfOfSocialSecurity
}
