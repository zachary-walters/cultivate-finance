package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionMarriedJoint struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedJoint) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedJoint) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandardDeductionMarriedJointTests = []struct {
	name                                   string
	model                                  calculator.Model
	incomePerBracketAfterStandardDeduction []float64
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
		incomePerBracketAfterStandardDeduction: []float64{1, 2, 3, 4, 5},
	},
}

func NewIncomePerBracketAfterStandardDeductionMarriedJoint(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionMarriedJoint()
	expected := calculator.IncomePerBracketAfterStandardDeductionMarriedJoint{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: calculator.NewAbstractIncomePerBracketAfterStandardDeduction(),
	}

	assert.Equal(t, expected, actual)
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedJointCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditional", &test.model, test.model.MarriedJointTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedJoint{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(&test.model, test.model.MarriedJointTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedJointCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model, test.model.MarriedJointTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedJoint{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(&test.model, test.model.MarriedJointTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedJointCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRoth", &test.model, test.model.MarriedJointTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedJoint{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(&test.model, test.model.MarriedJointTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedJointCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRothRetirement", &test.model, test.model.MarriedJointTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedJoint{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(&test.model, test.model.MarriedJointTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
