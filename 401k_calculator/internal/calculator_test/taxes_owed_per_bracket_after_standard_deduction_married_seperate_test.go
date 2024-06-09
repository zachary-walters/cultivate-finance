package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests = []struct {
	name                                                  string
	model                                                 calculator.Model
	incomePerBracketAfterStandardDeductionMarriedSeparate []float64
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
		incomePerBracketAfterStandardDeductionMarriedSeparate: []float64{
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
		IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: calculator.NewIncomePerBracketAfterStandardDeductionMarriedSeparate(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
				IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeparate,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
				IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeparate,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
				IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeparate,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionMarriedSeparate{
				IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionMarriedSeparate,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionMarriedSeparateCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
