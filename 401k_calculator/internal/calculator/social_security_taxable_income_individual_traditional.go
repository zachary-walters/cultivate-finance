package calculator

type SocialSecurityTaxableIncomeIndividualTraditionalCalculation Calculation

type SocialSecurityTaxableIncomeIndividualTraditional struct {
	AdjustedGrossIncomeTraditionalCalculation
}

func NewSocialSecurityTaxableIncomeIndividualTraditional() SocialSecurityTaxableIncomeIndividualTraditional {
	return SocialSecurityTaxableIncomeIndividualTraditional{
		AdjustedGrossIncomeTraditionalCalculation: NewAdjustedGrossIncomeTraditional(),
	}
}

func (c SocialSecurityTaxableIncomeIndividualTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditional := c.AdjustedGrossIncomeTraditionalCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncomeTraditional {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividualTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
