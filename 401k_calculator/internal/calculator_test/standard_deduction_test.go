package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockStandardDeduction struct {
	mock.Mock
}

func (m *MockStandardDeduction) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockStandardDeduction) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockStandardDeduction) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var standardDeductionTests = []struct {
	name  string
	model calculator.Model
}{
	{
		name: "Test Case Single",
		model: calculator.Model{
			STANDARD_DEDUCTION_SINGLE: 200,
			Input: calculator.Input{
				CurrentFilingStatus:    "single",
				RetirementFilingStatus: "single",
			},
		},
	},
	{
		name: "Test Case Married Joint",
		model: calculator.Model{
			STANDARD_DEDUCTION_MARRIED_JOINT: 200,
			Input: calculator.Input{
				CurrentFilingStatus:    "married-joint",
				RetirementFilingStatus: "married-joint",
			},
		},
	},
	{
		name: "Test Case Married Seperate",
		model: calculator.Model{
			STANDARD_DEDUCTION_MARRIED_SEPERATE: 200,
			Input: calculator.Input{
				CurrentFilingStatus:    "married-separate",
				RetirementFilingStatus: "married-separate",
			},
		},
	},
	{
		name: "Test Case Head of Household",
		model: calculator.Model{
			STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: 200,
			Input: calculator.Input{
				CurrentFilingStatus:    "head-of-household",
				RetirementFilingStatus: "head-of-household",
			},
		},
	},
	{
		name: "Test Case Unknown",
		model: calculator.Model{
			STANDARD_DEDUCTION_SINGLE: 200,
			Input: calculator.Input{
				CurrentFilingStatus:    "unknown",
				RetirementFilingStatus: "unknown",
			},
		},
	},
}

func TestNewStandardDeduction(t *testing.T) {
	actual := calculator.NewStandardDeduction()
	expected := calculator.StandardDeduction{}

	assert.Equal(t, expected, actual)
}

func TestStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range standardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.StandardDeduction{}

			actual := c.CalculateTraditional(&test.model)
			expected := -1.0

			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = test.model.STANDARD_DEDUCTION_SINGLE
			case "married-joint":
				expected = test.model.STANDARD_DEDUCTION_MARRIED_JOINT
			case "married-separate":
				expected = test.model.STANDARD_DEDUCTION_MARRIED_SEPERATE
			case "head-of-household":
				expected = test.model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
			default:
				expected = test.model.STANDARD_DEDUCTION_SINGLE
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range standardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.StandardDeduction{}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := -1.0

			switch test.model.Input.RetirementFilingStatus {
			case "single":
				expected = test.model.STANDARD_DEDUCTION_SINGLE
			case "married-joint":
				expected = test.model.STANDARD_DEDUCTION_MARRIED_JOINT
			case "married-separate":
				expected = test.model.STANDARD_DEDUCTION_MARRIED_SEPERATE
			case "head-of-household":
				expected = test.model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD
			default:
				expected = test.model.STANDARD_DEDUCTION_SINGLE
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range standardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.StandardDeduction{}

			actual := c.CalculateRoth(&test.model)
			expected := c.CalculateTraditional(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range standardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.StandardDeduction{}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
