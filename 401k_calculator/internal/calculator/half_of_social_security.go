package calculator

type HalfOfSocialSecurityCalculation Calculation

type HalfOfSocialSecurity struct{}

func NewHalfOfSocialSecurity() HalfOfSocialSecurity {
	return HalfOfSocialSecurity{}
}

func (c HalfOfSocialSecurity) Calculate(model Model) float64 {
	return model.Input.SocialSecurity * 0.5
}

func (c HalfOfSocialSecurity) CalculateRetirement(model Model) float64 {
	return c.Calculate(model)
}
