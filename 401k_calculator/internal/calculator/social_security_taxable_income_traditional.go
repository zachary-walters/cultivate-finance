package calculator

type SocialSecurityTaxableIncomeTraditionalCalculation Calculation

type SocialSecurityTaxableIncomeTraditional struct {
	SocialSecurityTaxableIncomeIndividualTraditionalCalculation
	SocialSecurityTaxableIncomeJointTraditionalCalculation
}

func NewSocialSecurityTaxableIncomeTraditional() SocialSecurityTaxableIncomeTraditional {
	return SocialSecurityTaxableIncomeTraditional{
		SocialSecurityTaxableIncomeIndividualTraditionalCalculation: NewSocialSecurityTaxableIncomeIndividualTraditional(),
		SocialSecurityTaxableIncomeJointTraditionalCalculation:      NewSocialSecurityTaxableIncomeJointTraditional(),
	}
}

func (c SocialSecurityTaxableIncomeTraditional) Calculate(model Model) float64 {
	socialSecurityTaxableIncomeIndividualTraditional := c.SocialSecurityTaxableIncomeIndividualTraditionalCalculation.Calculate(model)
	socialSecurityTaxableIncomeJointTraditional := c.SocialSecurityTaxableIncomeJointTraditionalCalculation.Calculate(model)

	switch model.Input.RetirementFilingStatus {
	case "single":
		return socialSecurityTaxableIncomeIndividualTraditional
	case "married-seperate":
		return socialSecurityTaxableIncomeIndividualTraditional
	case "married-joint":
		return socialSecurityTaxableIncomeJointTraditional
	case "head-of-household":
		return socialSecurityTaxableIncomeJointTraditional
	default:
		return 0
	}
}

func (c SocialSecurityTaxableIncomeTraditional) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
