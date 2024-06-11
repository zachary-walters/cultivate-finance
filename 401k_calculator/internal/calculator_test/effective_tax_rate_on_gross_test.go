package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockEffectiveTaxRateOnGross struct {
	mock.Mock
}

func (m *MockEffectiveTaxRateOnGross) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockEffectiveTaxRateOnGross) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var effectiveTaxRateOnGrossTests = []struct {
	name                                                                  string
	model                                                                 calculator.Model
	totalTaxesOwedAfterStandardDeductionAndContributions                  float64
	totalTaxesOwedAfterStandardDeduction                                  float64
	totalAnnualRetirementIncomeBeforeTax                                  float64
	totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal float64
}{
	{
		name: "Test Case 0",
		totalTaxesOwedAfterStandardDeductionAndContributions:                  100,
		totalTaxesOwedAfterStandardDeduction:                                  200,
		totalAnnualRetirementIncomeBeforeTax:                                  23409,
		totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal: 1435,
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAnnualIncome: 41987,
			},
		},
	},
	{
		name: "Test Case 1",
		totalTaxesOwedAfterStandardDeductionAndContributions:                  0,
		totalTaxesOwedAfterStandardDeduction:                                  0,
		totalAnnualRetirementIncomeBeforeTax:                                  12312,
		totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal: 98735,
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAnnualIncome: 10000,
			},
		},
	},
	{
		name: "Test Case 2",
		totalTaxesOwedAfterStandardDeductionAndContributions:                  10000,
		totalTaxesOwedAfterStandardDeduction:                                  214500,
		totalAnnualRetirementIncomeBeforeTax:                                  0,
		totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal: 0,
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAnnualIncome: 0,
			},
		},
	},
}

func TestNewEffectiveTaxRateOnGross(t *testing.T) {
	actual := calculator.NewEffectiveTaxRateOnGross()
	expected := calculator.EffectiveTaxRateOnGross{
		TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation:                  calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributions(),
		TotalTaxesOwedAfterStandardDeductionCalculation:                                  calculator.NewTotalTaxesOwedAfterStandardDeduction(),
		TotalAnnualRetirementIncomeBeforeTaxCalculation:                                  calculator.NewTotalAnnualRetirementIncomeBeforeTax(),
		TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculation: calculator.NewTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal(),
	}

	assert.Equal(t, expected, actual)
}

func TestEffectiveTaxRateOnGrossCalculateTraditional(t *testing.T) {
	for _, test := range effectiveTaxRateOnGrossTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionAndContributions := new(MockTotalTaxesOwedAfterStandardDeductionAndContributions)
			mockTotalTaxesOwedAfterStandardDeductionAndContributions.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributions)

			c := &calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: mockTotalTaxesOwedAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() float64 {
				if test.model.Input.CurrentAnnualIncome == 0 {
					return 0
				}

				return test.totalTaxesOwedAfterStandardDeductionAndContributions / test.model.Input.CurrentAnnualIncome
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestEffectiveTaxRateOnGrossCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range effectiveTaxRateOnGrossTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeduction)

			mockTotalAnnualRetirementIncomeBeforeTax := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTax.On("CalculateTraditionalRetirement", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTax)

			c := &calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTax,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				if test.totalAnnualRetirementIncomeBeforeTax == 0.0 {
					return 0
				}

				return test.totalTaxesOwedAfterStandardDeduction / test.totalAnnualRetirementIncomeBeforeTax
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestEffectiveTaxRateOnGrossCalculateRoth(t *testing.T) {
	for _, test := range effectiveTaxRateOnGrossTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionAndContributions := new(MockTotalTaxesOwedAfterStandardDeductionAndContributions)
			mockTotalTaxesOwedAfterStandardDeductionAndContributions.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributions)

			c := &calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionAndContributionsCalculation: mockTotalTaxesOwedAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.CalculateRoth(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestEffectiveTaxRateOnGrossCalculateRothRetirement(t *testing.T) {
	for _, test := range effectiveTaxRateOnGrossTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeduction)

			mockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal := new(MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal)
			mockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal.On("CalculateRothRetirement", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal)

			c := &calculator.EffectiveTaxRateOnGross{
				TotalTaxesOwedAfterStandardDeductionCalculation:                                  mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculation: mockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				if test.totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal == 0 {
					return 0
				}

				return test.totalTaxesOwedAfterStandardDeduction / test.totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
