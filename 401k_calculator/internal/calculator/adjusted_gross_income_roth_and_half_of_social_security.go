package calculator

type AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation Calculation

type AdjustedGrossIncomeRothAndHalfOfSocialSecurity struct {
	AdjustedGrossIncomeRothCalculation
	HalfOfSocialSecurityCalculation
}

func NewAdjustedGrossIncomeRothAndHalfOfSocialSecurity() AdjustedGrossIncomeRothAndHalfOfSocialSecurity {
	return AdjustedGrossIncomeRothAndHalfOfSocialSecurity{
		AdjustedGrossIncomeRothCalculation: NewAdjustedGrossIncomeRoth(),
		HalfOfSocialSecurityCalculation:    NewHalfOfSocialSecurity(),
	}
}

func (c AdjustedGrossIncomeRothAndHalfOfSocialSecurity) Calculate(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeRothCalculation.Calculate(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.Calculate(model)

	return adjustedGrossIncomeRoth + halfOfSocialSecurity
}

func (c AdjustedGrossIncomeRothAndHalfOfSocialSecurity) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
