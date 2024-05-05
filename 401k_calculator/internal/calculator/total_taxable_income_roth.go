package calculator

type TotalTaxableIncomeRothCalculation Calculation

type TotalTaxableIncomeRoth struct {
	AdjustedGrossIncomeRothCalculation
	SocialSecurityTaxableIncomeIndividualRothCalculation
}

func NewTotalTaxableIncomeRoth() TotalTaxableIncomeRoth {
	return TotalTaxableIncomeRoth{
		AdjustedGrossIncomeRothCalculation:                   NewAdjustedGrossIncomeRoth(),
		SocialSecurityTaxableIncomeIndividualRothCalculation: NewSocialSecurityTaxableIncomeIndividualRoth(),
	}
}

func (c TotalTaxableIncomeRoth) Calculate(model Model) float64 {
	adjustedGrossIncomeRoth := c.AdjustedGrossIncomeRothCalculation.Calculate(model)
	socialSecurityTaxableIncomeIndividualRoth := c.SocialSecurityTaxableIncomeIndividualRothCalculation.Calculate(model)

	return adjustedGrossIncomeRoth + socialSecurityTaxableIncomeIndividualRoth
}

func (c TotalTaxableIncomeRoth) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
