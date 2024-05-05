package calculator

type SocialSecurityTaxableIncomeJointRothCalculation Calculation

type SocialSecurityTaxableIncomeJointRoth struct {
	AdjustedGrossIncomeRothCalculation
}

func NewSocialSecurityTaxableIncomeJointRoth() SocialSecurityTaxableIncomeJointRoth {
	return SocialSecurityTaxableIncomeJointRoth{
		AdjustedGrossIncomeRothCalculation: NewAdjustedGrossIncomeRoth(),
	}
}

func (c SocialSecurityTaxableIncomeJointRoth) Calculate(model Model) float64 {
	totalTaxableIncomeRoth := c.AdjustedGrossIncomeRothCalculation.Calculate(model)

	for _, taxRate := range model.SocialSecurityTaxRatesJoint {
		if taxRate.Cap > totalTaxableIncomeRoth {
			return model.Input.SocialSecurity * taxRate.Rate
		}
	}

	return model.SocialSecurityTaxRatesJoint[len(model.SocialSecurityTaxRatesJoint)-1].Rate * model.Input.SocialSecurity
}

func (c SocialSecurityTaxableIncomeJointRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
