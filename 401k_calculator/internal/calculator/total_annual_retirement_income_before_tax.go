package calculator

type TotalAnnualRetirementIncomeBeforeTaxCalculation Calculation

type TotalAnnualRetirementIncomeBeforeTax struct{}

func NewTotalAnnualRetirementIncomeBeforeTax() TotalAnnualRetirementIncomeBeforeTax {
	return TotalAnnualRetirementIncomeBeforeTax{}
}

func (c TotalAnnualRetirementIncomeBeforeTax) CalculateTraditional(model *Model) float64 {
	return 0.0
}

func (c TotalAnnualRetirementIncomeBeforeTax) CalculateTraditionalRetirement(model *Model) float64 {
	return model.Input.AnnuityIncome +
		model.Input.OtherLongTermCapitalGains +
		model.Input.OtherTaxableIncome +
		model.Input.PensionIncome +
		model.Input.QualifiedDividends +
		model.Input.RentalNetIncome +
		model.Input.SocialSecurity +
		model.Input.WorkIncome +
		model.Input.YearlyWithdrawal
}

func (c TotalAnnualRetirementIncomeBeforeTax) CalculateRoth(model *Model) float64 {
	return 0.0
}

func (c TotalAnnualRetirementIncomeBeforeTax) CalculateRothRetirement(model *Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}
