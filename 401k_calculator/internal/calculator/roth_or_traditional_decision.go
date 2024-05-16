package calculator

type RothOrTraditionalDecisionCalculation DecisionCalculation

type RothOrTraditionalDecision struct {
	TotalDisbursementsCalculation
}

func NewRothOrTraditionalDecision() RothOrTraditionalDecision {
	return RothOrTraditionalDecision{
		TotalDisbursementsCalculation: NewTotalDisbursements(),
	}
}

func (c RothOrTraditionalDecision) Calculate(model Model) string {
	totalDisbursementsTraditionalRetirement := c.TotalDisbursementsCalculation.CalculateTraditionalRetirement(model)
	totalDisbursementsRothRetirement := c.TotalDisbursementsCalculation.CalculateRothRetirement(model)

	if totalDisbursementsTraditionalRetirement >= totalDisbursementsRothRetirement {
		return "Traditional"
	}

	return "Roth"
}
