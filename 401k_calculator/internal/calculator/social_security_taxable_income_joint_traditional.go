package calculator

type SocialSecurityTaxableIncomeJointTraditionalCalculation Calculation

type SocialSecurityTaxableIncomeJointTraditional struct {
	AdjustedGrossIncomeTraditionalCalculation
}

func NewSocialSecurityTaxableIncomeJointTraditional() SocialSecurityTaxableIncomeJointTraditional {
	return SocialSecurityTaxableIncomeJointTraditional{
		AdjustedGrossIncomeTraditionalCalculation: NewAdjustedGrossIncomeTraditional(),
	}
}

func (c SocialSecurityTaxableIncomeJointTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditional := c.AdjustedGrossIncomeTraditionalCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > adjustedGrossIncomeTraditional {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJointTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
