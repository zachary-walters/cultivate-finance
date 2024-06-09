package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandarddeductionAndContributionsMarriedSeparateTests = []struct {
	name                                                   string
	model                                                  calculator.Model
	incomePerBracketAfterStandardDeductionAndContributions []float64
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
		incomePerBracketAfterStandardDeductionAndContributions: []float64{912, 3124, 6, 346, 0},
	},
	{
		name: "Test Case 1",
		model: calculator.Model{
			MarriedSeparateTaxRates: []calculator.TaxRate{
				{
					Cap:  3.0,
					Rate: 0.1,
				},
				{
					Cap:  2.0,
					Rate: 0.0,
				},
				{
					Cap:  1.0,
					Rate: 0.4,
				},
			},
		},
		incomePerBracketAfterStandardDeductionAndContributions: []float64{1},
	},
	{
		name: "Test Case 2",
	},
}

func TestNewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate()
	expected := calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
		AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: calculator.NewAbstractIncomePerBracketAfterStandardDeductionAndContributions(),
	}

	assert.Equal(t, expected, actual)
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandarddeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateTraditional", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandarddeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateTraditionalRetirement", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditionalRetirement(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandarddeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateRoth", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRoth(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandarddeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateRothRetirement", &test.model, test.model.MarriedSeparateTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRothRetirement(&test.model, test.model.MarriedSeparateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
