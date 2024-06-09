package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests = []struct {
	name                                                  string
	model                                                 calculator.Model
	incomePerBracketAfterStandardDeductionHeadOfHousehold []float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			HeadOfHouseholdTaxRates: []calculator.TaxRate{
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
		incomePerBracketAfterStandardDeductionHeadOfHousehold: []float64{
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
		IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: calculator.NewIncomePerBracketAfterStandardDeductionHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionHeadOfHouseholdCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
