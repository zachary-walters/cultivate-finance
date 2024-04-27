package calculator

type AbstractIncomePerBracketCalculation interface {
	Calculate([]TaxRate, int, float64) float64
}

type AbstractIncomePerBracket struct{}

func NewAbstractIncomePerBracket() AbstractIncomePerBracket {
	return AbstractIncomePerBracket{}
}

func (c AbstractIncomePerBracket) Calculate(taxRates []TaxRate, bracketSequence int, income float64) float64 {
	incomePerBracket := 0.0
	previousCap := 0.0
	sum := 0.0
	overflow := false
	for idx, taxRate := range taxRates {
		if idx > bracketSequence {
			return incomePerBracket
		} else if overflow { // when given a bracketSequence above income
			return 0
		} else if income > taxRate.Cap {
			incomePerBracket = taxRate.Cap - previousCap
			sum += incomePerBracket
		} else {
			incomePerBracket = income - sum
			overflow = true
		}

		previousCap = taxRate.Cap
	}

	return incomePerBracket
}
