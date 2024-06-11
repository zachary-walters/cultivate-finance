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

var totalTaxesOwedAfterStandardDeductionTests = []struct {
	name                                                string
	model                                               calculator.Model
	totalTaxesOwedAfterStandardDeductionSingle          float64
	totalTaxesOwedAfterStandardDeductionMarriedJoint    float64
	totalTaxesOwedAfterStandardDeductionMarriedSeparate float64
	totalTaxesOwedAfterStandardDeductionHeadOfHousehold float64
}{
	{
		name: "Test Case Single",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "single",
				RetirementFilingStatus: "single",
			},
		},
		totalTaxesOwedAfterStandardDeductionSingle:          43241,
		totalTaxesOwedAfterStandardDeductionMarriedJoint:    509865,
		totalTaxesOwedAfterStandardDeductionMarriedSeparate: 23487,
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold: 123,
	},
	{
		name: "Test Case Married Joint",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-joint",
				RetirementFilingStatus: "married-joint",
			},
		},
		totalTaxesOwedAfterStandardDeductionSingle:          43241,
		totalTaxesOwedAfterStandardDeductionMarriedJoint:    509865,
		totalTaxesOwedAfterStandardDeductionMarriedSeparate: 23487,
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold: 123,
	},
	{
		name: "Test Case Married Separate",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-separate",
				RetirementFilingStatus: "married-separate",
			},
		},
		totalTaxesOwedAfterStandardDeductionSingle:          43241,
		totalTaxesOwedAfterStandardDeductionMarriedJoint:    509865,
		totalTaxesOwedAfterStandardDeductionMarriedSeparate: 23487,
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold: 123,
	},
	{
		name: "Test Case Head of Household",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "head-of-household",
				RetirementFilingStatus: "head-of-household",
			},
		},
		totalTaxesOwedAfterStandardDeductionSingle:          43241,
		totalTaxesOwedAfterStandardDeductionMarriedJoint:    509865,
		totalTaxesOwedAfterStandardDeductionMarriedSeparate: 23487,
		totalTaxesOwedAfterStandardDeductionHeadOfHousehold: 123,
	},
	{
		name: "Test Case default",
	},
}

func TestNewTotalTaxesOwedAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeduction()
	expected := calculator.TotalTaxesOwedAfterStandardDeduction{
		TotalTaxesOwedAfterStandardDeductionSingleCalculation:          calculator.NewTotalTaxesOwedAfterStandardDeductionSingle(),
		TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedJoint(),
		TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedSeparate(),
		TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: calculator.NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			mockTotalTaxesOwedAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			c := calculator.TotalTaxesOwedAfterStandardDeduction{
				TotalTaxesOwedAfterStandardDeductionSingleCalculation:          mockTotalTaxesOwedAfterStandardDeductionSingle,
				TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    mockTotalTaxesOwedAfterStandardDeductionMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold,
			}

			expected := 0.0
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedJoint
			case "married-separate":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold
			default:
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			}

			actual := c.CalculateTraditional(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			mockTotalTaxesOwedAfterStandardDeductionSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			c := calculator.TotalTaxesOwedAfterStandardDeduction{
				TotalTaxesOwedAfterStandardDeductionSingleCalculation:          mockTotalTaxesOwedAfterStandardDeductionSingle,
				TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    mockTotalTaxesOwedAfterStandardDeductionMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold,
			}

			expected := 0.0
			switch test.model.Input.RetirementFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedJoint
			case "married-separate":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold
			default:
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			mockTotalTaxesOwedAfterStandardDeductionSingle.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			c := calculator.TotalTaxesOwedAfterStandardDeduction{
				TotalTaxesOwedAfterStandardDeductionSingleCalculation:          mockTotalTaxesOwedAfterStandardDeductionSingle,
				TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    mockTotalTaxesOwedAfterStandardDeductionMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold,
			}

			expected := 0.0
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedJoint
			case "married-separate":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold
			default:
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			}

			actual := c.CalculateRoth(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxesOwedAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint := new(MockTotalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate := new(MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold := new(MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			mockTotalTaxesOwedAfterStandardDeductionSingle.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionSingle)
			mockTotalTaxesOwedAfterStandardDeductionMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedJoint)
			mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionMarriedSeparate)
			mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold)

			c := calculator.TotalTaxesOwedAfterStandardDeduction{
				TotalTaxesOwedAfterStandardDeductionSingleCalculation:          mockTotalTaxesOwedAfterStandardDeductionSingle,
				TotalTaxesOwedAfterStandardDeductionMarriedJointCalculation:    mockTotalTaxesOwedAfterStandardDeductionMarriedJoint,
				TotalTaxesOwedAfterStandardDeductionMarriedSeparateCalculation: mockTotalTaxesOwedAfterStandardDeductionMarriedSeparate,
				TotalTaxesOwedAfterStandardDeductionHeadOfHouseholdCalculation: mockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold,
			}

			expected := 0.0
			switch test.model.Input.RetirementFilingStatus {
			case "single":
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			case "married-joint":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedJoint
			case "married-separate":
				expected = test.totalTaxesOwedAfterStandardDeductionMarriedSeparate
			case "head-of-household":
				expected = test.totalTaxesOwedAfterStandardDeductionHeadOfHousehold
			default:
				expected = test.totalTaxesOwedAfterStandardDeductionSingle
			}

			actual := c.CalculateRothRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
