package calculator

type TotalDisbursementsAfterTaxCalculation Calculation

type TotalDisbursementsAfterTax struct {
	BalancesTraditionalCalculation
}

func NewTotalDisbursementsAfterTax() TotalDisbursementsAfterTax {
	return TotalDisbursementsAfterTax{
		BalancesTraditionalCalculation: NewBalancesTraditional(),
	}
}

func (c TotalDisbursementsAfterTax) Calculate(model Model) float64 {
	traditionalBalances := c.BalancesTraditionalCalculation.Calculate(model)

	var totalDisbursementsAfterTax float64

	for _, income := range traditionalBalances.AfterTaxIncome {
		totalDisbursementsAfterTax += income
	}

	return totalDisbursementsAfterTax
}

func (c TotalDisbursementsAfterTax) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
