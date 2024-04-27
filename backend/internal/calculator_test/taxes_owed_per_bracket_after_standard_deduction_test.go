package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestTaxesOwedPerBracketAfterStandardDeduction(t *testing.T) {
	tests := []struct {
		name                                                     string
		model                                                    calculator.Model
		taxesOwedPerBracketAfterStandardDeductionSingle          []float64
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint    []float64
		taxesOwedPerBracketAfterStandardDeductionMarriedSeperate []float64
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold []float64
	}{
		{
			name: "Test Case 0",
			taxesOwedPerBracketAfterStandardDeductionSingle: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "single",
				},
			},
		},
		{
			name: "Test Case 1",
			taxesOwedPerBracketAfterStandardDeductionMarriedJoint: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-joint",
				},
			},
		},
		{
			name: "Test Case 2",
			taxesOwedPerBracketAfterStandardDeductionMarriedSeperate: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-seperate",
				},
			},
		},
		{
			name: "Test Case 3",
			taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "head-of-household",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeperate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate,
				TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.Calculate(test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionSingle.Calculate(test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.Calculate(test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate.Calculate(test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.Calculate(test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}
