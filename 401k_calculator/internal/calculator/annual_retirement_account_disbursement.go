package calculator

type AnnualRetirementAccountDisbursementCalculation Calculation

type AnnualRetirementAccountDisbursement struct {
	EffectiveTaxRateOnGrossCalculation
	TaxOnTraditionalIRAWithdrawalCalculation
}

func NewAnnualRetirementAccountDisbursement() AnnualRetirementAccountDisbursement {
	return AnnualRetirementAccountDisbursement{
		EffectiveTaxRateOnGrossCalculation:       NewEffectiveTaxRateOnGross(),
		TaxOnTraditionalIRAWithdrawalCalculation: NewTaxOnTraditionalIRAWithdrawal(),
	}
}

func (c AnnualRetirementAccountDisbursement) CalculateTraditional(model Model) float64 {
	return 0
}

func (c AnnualRetirementAccountDisbursement) CalculateTraditionalRetirement(model Model) float64 {
	taxOnTraditionalIRAWithdrawal := c.TaxOnTraditionalIRAWithdrawalCalculation.CalculateTraditionalRetirement(model)

	return model.Input.YearlyWithdrawal - taxOnTraditionalIRAWithdrawal
}

func (c AnnualRetirementAccountDisbursement) CalculateRoth(model Model) float64 {
	return 0
}

func (c AnnualRetirementAccountDisbursement) CalculateRothRetirement(model Model) float64 {
	effectiveTaxRateOnGross := c.EffectiveTaxRateOnGrossCalculation.CalculateRothRetirement(model)
	taxOnTraditionalIRAWithdrawal := c.TaxOnTraditionalIRAWithdrawalCalculation.CalculateRothRetirement(model)

	return (model.Input.YearlyWithdrawal - taxOnTraditionalIRAWithdrawal) * (1 - effectiveTaxRateOnGross)
}
