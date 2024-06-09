package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandadDeductionAndContributionsMarriedSeparateTests = []struct {
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
		IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsMarriedSeparateTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate)
			mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionMarriedSeparate)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate{
				IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionMarriedSeparate := c.IncomePerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.MarriedSeparateTaxRates))

			for idx, taxRate := range test.model.MarriedSeparateTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionMarriedSeparate[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
