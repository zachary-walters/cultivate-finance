package calculator

type SocialSecurityTaxableIncomeIndividualRothCalculation Calculation

type SocialSecurityTaxableIncomeIndividualRoth struct {
	AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation
}

func NewSocialSecurityTaxableIncomeIndividualRoth() SocialSecurityTaxableIncomeIndividualRoth {
	return SocialSecurityTaxableIncomeIndividualRoth{
		AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation: NewAdjustedGrossIncomeRothAndHalfOfSocialSecurity(),
	}
}

func (c SocialSecurityTaxableIncomeIndividualRoth) Calculate(model Model) float64 {
	adjustedGrossIncomeRothAndHalfOfSocialSecurity := c.AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncomeRothAndHalfOfSocialSecurity {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividualRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
