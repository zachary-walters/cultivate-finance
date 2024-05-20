package calculator

type SocialSecurityTaxableIncomeIndividualCalculation Calculation

type SocialSecurityTaxableIncomeIndividual struct {
	AdjustedGrossIncomeCalculation
}

func NewSocialSecurityTaxableIncomeIndividual() SocialSecurityTaxableIncomeIndividual {
	return SocialSecurityTaxableIncomeIndividual{
		AdjustedGrossIncomeCalculation: NewAdjustedGrossIncome(),
	}
}

func (c SocialSecurityTaxableIncomeIndividual) CalculateTraditional(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateTraditional(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncome {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividual) CalculateTraditionalRetirement(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateTraditionalRetirement(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncome {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividual) CalculateRoth(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateRoth(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncome {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeIndividual) CalculateRetirement(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateRothRetirement(model)

	for _, taxRate := range model.SocialSecurityTaxRatesIndividual {
		if taxRate.Cap > adjustedGrossIncome {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesIndividual[len(model.SocialSecurityTaxRatesIndividual)-1].Rate * model.Input.SocialSecurity
}
