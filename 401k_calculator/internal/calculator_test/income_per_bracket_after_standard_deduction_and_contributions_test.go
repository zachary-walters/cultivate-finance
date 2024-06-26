package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditional(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributions) CalculateRoth(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributions) CalculateRothRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandardDeductionAndContributionsTests = []struct {
	name                                                                  string
	model                                                                 calculator.Model
	incomePerBracketAfterStandardDeductionAndContributionsSingle          []float64
	incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint    []float64
	incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate []float64
	incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold []float64
}{
	{
		name: "Test Case 0",
		incomePerBracketAfterStandardDeductionAndContributionsSingle: []float64{1, 2, 3, 4, 5},
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
		incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint: []float64{1, 2, 3, 4, 5},
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
		incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-seperate",
				RetirementFilingStatus: "married-seperate",
			},
			MarriedSeparateTaxRates: []calculator.TaxRate{
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
		incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold: []float64{1, 2, 3, 4, 5},
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
	{
		name: "Test Case 4",
	},
}

func TestNewIncomePerBracketAfterStandardDeductionAndContributinos(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionAndContributions()
	expected := calculator.IncomePerBracketAfterStandardDeductionAndContributions{
		IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation:          calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
		IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
		IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate(),
		IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeductionAndContributions{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				case "married-joint":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint
				case "married-seperate":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate
				case "head-of-household":
					return test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold
				default:
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				}
			}()

			assert.Equal(t, expected, actual)

		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeductionAndContributions{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				case "married-joint":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint
				case "married-seperate":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate
				case "head-of-household":
					return test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold
				default:
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				}
			}()

			assert.Equal(t, expected, actual)

		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeductionAndContributions{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				case "married-joint":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint
				case "married-seperate":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate
				case "head-of-household":
					return test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold
				default:
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				}
			}()

			assert.Equal(t, expected, actual)

		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.IncomePerBracketAfterStandardDeductionAndContributions{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				case "married-joint":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedJoint
				case "married-seperate":
					return test.incomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate
				case "head-of-household":
					return test.incomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold
				default:
					return test.incomePerBracketAfterStandardDeductionAndContributionsSingle
				}
			}()

			assert.Equal(t, expected, actual)

		})
	}
}
