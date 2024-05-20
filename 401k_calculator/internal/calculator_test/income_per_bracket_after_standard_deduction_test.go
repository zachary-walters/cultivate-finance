package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeduction) CalculateTraditional(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeduction) CalculateRoth(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func TestIncomePerBracketAfterStandardDeductionCalculate(t *testing.T) {
	tests := []struct {
		name                                                  string
		model                                                 calculator.Model
		incomePerBracketAfterStandardDeductionSingle          []float64
		incomePerBracketAfterStandardDeductionMarriedJoint    []float64
		incomePerBracketAfterStandardDeductionMarriedSeperate []float64
		incomePerBracketAfterStandardDeductionHeadOfHousehold []float64
	}{
		{
			name: "Test Case 0",
			incomePerBracketAfterStandardDeductionSingle: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "single",
				},
				SingleTaxRates: []calculator.TaxRate{
					{
						Cap:  12.0,
						Rate: 0.123,
					},
					{
						Cap:  4214.0,
						Rate: 0.646546,
					},
					{
						Cap:  4564.0,
						Rate: 0.231,
					},
				},
			},
		},
		{
			name: "Test Case 1",
			incomePerBracketAfterStandardDeductionMarriedJoint: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-joint",
				},
				MarriedJointTaxRates: []calculator.TaxRate{
					{
						Cap:  123.0,
						Rate: 0.123,
					},
					{
						Cap:  4214.0,
						Rate: 0.646546,
					},
					{
						Cap:  4564.0,
						Rate: 0.231,
					},
				},
			},
		},
		{
			name: "Test Case 2",
			incomePerBracketAfterStandardDeductionMarriedSeperate: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-seperate",
				},
				MarriedSeperateTaxRates: []calculator.TaxRate{
					{
						Cap:  123.0,
						Rate: 0.123,
					},
					{
						Cap:  4214.0,
						Rate: 0.646546,
					},
					{
						Cap:  4564.0,
						Rate: 0.231,
					},
				},
			},
		},
		{
			name: "Test Case 3",
			incomePerBracketAfterStandardDeductionHeadOfHousehold: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "head-of-household",
				},
				HeadOfHouseholdTaxRates: []calculator.TaxRate{
					{
						Cap:  123.0,
						Rate: 0.123,
					},
					{
						Cap:  4214.0,
						Rate: 0.646546,
					},
					{
						Cap:  4564.0,
						Rate: 0.231,
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeduction{
				IncomePerBracketAfterStandardDeductionSingleCalculation:          mockIncomePerBracketAfterStandardDeductionSingle,
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionMarriedJoint,
				IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeperate,
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return test.incomePerBracketAfterStandardDeductionSingle
				case "married-joint":
					return test.incomePerBracketAfterStandardDeductionMarriedJoint
				case "married-seperate":
					return test.incomePerBracketAfterStandardDeductionMarriedSeperate
				case "head-of-household":
					return test.incomePerBracketAfterStandardDeductionHeadOfHousehold
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})

	}
}
