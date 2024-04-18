package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestTotalTaxesOwedAfterStandardDeductionAndContributions(t *testing.T) {
	totalTaxesOwedAfterStandardDeductionAndContributionsSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)

	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case Single",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "single",
				},
			},
		},
		// {
		// 	name: "Test Case Married Joint",
		// 	model: calculator.Model{
		// 		Input: calculator.Input{
		// 			CurrentFilingStatus: "married_joint",
		// 		},
		// 	},
		// },
		// {
		// 	name: "Test Case Married Seperate",
		// 	model: calculator.Model{
		// 		Input: calculator.Input{
		// 			CurrentFilingStatus: "married_seperate",
		// 		},
		// 	},
		// },
		// {
		// 	name: "Test Case Head of Household",
		// 	model: calculator.Model{
		// 		Input: calculator.Input{
		// 			CurrentFilingStatus: "head_of_household",
		// 		},
		// 	},
		// },
	}

	c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
		TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
	}

	for _, test := range tests {
		totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("Calculate", test.model).Return(1337.0)
		expected := 0.0
		actual := c.Calculate(test.model)
		t.Run(test.name, func(t *testing.T) {
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = 1337.0
			case "married_joint":
			case "married_seperate":
			case "head_of_household":
			default:
			}

			assert.Equal(t, expected, actual)

		})
	}
}
