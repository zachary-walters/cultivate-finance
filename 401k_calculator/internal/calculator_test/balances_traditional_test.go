package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockBalancesTraditional struct {
	mock.Mock
}

func (m *MockBalancesTraditional) Calculate(model *calculator.Model) calculator.ChartData {
	args := m.Called(model)
	return args.Get(0).(calculator.ChartData)
}

var balancesTraditionalTests = []struct {
	name                                string
	model                               calculator.Model
	annualGrowthLessInflation           float64
	annualRetirementAccountDisbursement float64
	topTierTaxRate                      float64
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

func TestNewBalancesTraditional(t *testing.T) {
	actual := calculator.NewBalancesTraditional()
	expected := calculator.BalancesTraditional{
		Limit:                                133,
		AnnualGrowthLessInflationCalculation: calculator.NewAnnualGrowthLessInflation(),
		AnnualRetirementAccountDisbursementCalculation: calculator.NewAnnualRetirementAccountDisbursement(),
		TopTierTaxRateCalculation:                      calculator.NewTopTierTaxRate(),
	}

	assert.Equal(t, expected, actual)
}

func TestBalancesTraditionalCalculate(t *testing.T) {
	for _, test := range balancesTraditionalTests {
		t.Run(test.name, func(t *testing.T) {
			mockAnnualGrowthLessInflation := new(MockAnnualGrowthLessInflation)
			mockAnnualGrowthLessInflation.On("CalculateTraditionalRetirement", &test.model).Return(test.annualGrowthLessInflation)

			mockAnnualRetirementDisbursement := new(MockAnnualRetirementAccountDisbursement)
			mockAnnualRetirementDisbursement.On("CalculateTraditionalRetirement", &test.model).Return(test.annualRetirementAccountDisbursement)

			mockTopTierTaxRate := new(MockTopTierTaxRate)
			mockTopTierTaxRate.On("CalculateTraditionalRetirement", &test.model).Return(test.topTierTaxRate)

			c := calculator.BalancesTraditional{
				AnnualGrowthLessInflationCalculation:           mockAnnualGrowthLessInflation,
				AnnualRetirementAccountDisbursementCalculation: mockAnnualRetirementDisbursement,
				TopTierTaxRateCalculation:                      mockTopTierTaxRate,

				Limit: 133,
			}

			actual := c.Calculate(&test.model)
			expected := func() calculator.ChartData {
				chartData := calculator.ChartData{
					BeginningBalance: make(map[int32]float64, c.Limit),
					Contribution:     make(map[int32]float64, c.Limit),
					Withdrawal:       make(map[int32]float64, c.Limit),
					InterestEarned:   make(map[int32]float64, c.Limit),
					EndingBalance:    make(map[int32]float64, c.Limit),
					AfterTaxIncome:   make(map[int32]float64, c.Limit),
				}

				for i := test.model.Input.CurrentAge; i < c.Limit; i++ {
					if i == test.model.Input.CurrentAge {
						chartData.BeginningBalance[i] = float64(0)
					} else {
						chartData.BeginningBalance[i] = chartData.EndingBalance[i-1]
					}

					if i < test.model.Input.RetirementAge {
						chartData.Contribution[i] = float64(test.model.Input.AnnualContributionsPreTax)
						chartData.Withdrawal[i] = float64(0)
						chartData.AfterTaxIncome[i] = float64(0)
					} else {
						chartData.Contribution[i] = float64(0)
						chartData.Withdrawal[i] = float64(test.model.Input.YearlyWithdrawal)
						chartData.AfterTaxIncome[i] = test.annualRetirementAccountDisbursement
					}

					chartData.InterestEarned[i] = (chartData.BeginningBalance[i] +
						chartData.Contribution[i] -
						chartData.Withdrawal[i]) *
						test.annualGrowthLessInflation

					chartData.EndingBalance[i] = chartData.BeginningBalance[i] +
						chartData.Contribution[i] -
						chartData.Withdrawal[i] +
						chartData.InterestEarned[i]

					if chartData.EndingBalance[i] <= 0 {
						chartData.EndingBalance[i] = 0.0
						chartData.InterestEarned[i] = 0.0
						chartData.Withdrawal[i] = chartData.EndingBalance[i-1]
						chartData.AfterTaxIncome[i] = chartData.Withdrawal[i] * (1 - test.topTierTaxRate)
					}
				}

				return chartData
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
