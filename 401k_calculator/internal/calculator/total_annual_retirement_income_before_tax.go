package calculator

type TotalAnnualRetirementIncomeBeforeTaxCalculation Calculation

type TotalAnnualRetirementIncomeBeforeTax struct{}

func NewTotalAnnualRetirementIncomeBeforeTax() TotalAnnualRetirementIncomeBeforeTax {
	return TotalAnnualRetirementIncomeBeforeTax{}
}

func (c TotalAnnualRetirementIncomeBeforeTax) Calculate(model Model) float64 {
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

func (c TotalAnnualRetirementIncomeBeforeTax) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
