package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdTests = []struct {
	name                                                      string
	model                                                     calculator.Model
	taxesOwedPerBracketAfterStandardDeductionAndContributions []float64
}{
	{
		name:  "Test Case 0",
		model: calculator.Model{},
		taxesOwedPerBracketAfterStandardDeductionAndContributions: []float64{
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		},
	},
	{
		name: "Test Case 1",
	},
}

func TestNewTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			for _, value := range test.taxesOwedPerBracketAfterStandardDeductionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold{}

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			for _, value := range test.taxesOwedPerBracketAfterStandardDeductionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHouseholdTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsHeadOfHousehold{}

		actual := c.CalculateRothRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}
