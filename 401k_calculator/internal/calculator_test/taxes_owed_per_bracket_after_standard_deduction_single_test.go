package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionSingle struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionSingle) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandardDeductionSingleTests = []struct {
	name                                         string
	model                                        calculator.Model
	incomePerBracketAfterStandardDeductionSingle []float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
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
		incomePerBracketAfterStandardDeductionSingle: []float64{
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionSingle(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
		IncomePerBracketAfterStandardDeductionSingleCalculation: calculator.NewIncomePerBracketAfterStandardDeductionSingle(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionSingleCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionSingleCalculation: mockIncomePerBracketAfterStandardDeductionSingle,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionSingleCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionSingleCalculation: mockIncomePerBracketAfterStandardDeductionSingle,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionSingleCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionSingleCalculation: mockIncomePerBracketAfterStandardDeductionSingle,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionSingleCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionSingle := new(MockIncomePerBracketAfterStandardDeductionSingle)
			mockIncomePerBracketAfterStandardDeductionSingle.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionSingle{
				IncomePerBracketAfterStandardDeductionSingleCalculation: mockIncomePerBracketAfterStandardDeductionSingle,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionSingleCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
