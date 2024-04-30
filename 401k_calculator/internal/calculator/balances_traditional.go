package calculator

type BalancesTraditionalCalculation ChartCalculation

type BalancesTraditional struct {
	Limit int

	AnnualGrowthLessInflationCalculation
	EffectiveTaxRateOnGrossCalculation
	NetDistributionAfterTaxesCalculation
}

func NewBalancesTraditional() BalancesTraditional {
	return BalancesTraditional{
		Limit:                                133,
		AnnualGrowthLessInflationCalculation: NewAnnualGrowthLessInflation(),
		EffectiveTaxRateOnGrossCalculation:   NewEffectiveTaxRateOnGross(),
		NetDistributionAfterTaxesCalculation: NewNetDistributionAfterTaxes(),
	}
}

func (c BalancesTraditional) Calculate(model Model) ChartData {
	annualGrowthLessInflation := c.AnnualGrowthLessInflationCalculation.CalculateRetirement(model)
	effectiveTaxRateOnGross := c.EffectiveTaxRateOnGrossCalculation.CalculateRetirement(model)
	netDistributionAfterTaxes := c.NetDistributionAfterTaxesCalculation.CalculateRetirement(model)

	chartData := ChartData{
		BeginningBalance: make(map[int]float64, c.Limit),
		Contribution:     make(map[int]float64, c.Limit),
		Withdrawal:       make(map[int]float64, c.Limit),
		InterestEarned:   make(map[int]float64, c.Limit),
		EndingBalance:    make(map[int]float64, c.Limit),
		AfterTaxIncome:   make(map[int]float64, c.Limit),
	}

	for i := model.Input.CurrentAge; i < c.Limit; i++ {
		if i == model.Input.CurrentAge {
			chartData.BeginningBalance[i] = float64(0)
		} else {
			chartData.BeginningBalance[i] = chartData.EndingBalance[i-1]
		}

		if i < model.Input.RetirementAge {
			chartData.Contribution[i] = float64(model.Input.AnnualContributionsPreTax)
			chartData.Withdrawal[i] = float64(0)
			chartData.AfterTaxIncome[i] = float64(0)
		} else {
			chartData.Contribution[i] = float64(0)
			chartData.Withdrawal[i] = float64(model.Input.YearlyWithdrawal)
			chartData.AfterTaxIncome[i] = netDistributionAfterTaxes
		}

		chartData.InterestEarned[i] = (chartData.BeginningBalance[i] +
			chartData.Contribution[i] -
			chartData.Withdrawal[i]) *
			annualGrowthLessInflation

		chartData.EndingBalance[i] = chartData.BeginningBalance[i] +
			chartData.Contribution[i] -
			chartData.Withdrawal[i] +
			chartData.InterestEarned[i]

		if chartData.EndingBalance[i] <= 0 {
			chartData.EndingBalance[i] = 0.0
			chartData.InterestEarned[i] = 0.0
			chartData.Withdrawal[i] = chartData.EndingBalance[i-1]
			chartData.AfterTaxIncome[i] = chartData.Withdrawal[i] * (1 - effectiveTaxRateOnGross)
		}
	}

	return chartData
}
