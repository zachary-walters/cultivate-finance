package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionMarriedSeperate struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionMarriedSeperate) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandardDeductionMarriedSeperateTests = []struct {
	name                                   string
	model                                  calculator.Model
	incomePerBracketAfterStandardDeduction []float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			MarriedSeperateTaxRates: []calculator.TaxRate{
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionMarriedSeperate(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionMarriedSeperate()
	expected := calculator.IncomePerBracketAfterStandardDeductionMarriedSeperate{
		AbstractIncomePerBracketAfterStandardDeductionCalculation: calculator.NewAbstractIncomePerBracketAfterStandardDeduction(),
	}

	assert.Equal(t, expected, actual)
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeperateCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeperateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditional", &test.model, test.model.MarriedSeperateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeperate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditional(&test.model, test.model.MarriedSeperateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeperateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeperateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model, test.model.MarriedSeperateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeperate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateTraditionalRetirement(&test.model, test.model.MarriedSeperateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeperateCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeperateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRoth", &test.model, test.model.MarriedSeperateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeperate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRoth(&test.model, test.model.MarriedSeperateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockIncomePerBracketAfterStandardDeductionMarriedSeperateCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionMarriedSeperateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeduction := new(MockAbstractIncomePerBracketAfterStandardDeduction)
			mockIncomePerBracketAfterStandardDeduction.On("CalculateRothRetirement", &test.model, test.model.MarriedSeperateTaxRates).Return(test.incomePerBracketAfterStandardDeduction)

			c := &calculator.IncomePerBracketAfterStandardDeductionMarriedSeperate{
				AbstractIncomePerBracketAfterStandardDeductionCalculation: mockIncomePerBracketAfterStandardDeduction,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionCalculation.CalculateRothRetirement(&test.model, test.model.MarriedSeperateTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
