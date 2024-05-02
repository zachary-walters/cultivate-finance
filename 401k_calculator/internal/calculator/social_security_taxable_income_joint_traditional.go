package calculator

type SocialSecurityTaxableIncomeJointTraditionalCalculation Calculation

type SocialSecurityTaxableIncomeJointTraditional struct {
	AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation
}

func NewSocialSecurityTaxableIncomeJointTraditional() SocialSecurityTaxableIncomeJointTraditional {
	return SocialSecurityTaxableIncomeJointTraditional{
		AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation: NewAdjustedGrossIncomeTraditionalAndHalfOfSocialSecurity(),
	}
}

func (c SocialSecurityTaxableIncomeJointTraditional) Calculate(model Model) float64 {
	adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity := c.AdjustedGrossIncomeTraditionalAndHalfOfSocialSecurityCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > adjustedGrossIncomeTraditionalAndHalfOfSocialSecurity {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJointTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
