package calculator

type SocialSecurityTaxableIncomeRothCalculation Calculation

type SocialSecurityTaxableIncomeRoth struct {
	SocialSecurityTaxableIncomeIndividualRothCalculation
	SocialSecurityTaxableIncomeJointRothCalculation
}

func NewSocialSecurityTaxableIncomeRoth() SocialSecurityTaxableIncomeRoth {
	return SocialSecurityTaxableIncomeRoth{
		SocialSecurityTaxableIncomeIndividualRothCalculation: NewSocialSecurityTaxableIncomeIndividualRoth(),
		SocialSecurityTaxableIncomeJointRothCalculation:      NewSocialSecurityTaxableIncomeJointRoth(),
	}
}

func (c SocialSecurityTaxableIncomeRoth) Calculate(model Model) float64 {
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualRothCalculation.Calculate(model)
	socialSecurityTaxableIncomeJointRoth := c.SocialSecurityTaxableIncomeJointRothCalculation.Calculate(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return socialSecurityTaxableIncomeIndividualRoth
	case "married-seperate":
		return socialSecurityTaxableIncomeIndividualRoth
	case "married-joint":
		return socialSecurityTaxableIncomeJointRoth
	case "head-of-household":
		return socialSecurityTaxableIncomeJointRoth
	default:
		return 0
	}
}

func (c SocialSecurityTaxableIncomeRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
