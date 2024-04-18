package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeduction) CalculateRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalTaxesOwedAfterStandardDeductionCalculate(t *testing.T) {
	totalTaxesOwedAfterStandardDeductionSingle := new(MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle)

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
		totalTaxesOwedAfterStandardDeductionSingle.On("Calculate", test.model).Return(1337.0)

		totalTaxesOwedAfterStandardDeduction := calculator.TotalTaxesOwedAfterStandardDeduction{
			TotalTaxesOwedAfterStandardDeductionSingleCalculation: totalTaxesOwedAfterStandardDeductionSingle,
		}
		t.Run(test.name, func(t *testing.T) {
			expected := 0.0
			switch test.model.Input.CurrentFilingStatus {
			case "single":
				expected = 1337.0
			case "married_joint":
			case "married_seperate":
			case "head_of_household":
			default:
				expected = 0
			}

			actual := totalTaxesOwedAfterStandardDeduction.Calculate(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
