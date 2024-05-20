package calculator

type BalancesRothMatchingNetContributionsCalculation ChartCalculation

type BalancesRothMatchingNetContributions struct {
	Limit int32

	AnnualGrowthLessInflationCalculation
	AnnualRetirementAccountDisbursementCalculation
	EquivalentRothContributionsCalculation
}

func NewBalancesRothMatchingNetContributions() BalancesRothMatchingNetContributions {
	return BalancesRothMatchingNetContributions{
		Limit: 133,

		AnnualGrowthLessInflationCalculation:           NewAnnualGrowthLessInflation(),
		AnnualRetirementAccountDisbursementCalculation: NewAnnualRetirementAccountDisbursement(),
		EquivalentRothContributionsCalculation:         NewEquivalentRothContributions(),
	}
}

func (c BalancesRothMatchingNetContributions) Calculate(model *Model) ChartData {
	annualGrowthLessInflation := c.AnnualGrowthLessInflationCalculation.CalculateRothRetirement(model)
	equivalentRothContributions := c.EquivalentRothContributionsCalculation.CalculateRothRetirement(model)
	annualRetirementAccountDisbursementTraditional := c.AnnualRetirementAccountDisbursementCalculation.CalculateTraditionalRetirement(model)
	annualRetirementAccountDisbursementRoth := c.AnnualRetirementAccountDisbursementCalculation.CalculateRothRetirement(model)

	chartData := ChartData{
		BeginningBalance: make(map[int32]float64, c.Limit),
		Contribution:     make(map[int32]float64, c.Limit),
		Withdrawal:       make(map[int32]float64, c.Limit),
		InterestEarned:   make(map[int32]float64, c.Limit),
		EndingBalance:    make(map[int32]float64, c.Limit),
		AfterTaxIncome:   make(map[int32]float64, c.Limit),
	}

	for i := model.Input.CurrentAge; i < c.Limit; i++ {
		if i == model.Input.CurrentAge {
			chartData.BeginningBalance[i] = float64(0)
		} else {
			chartData.BeginningBalance[i] = chartData.EndingBalance[i-1]
		}

		if i < model.Input.RetirementAge {
			chartData.Contribution[i] = float64(equivalentRothContributions)
			chartData.Withdrawal[i] = float64(0)
			chartData.AfterTaxIncome[i] = float64(0)
		} else {
			chartData.Contribution[i] = float64(0)
			chartData.Withdrawal[i] = float64(annualRetirementAccountDisbursementTraditional)
			chartData.AfterTaxIncome[i] = annualRetirementAccountDisbursementRoth
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
			chartData.AfterTaxIncome[i] = chartData.Withdrawal[i]
		}
	}

	return chartData
}
