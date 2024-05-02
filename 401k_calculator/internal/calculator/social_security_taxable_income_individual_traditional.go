package calculator

type SocialSecurityTaxableIncomeIndividualTraditionalCalculation Calculation

type SocialSecurityTaxableIncomeIndividualTraditional struct {
	AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation
}

func NewSocialSecurityTaxableIncomeIndividualTraditional() SocialSecurityTaxableIncomeIndividualTraditional {
	return SocialSecurityTaxableIncomeIndividualTraditional{
		AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation: NewAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity(),
	}
}

func (c SocialSecurityTaxableIncomeIndividualTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity := c.AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividualTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
