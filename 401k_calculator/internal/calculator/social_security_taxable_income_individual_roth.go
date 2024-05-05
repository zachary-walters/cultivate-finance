package calculator

type SocialSecurityTaxableIncomeIndividualRothCalculation Calculation

type SocialSecurityTaxableIncomeIndividualRoth struct {
	AdjustedGrossIncomeRothCalculation
}

func NewSocialSecurityTaxableIncomeIndividualRoth() SocialSecurityTaxableIncomeIndividualRoth {
	return SocialSecurityTaxableIncomeIndividualRoth{
		AdjustedGrossIncomeRothCalculation: NewAdjustedGrossIncomeRoth(),
	}
}

func (c SocialSecurityTaxableIncomeIndividualRoth) Calculate(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeRothCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividualRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
