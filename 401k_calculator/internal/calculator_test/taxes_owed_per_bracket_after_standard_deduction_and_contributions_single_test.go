package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandadDeductionAndContributionsSingleTests = []struct {
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
		IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsSingle(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsSingleTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle := new(MockIncomePerBracketAfterStandardDeductionAndContributionsSingle)
			mockIncomePerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionSingle)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle{
				IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsSingle,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionSingle := c.IncomePerBracketAfterStandardDeductionAndContributionsSingleCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.SingleTaxRates))

			for idx, taxRate := range test.model.SingleTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionSingle[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
