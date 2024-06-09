package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionMarriedSeparate struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandardDeductionMarriedSeparateTests = []struct {
	name                                   string
	model                                  calculator.Model
	incomePerBracketAfterStandardDeduction []float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			MarriedSeparateTaxRates: []calculator.TaxRate{
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

func NewIncomePerBracketAfterStandardDeductionMarriedSeparate(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionMarriedSeparate()
	expected := calculator.IncomePerBracketAfterStandardDeductionMarriedSeparate{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: calculator.NewAbstractIncomePerBracketAfterStandardDeduction(),
	}

	assert.Equal(t, expected, actual)
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditional", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRoth", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRothRetirement", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
