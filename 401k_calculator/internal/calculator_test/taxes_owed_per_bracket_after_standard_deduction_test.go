package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) CalculateTraditional(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) CalculateRoth(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func TestTaxesOwedPerBracketAfterStandardDeductionCalculateTraditional(t *testing.T) {
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

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeperate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate,
				TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionSingle.CalculateTraditional(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.CalculateTraditional(&test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate.CalculateTraditional(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.CalculateTraditional(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}
