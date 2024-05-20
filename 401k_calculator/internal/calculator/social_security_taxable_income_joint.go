package calculator

type SocialSecurityTaxableIncomeJointCalculation Calculation

type SocialSecurityTaxableIncomeJoint struct {
	AdjustedGrossIncomeCalculation
}

func NewSocialSecurityTaxableIncomeJoint() SocialSecurityTaxableIncomeJoint {
	return SocialSecurityTaxableIncomeJoint{
		AdjustedGrossIncomeCalculation: NewAdjustedGrossIncome(),
	}
}

func (c SocialSecurityTaxableIncomeJoint) CalculateTraditional(model *Model) float64 {
	totalTaxableIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateTraditional(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > totalTaxableIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJoint) CalculateTraditionalRetirement(model *Model) float64 {
	totalTaxableIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateTraditionalRetirement(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > totalTaxableIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJoint) CalculateRoth(model *Model) float64 {
	totalTaxableIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateRoth(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > totalTaxableIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJoint) CalculateRothRetirement(model *Model) float64 {
	totalTaxableIncomeRoth := c.AdjustedGrossIncomeCalculation.CalculateRothRetirement(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > totalTaxableIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}
