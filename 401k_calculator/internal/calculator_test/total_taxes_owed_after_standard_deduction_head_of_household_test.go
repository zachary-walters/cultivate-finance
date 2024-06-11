package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionHeadOfHousehold) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests = []struct {
	name                                       string
	model                                      calculator.Model
	taxesOwedPerBracketAfterStandardDeductions []float64
}{
	{
		name:  "Test Case 0",
		model: calculator.Model{},
		taxesOwedPerBracketAfterStandardDeductions: []float64{
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		},
	},
}

func TestNewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionHeadOfHousehold()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
		TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.CalculateTraditionalRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.CalculateRoth(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionHeadOfHousehold)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionHeadOfHousehold{
			TaxesOwedPerBracketAfterStandardDeductionHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRothRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsHeadOfHousehold.CalculateRothRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
