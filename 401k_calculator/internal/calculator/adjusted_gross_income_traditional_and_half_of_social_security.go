package calculator

type AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation Calculation

type AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity struct {
	AdjustedGrossIncomeTraditionalCalculation
	HalfOfSocialSecurityCalculation
}

func NewAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity() AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity {
	return AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity{
		AdjustedGrossIncomeTraditionalCalculation: NewAdjustedGrossIncomeTraditional(),
		HalfOfSocialSecurityCalculation:           NewHalfOfSocialSecurity(),
	}
}

func (c AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditional := c.AdjustedGrossIncomeTraditionalCalculation.Calculate(model)
	halfOfSocialSecurity := c.HalfOfSocialSecurityCalculation.Calculate(model)

	return adjustedGrossIncomeTraditional + halfOfSocialSecurity
}

func (c AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
