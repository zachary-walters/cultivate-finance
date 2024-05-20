package calculator

type TotalTaxableIncomeCalculation Calculation

type TotalTaxableIncome struct {
	AdjustedGrossIncomeCalculation
	SocialSecurityTaxableIncomeIndividualCalculation
}

func NewTotalTaxableIncome() TotalTaxableIncome {
	return TotalTaxableIncome{
		AdjustedGrossIncomeCalculation:                   NewAdjustedGrossIncome(),
		SocialSecurityTaxableIncomeIndividualCalculation: NewSocialSecurityTaxableIncomeIndividual(),
	}
}

func (c TotalTaxableIncome) CalculateTraditional(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateTraditional(model)
	socialSecurityTaxableIncomeIndividual := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateTraditional(model)

	return adjustedGrossIncome + socialSecurityTaxableIncomeIndividual
}

func (c TotalTaxableIncome) CalculateTraditionalRetirement(model *Model) float64 {
	return c.CalculateTraditional(model)
}

func (c TotalTaxableIncome) CalculateRoth(model *Model) float64 {
	adjustedGrossIncome := c.AdjustedGrossIncomeCalculation.CalculateRoth(model)
	socialSecurityTaxableIncomeIndividual := c.SocialSecurityTaxableIncomeIndividualCalculation.CalculateRoth(model)

	return adjustedGrossIncome + socialSecurityTaxableIncomeIndividual
}

func (c TotalTaxableIncome) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateRoth(model)
}
