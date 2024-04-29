package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockStandardDeduction struct {
	mock.Mock
}

func (m *MockStandardDeduction) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockStandardDeduction) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestStandardDeductionCalculate(t *testing.T) {
	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case Single",
			model: calculator.Model{
				STANDARD_DEDUCTION_SINGLE: 200,
				Input: calculator.Input{
					CurrentFilingStatus: "single",
				},
			},
		},
		{
			name: "Test Case Married Joint",
			model: calculator.Model{
				STANDARD_DEDUCTION_MARRIED_JOINT: 200,
				Input: calculator.Input{
					CurrentFilingStatus: "married-joint",
				},
			},
		},
		{
			name: "Test Case Married Seperate",
			model: calculator.Model{
				STANDARD_DEDUCTION_MARRIED_SEPERATE: 200,
				Input: calculator.Input{
					CurrentFilingStatus: "married-seperate",
				},
			},
		},
		{
			name: "Test Case Head of Household",
			model: calculator.Model{
				STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: 200,
				Input: calculator.Input{
					CurrentFilingStatus: "head-of-household",
				},
			},
		},
		{
			name: "Test Case Unknown",
			model: calculator.Model{
				STANDARD_DEDUCTION_SINGLE: 200,
				Input: calculator.Input{
					CurrentFilingStatus: "unknown",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.StandardDeduction{}

			actual := c.Calculate(test.model)
			expected := -1.0

			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = test.model.STANDARD_DEDUCTION_SINGLE
			case "married-joint":
				expected = test.model.STANDARD_DEDUCTION_MARRIED_JOINT
			case "married-seperate":
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
