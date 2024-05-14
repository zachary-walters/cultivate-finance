package calculator

type SocialSecurityTaxableIncomeCalculation Calculation

type SocialSecurityTaxableIncome struct {
	SocialSecurityTaxableIncomeIndividualCalculation
	SocialSecurityTaxableIncomeJointCalculation
}

func NewSocialSecurityTaxableIncome() SocialSecurityTaxableIncome {
	return SocialSecurityTaxableIncome{
		SocialSecurityTaxableIncomeIndividualCalculation: NewSocialSecurityTaxableIncomeIndividual(),
		SocialSecurityTaxableIncomeJointCalculation:      NewSocialSecurityTaxableIncomeJoint(),
	}
}

func (c SocialSecurityTaxableIncome) CalculateTraditional(model Model) float64 {
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateTraditional(model)
	socialSecurityTaxableIncomeJointRoth := c.SocialSecurityTaxableIncomeJointCalculation.CalculateTraditional(model)

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

func (c SocialSecurityTaxableIncome) CalculateTraditionalRetirement(model Model) float64 {
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateTraditionalRetirement(model)
	socialSecurityTaxableIncomeJointRoth := c.SocialSecurityTaxableIncomeJointCalculation.CalculateTraditionalRetirement(model)

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

func (c SocialSecurityTaxableIncome) CalculateRoth(model Model) float64 {
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateRoth(model)
	socialSecurityTaxableIncomeJointRoth := c.SocialSecurityTaxableIncomeJointCalculation.CalculateRoth(model)

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

func (c SocialSecurityTaxableIncome) CalculateRothRetirement(model Model) float64 {
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateRothRetirement(model)
	socialSecurityTaxableIncomeJointRoth := c.SocialSecurityTaxableIncomeJointCalculation.CalculateRothRetirement(model)

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
