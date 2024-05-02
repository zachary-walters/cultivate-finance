package calculator

type SocialSecurityTaxableIncomeJointRothCalculation Calculation

type SocialSecurityTaxableIncomeJointRoth struct {
	AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation
}

func NewSocialSecurityTaxableIncomeJointRoth() SocialSecurityTaxableIncomeJointRoth {
	return SocialSecurityTaxableIncomeJointRoth{
		AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation: NewAdjustedGrossIncomeRothAndHalfOfSocialSecurity(),
	}
}

func (c SocialSecurityTaxableIncomeJointRoth) Calculate(model Model) float64 {
	adjustedGrossIncomeRothAndHalfOfSocialSecurity := c.AdjustedGrossIncomeRothAndHalfOfSocialSecurityCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > adjustedGrossIncomeRothAndHalfOfSocialSecurity {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJointRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
