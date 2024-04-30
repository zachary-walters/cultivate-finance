package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
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
	totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
	totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate)
	totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

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
		{
			name: "Test Case Married Joint",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-joint",
				},
			},
		},
		{
			name: "Test Case Married Seperate",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-seperate",
				},
			},
		},
		{
			name: "Test Case Head of Household",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "head_of_household",
				},
			},
		},
	}

	c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
		TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
		TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint,
		TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperateCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate,
		TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold,
	}

	for _, test := range tests {
		totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("Calculate", test.model).Return(1337.0)
		totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint.On("Calculate", test.model).Return(349587.0)
		totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeperate.On("Calculate", test.model).Return(10003.0)
		totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold.On("Calculate", test.model).Return(4387.8)

		expected := 0.0
		actual := c.Calculate(test.model)
		t.Run(test.name, func(t *testing.T) {
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = 1337.0
			case "married-joint":
				expected = 349587.0
			case "married-seperate":
				expected = 10003.0
			case "head-of-household":
				expected = 4387.8
			default:
			}

			assert.Equal(t, expected, actual)

		})
	}
}
