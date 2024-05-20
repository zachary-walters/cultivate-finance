package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxesOwedAfterStandardDeductionCalculate(t *testing.T) {
	totalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionSingle)
	totalTaxesOwedAfterStandardDeductionMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionMarriedJoint)
	totalTaxesOwedAfterStandardDeductionMarriedSeperate := new(MockTotalTaxesOwedAfterStandardDeductionMarriedSeperate)
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold)

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
					CurrentFilingStatus: "married_joint",
				},
			},
		},
		{
			name: "Test Case Married Seperate",
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married_seperate",
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

	for _, test := range tests {
		totalTaxesOwedAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(1337.0)
		totalTaxesOwedAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(90245.7)
		totalTaxesOwedAfterStandardDeductionMarriedSeperate.On("CalculateTraditional", &test.model).Return(345.89)
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(1233214.908)

		totalTaxesOwedAfterStandardDeduction := calculator.TotalTaxesOwedAfterStandardDeduction{
			TotalTaxesOwedAfterStandardDeductionSingleCalculation:          totalTaxesOwedAfterStandardDeductionSingle,
			TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionMarriedJoint,
			TotalTaxesOwedAfterStandardDeductionMarriedSeperateCalculation: totalTaxesOwedAfterStandardDeductionMarriedSeperate,
			TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionHeadOfHousehold,
		}
		t.Run(test.name, func(t *testing.T) {
			expected := 0.0
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = 1337.0
			case "married-joint":
				expected = 90245.7
			case "married-seperate":
				expected = 345.89
			case "head-of-household":
				expected = 1233214.908
			default:
				expected = 0
			}

			actual := totalTaxesOwedAfterStandardDeduction.CalculateTraditional(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
