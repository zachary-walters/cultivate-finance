package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	mock.Mock
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

var incomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdTests = []struct {
	name                                                   string
	model                                                  calculator.Model
	incomePerBracketAfterStandardDeductionAndContributions []float64
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
	},
}

func TestNewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(t *testing.T) {
	actual := calculator.NewIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold()
	expected := calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
		AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: calculator.NewAbstractIncomePerBracketAfterStandardDeductionAndContributions(),
	}

	assert.Equal(t, expected, actual)
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditional(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateTraditional", &test.model, test.model.HeadOfHouseholdTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditional(&test.model, test.model.HeadOfHouseholdTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateTraditionalRetirement", &test.model, test.model.HeadOfHouseholdTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateTraditionalRetirement(&test.model, test.model.HeadOfHouseholdTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRoth(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateRoth", &test.model, test.model.HeadOfHouseholdTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRoth(&test.model, test.model.HeadOfHouseholdTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestIncomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRothRetirement(t *testing.T) {
	for _, test := range incomePerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomePerBracketAfterStandardDeductionAndContributions := new(MockAbstractIncomePerBracketAfterStandardDeductionAndContributions)
			mockIncomePerBracketAfterStandardDeductionAndContributions.On("CalculateRothRetirement", &test.model, test.model.HeadOfHouseholdTaxRates).Return(test.incomePerBracketAfterStandardDeductionAndContributions)

			c := &calculator.IncomePerBracketAfterStandardDeductionAndContributionsHeadOfHousehold{
				AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation: mockIncomePerBracketAfterStandardDeductionAndContributions,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.AbstractIncomePerBracketAfterStandardDeductionAndContributionsCalculation.CalculateRothRetirement(&test.model, test.model.HeadOfHouseholdTaxRates)

			assert.Equal(t, expected, actual)
		})
	}
}
