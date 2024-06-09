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

var taxesOwedPerBracketAfterStandardDeductionTests = []struct {
	name                                                     string
	model                                                    calculator.Model
	taxesOwedPerBracketAfterStandardDeductionSingle          []float64
	taxesOwedPerBracketAfterStandardDeductionMarriedJoint    []float64
	taxesOwedPerBracketAfterStandardDeductionMarriedSeparate []float64
	taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold []float64
}{
	{
		name: "Test Case 0",
		taxesOwedPerBracketAfterStandardDeductionSingle: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "single",
				RetirementFilingStatus: "single",
			},
		},
	},
	{
		name: "Test Case 1",
		taxesOwedPerBracketAfterStandardDeductionMarriedJoint: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-joint",
				RetirementFilingStatus: "married-joint",
			},
		},
	},
	{
		name: "Test Case 2",
		taxesOwedPerBracketAfterStandardDeductionMarriedSeparate: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "married-seperate",
				RetirementFilingStatus: "married-seperate",
			},
		},
	},
	{
		name: "Test Case 3",
		taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold: []float64{1, 2, 3, 4, 5},
		model: calculator.Model{
			Input: calculator.Input{
				CurrentFilingStatus:    "head-of-household",
				RetirementFilingStatus: "head-of-household",
			},
		},
	},
}

func TestNewTaxesOwedPerBracketAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeduction()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeduction{
		TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
		TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
		TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate(),
		TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate,
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
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.CalculateTraditional(&test.model)
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

func TestTaxesOwedPerBracketAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionSingle.CalculateTraditionalRetirement(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.CalculateTraditionalRetirement(&test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.CalculateTraditionalRetirement(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.CalculateTraditionalRetirement(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionSingle.CalculateRoth(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.CalculateRoth(&test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.CalculateRoth(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.CalculateRoth(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionSingle.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionMarriedSeparate)
			mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeduction{
				TaxesOwedPerBracketAfterStandardDeductionSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionSingle,
				TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate,
				TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() []float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionSingle.CalculateRothRetirement(&test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint.CalculateRothRetirement(&test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate.CalculateRothRetirement(&test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold.CalculateRothRetirement(&test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}
