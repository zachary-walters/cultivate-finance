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
						chartData.Contribution[i] = float64(test.equivalentRothContributions)
						chartData.Withdrawal[i] = float64(0)
						chartData.AfterTaxIncome[i] = float64(0)
					} else {
						chartData.Contribution[i] = float64(0)
						chartData.Withdrawal[i] = float64(test.annualRetirementAccountDisbursementTraditional)
						chartData.AfterTaxIncome[i] = test.annualRetirementAccountDisbursementRoth
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
						chartData.AfterTaxIncome[i] = chartData.Withdrawal[i]
					}
				}

				return chartData
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
