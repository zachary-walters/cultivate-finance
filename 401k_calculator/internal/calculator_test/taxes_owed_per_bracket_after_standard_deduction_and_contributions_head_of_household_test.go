package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var taxesOwedPerBracketAfterStandadDeductionAndContributionsHeadOfHouseholdTests = []struct {
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

func TestNewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(t *testing.T) {
	actual := calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold()
	expected := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
		IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditional(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditional(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation.CalculateTraditional(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation.CalculateTraditionalRetirement(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRoth(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRoth(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation.CalculateRoth(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRothRetirement(t *testing.T) {
	for _, test := range taxesOwedPerBracketAfterStandadDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)
			mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.incomePerBracketAfterStandardDeductionHeadOfHousehold)

			c := &calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.CalculateRothRetirement(&test.model)

			incomePerBracketAfterStandarddeductionHeadOfHousehold := c.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation.CalculateRothRetirement(&test.model)
			expected := make([]float64, len(test.model.HeadOfHouseholdTaxRates))

			for idx, taxRate := range test.model.HeadOfHouseholdTaxRates {
				expected[idx] = incomePerBracketAfterStandarddeductionHeadOfHousehold[idx] * taxRate.Rate
			}

			assert.Equal(t, expected, actual)
		})
	}
}
