package calculator

type HalfOfSocialSecurityCalculation Calculation

type HalfOfSocialSecurity struct{}

func NewHalfOfSocialSecurity() HalfOfSocialSecurity {
	return HalfOfSocialSecurity{}
}

func (c HalfOfSocialSecurity) CalculateTraditional(model Model) float64 {
	return c.CalculateTraditionalRetirement(model)
}

func (c HalfOfSocialSecurity) CalculateTraditionalRetirement(model Model) float64 {
	return model.Input.SocialSecurity * 0.5
}

func (c HalfOfSocialSecurity) CalculateRoth(model Model) float64 {
	return c.CalculateTraditional(model)
}

func (c HalfOfSocialSecurity) CalculateRothRetirement(model Model) float64 {
	return c.CalculateTraditional(model)
}
