package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandardDeductionMarriedJointTests = []struct {
	name                                               string
	model                                              calculator.Model
	incomePerBracketAfterStandardDeductionMarriedJoint []float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			MarriedJointTaxRates: []calculator.TaxRate{
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
		incomePerBracketAfterStandardDeductionMarriedJoint: []float64{
			1.0,
			1.0,
			1.0,
			1.0,
			1.0,
			1.0,
			1.0,
		},
	},
}

func TestNewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
		IncomePerBracketAfterStandardDeductionMarriedJointCalculation: calculator.NewIncomePerBracketAfterStandardDeductionMarriedJoint(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionMarriedJoint,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionMarriedJoint,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionMarriedJoint,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedJoint := new(MockIncomePerBracketAfterStandardDeductionMarriedJoint)
			mockIncomePerBracketAfterStandardDeductionMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedJoint)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedJoint{
				IncomePerBracketAfterStandardDeductionMarriedJointCalculation: mockIncomePerBracketAfterStandardDeductionMarriedJoint,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedJoint := c.IncomePerBracketAfterStandardDeductionMarriedJointCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedJointTaxRates))

			for idx, taxRate := range test.model.MarriedJointTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedJoint[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
