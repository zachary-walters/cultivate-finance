package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockBalancesRothMatchingNet struct {
	mock.Mock
}

func (m *MockBalancesRothMatchingNet) Calculate(model *calculator.Model) calculator.ChartData {
	args := m.Called(model)
	return args.Get(0).(calculator.ChartData)
}

var balancesRothMatchingNetTests = []struct {
	name                                           string
	model                                          calculator.Model
	annualGrowthLessInflation                      float64
	annualRetirementAccountDisbursementTraditional float64
	annualRetirementAccountDisbursementRoth        float64
	equivalentRothContributions                    float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAge:    30,
				RetirementAge: 65,
			},
		},
	},
}

func TestNewBalancesRothMatchingNetContributions(t *testing.T) {
	actual := calculator.NewBalancesRothMatchingNetContributions()
	expected := calculator.BalancesRothMatchingNetContributions{
		Limit: 133,

		AnnualGrowthLessInflationCalculation:           calculator.NewAnnualGrowthLessInflation(),
		AnnualRetirementAccountDisbursementCalculation: calculator.NewAnnualRetirementAccountDisbursement(),
		EquivalentRothContributionsCalculation:         calculator.NewEquivalentRothContributions(),
	}
	assert.Equal(t, int32(133), actual.Limit)
	assert.Equal(t, expected, actual)
}

func TestCalculate(t *testing.T) {
	for _, test := range balancesRothMatchingNetTests {
		t.Run(test.name, func(t *testing.T) {
			mockAnnualGrowthLessInflation := new(MockAnnualGrowthLessInflation)
			mockAnnualGrowthLessInflation.On("CalculateRothRetirement", &test.model).Return(test.annualGrowthLessInflation)

			mockAnnualRetirementDisbursement := new(MockAnnualRetirementAccountDisbursement)
			mockAnnualRetirementDisbursement.On("CalculateRothRetirement", &test.model).Return(test.annualRetirementAccountDisbursementRoth)
			mockAnnualRetirementDisbursement.On("CalculateTraditionalRetirement", &test.model).Return(test.annualRetirementAccountDisbursementTraditional)

			mockEquivalentRothContributions := new(MockEquivalentRothContributions)
			mockEquivalentRothContributions.On("CalculateRothRetirement", &test.model).Return(test.equivalentRothContributions)

			c := calculator.BalancesRothMatchingNetContributions{
				AnnualGrowthLessInflationCalculation:           mockAnnualGrowthLessInflation,
				AnnualRetirementAccountDisbursementCalculation: mockAnnualRetirementDisbursement,
				EquivalentRothContributionsCalculation:         mockEquivalentRothContributions,
			}

			c.Calculate(&test.model)

		})
	}

	// Test BeginningBalance
	// assert.Equal(t, float64(0), actual.BeginningBalance[model.Input.CurrentAge])

	// // Test Contribution
	// assert.Equal(t, float64(c.EquivalentRothContributionsCalculation.CalculateRothRetirement(model)), chartData.Contribution[model.Input.CurrentAge])

	// // Test Withdrawal
	// assert.Equal(t, float64(0), chartData.Withdrawal[model.Input.CurrentAge])

	// // Test InterestEarned
	// expectedInterest := (chartData.BeginningBalance[model.Input.CurrentAge] +
	// 	chartData.Contribution[model.Input.CurrentAge] -
	// 	chartData.Withdrawal[model.Input.CurrentAge]) *
	// 	c.AnnualGrowthLessInflationCalculation.CalculateRothRetirement(model)
	// assert.Equal(t, expectedInterest, chartData.InterestEarned[model.Input.CurrentAge])

	// // Test EndingBalance
	// expectedEndingBalance := chartData.BeginningBalance[model.Input.CurrentAge] +
	// 	chartData.Contribution[model.Input.CurrentAge] -
	// 	chartData.Withdrawal[model.Input.CurrentAge] +
	// 	chartData.InterestEarned[model.Input.CurrentAge]
	// assert.Equal(t, expectedEndingBalance, chartData.EndingBalance[model.Input.CurrentAge])

	// // Test AfterTaxIncome
	// assert.Equal(t, float64(0), chartData.AfterTaxIncome[model.Input.CurrentAge])
}
