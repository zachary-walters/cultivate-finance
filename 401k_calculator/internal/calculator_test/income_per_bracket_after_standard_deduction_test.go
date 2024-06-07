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

var incomePerBracketAfterStandardDeductionTests = []struct {
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
				CurrentFilingStatus:    "single",
				RetirementFilingStatus: "single",
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
				CurrentFilingStatus:    "married-joint",
				RetirementFilingStatus: "married-joint",
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
				CurrentFilingStatus:    "married-seperate",
				RetirementFilingStatus: "married-seperate",
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
				CurrentFilingStatus:    "head-of-household",
				RetirementFilingStatus: "head-of-household",
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

func TestNewTaxesOwedPerBracketAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeduction()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeduction{
		TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
		TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
		TaxesOwedPerBracketAfterStandardDeductionMarriedSeperateCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate(),
		TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestIncomePerBracketAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionTests {
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

func TestIncomePerBracketAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeduction{
				IncomePerBracketAfterStandardDeductionSingleCalculation:          mockIncomePerBracketAfterStandardDeductionSingle,
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionMarriedJoint,
				IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeperate,
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
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

func TestIncomePerBracketAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeduction{
				IncomePerBracketAfterStandardDeductionSingleCalculation:          mockIncomePerBracketAfterStandardDeductionSingle,
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionMarriedJoint,
				IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeperate,
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)
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

func TestIncomePerBracketAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedSeperate.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeperate)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeduction{
				IncomePerBracketAfterStandardDeductionSingleCalculation:          mockIncomePerBracketAfterStandardDeductionSingle,
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionMarriedJoint,
				IncomePerBracketAfterStandardDeductionMarriedSeperateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeperate,
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
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
