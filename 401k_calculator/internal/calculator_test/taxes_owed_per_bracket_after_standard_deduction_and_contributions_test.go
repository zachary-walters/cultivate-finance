package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) CalculateTraditional(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) CalculateTraditionalRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) CalculateRoth(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) CalculateRothRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandardDeductionAndContributionsTests = []struct {
	name                                                                     string
	model                                                                    calculator.Model
	taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle          []float64
	taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint    []float64
	taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate []float64
	taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold []float64
}{
	{
		name: "Test Case 0",
		taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "single",
				RetirementFilingStatus: "single",
			},
		},
	},
	{
		name: "Test Case 1",
		taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-joint",
				RetirementFilingStatus: "married-joint",
			},
		},
	},
	{
		name: "Test Case 2",
		taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-separate",
				RetirementFilingStatus: "married-separate",
			},
		},
	},
	{
		name: "Test Case 3",
		taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "head-of-household",
				RetirementFilingStatus: "head-of-household",
			},
		},
	},
}

func TestNewTaxesOwedPerBracketAfterStandardDeductionAndContributions(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributions()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate(),
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.CalculateTraditional(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.CalculateTraditional(&test.model)
				case "married-separate":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.CalculateTraditional(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.CalculateTraditional(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.CalculateTraditionalRetirement(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.CalculateTraditionalRetirement(&test.model)
				case "married-separate":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.CalculateTraditionalRetirement(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.CalculateTraditionalRetirement(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, expected)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.CalculateRoth(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.CalculateRoth(&test.model)
				case "married-separate":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.CalculateRoth(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.CalculateRoth(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionAndContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.CalculateRothRetirement(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.CalculateRothRetirement(&test.model)
				case "married-separate":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.CalculateRothRetirement(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.CalculateRothRetirement(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, expected)
		})
	}
}
