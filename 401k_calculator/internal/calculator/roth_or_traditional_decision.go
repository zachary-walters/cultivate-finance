package calculator

type RothOrTraditionalDecisionCalculation DecisionCalculation

type RothOrTraditionalDecision struct {
	TaxRateOfSavingsCalculation
	EffectiveTaxRateOnGrossCalculation
}

func NewRothOrTraditionalDecision() RothOrTraditionalDecision {
	return RothOrTraditionalDecision{
		TaxRateOfSavingsCalculation:        NewTaxRateOfSavings(),
		EffectiveTaxRateOnGrossCalculation: NewEffectiveTaxRateOnGross(),
	}
}

func (c RothOrTraditionalDecision) Calculate(model Model) string {
	taxRateOfSavings := c.TaxRateOfSavingsCalculation.CalculateTraditional(model)
	effectiveTaxRateOnGross := c.EffectiveTaxRateOnGrossCalculation.CalculateTraditional(model)

	if taxRateOfSavings >= effectiveTaxRateOnGross {
		return "Traditional"
	}

	return "Roth"
}

func (c RothOrTraditionalDecision) CalculateRetirement(model Model) string {
	return c.Calculate(model)
}
