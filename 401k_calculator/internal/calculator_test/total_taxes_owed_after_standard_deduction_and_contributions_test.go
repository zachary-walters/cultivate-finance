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

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributions) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedAfterStandardDeductionAndContributionsTests = []struct {
	name                                                                string
	model                                                               calculator.Model
	totalTaxesOwedAfterStandardDeductionAndContributionsSingle          float64
	totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint    float64
	totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate float64
	totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold float64
}{
	{
		name: "Test Case Single",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus: "single",
			},
		},
		totalTaxesOwedAfterStandardDeductionAndContributionsSingle: 123456,
	},
	{
		name: "Test Case Married Joint",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus: "married-joint",
			},
		},
		totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint: 123456,
	},
	{
		name: "Test Case Married Separate",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus: "married-seperate",
			},
		},
		totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate: 123456,
	},
	{
		name: "Test Case Head of Household",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus: "head_of_household",
			},
		},
		totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold: 123456,
	},
}

func TestNewTotalTaxesOwedAfterStandardDeductionAndContributions(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributions()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
		TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(),
		TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint(),
		TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateCalculation: calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate(),
		TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedAfterStandardDeductionAndContributionsCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			totalTaxesOwedAfterStandardDeductionAndContributionsSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
				TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			expected := 0.0
			actual := c.CalculateTraditional(&test.model)
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint
			case "married-seperate":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold
			default:
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionAndContributionsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsTests {

		t.Run(test.name, func(t *testing.T) {
			totalTaxesOwedAfterStandardDeductionAndContributionsSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
				TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			expected := c.CalculateTraditional(&test.model)
			actual := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionAndContributionsCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			totalTaxesOwedAfterStandardDeductionAndContributionsSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
				TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			expected := 0.0
			actual := c.CalculateRoth(&test.model)
			switch test.model.Input.RetirementFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint
			case "married-seperate":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold
			default:
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionAndContributionsCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsTests {

		t.Run(test.name, func(t *testing.T) {
			totalTaxesOwedAfterStandardDeductionAndContributionsSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			totalTaxesOwedAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsSingle)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint)
			totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate)
			totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributions{
				TotalTaxesOwedAfterStandardDeductionAndContributionsSingleCalculation:          totalTaxesOwedAfterStandardDeductionAndContributionsSingle,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointCalculation:    totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			expected := c.CalculateTraditionalRetirement(&test.model)
			actual := c.CalculateRothRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
