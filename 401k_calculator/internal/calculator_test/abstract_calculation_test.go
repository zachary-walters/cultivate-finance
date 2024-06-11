package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

func TestNewModel(t *testing.T) {
	input := calculator.Input{}

	model := calculator.NewModel(input)

	assert.Equal(t, input, model.Input)

	input = calculator.Input{
		CurrentAge:                8975,
		RetirementAge:             392,
		RetirementFilingStatus:    "status",
		CurrentFilingStatus:       "status1",
		AnnualContributionsPreTax: 534532,
		AnnuityIncome:             334225,
	}

	model = calculator.NewModel(input)

	assert.Equal(t, input, model.Input)
	assert.Equal(t, calculator.Constants.InflationRate, model.InflationRate)
	assert.Equal(t, calculator.Constants.SingleTaxRates, model.SingleTaxRates)
	assert.Equal(t, calculator.Constants.MarriedJointTaxRates, model.MarriedJointTaxRates)
	assert.Equal(t, calculator.Constants.MarriedSeparateTaxRates, model.MarriedSeparateTaxRates)
	assert.Equal(t, calculator.Constants.HeadOfHouseholdTaxRates, model.HeadOfHouseholdTaxRates)
	assert.Equal(t, calculator.Constants.SocialSecurityTaxRatesIndividual, model.SocialSecurityTaxRatesIndividual)
	assert.Equal(t, calculator.Constants.SocialSecurityTaxRatesJoint, model.SocialSecurityTaxRatesJoint)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_SINGLE, model.STANDARD_DEDUCTION_SINGLE)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_MARRIED_JOINT, model.STANDARD_DEDUCTION_MARRIED_JOINT)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_MARRIED_SEPERATE, model.STANDARD_DEDUCTION_MARRIED_SEPERATE)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD, model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD)
}
