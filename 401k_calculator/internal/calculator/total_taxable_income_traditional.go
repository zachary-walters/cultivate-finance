package calculator

type TotalTaxableIncomeTraditionalCalculation Calculation

type TotalTaxableIncomeTraditional struct {
	AdjustedGrossIncomeTraditionalCalculation
	SocialSecurityTaxableIncomeIndividualTraditionalCalculation
}

func NewTotalTaxableIncomeTraditional() TotalTaxableIncomeTraditional {
	return TotalTaxableIncomeTraditional{
		AdjustedGrossIncomeTraditionalCalculation:                   NewAdjustedGrossIncomeTraditional(),
		SocialSecurityTaxableIncomeIndividualTraditionalCalculation: NewSocialSecurityTaxableIncomeIndividualTraditional(),
	}
}

func (c TotalTaxableIncomeTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditional := c.AdjustedGrossIncomeTraditionalCalculation.Calculate(model)
	socialSecurityTaxableIncomeIndividualTraditionalCalculation := c.SocialSecurityTaxableIncomeIndividualTraditionalCalculation.Calculate(model)

	return adjustedGrossIncomeTraditional + socialSecurityTaxableIncomeIndividualTraditionalCalculation
}

func (c TotalTaxableIncomeTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
