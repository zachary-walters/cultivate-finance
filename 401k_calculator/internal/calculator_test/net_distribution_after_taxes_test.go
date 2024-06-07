package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockNetDistributionAfterTaxes struct {
	mock.Mock
}

func (m *MockNetDistributionAfterTaxes) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockNetDistributionAfterTaxes) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockNetDistributionAfterTaxes) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockNetDistributionAfterTaxes) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var netDistributionAfterTaxesTests = []struct {
	name                                 string
	totalTaxesOwedAfterStandadDeduction  float64
	totalAnnualRetirementIncomeBeforeTax float64
}{
	{
		name:                                 "Test Case 0",
		totalTaxesOwedAfterStandadDeduction:  1337,
		totalAnnualRetirementIncomeBeforeTax: 3942039,
	},
	{
		name:                                 "Test Case 1",
		totalTaxesOwedAfterStandadDeduction:  1337,
		totalAnnualRetirementIncomeBeforeTax: 3447,
	},
	{
		name: "Test Case 2",
	},
}

func TestNewNetDistributionAfterTaxes(t *testing.T) {
	actual := calculator.NewNetDistributionAfterTaxes()
	expected := calculator.NetDistributionAfterTaxes{
		TotalTaxesOwedAfterStandardDeductionCalculation: calculator.NewTotalTaxesOwedAfterStandardDeduction(),
		TotalAnnualRetirementIncomeBeforeTaxCalculation: calculator.NewTotalAnnualRetirementIncomeBeforeTax(),
	}

	assert.Equal(t, expected, actual)
}

func TestNetDistributionAfterTaxesCalculateTraditional(t *testing.T) {
	testModel := calculator.Model{}

	for _, test := range netDistributionAfterTaxesTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditional", &testModel).Return(test.totalTaxesOwedAfterStandadDeduction)

			mockTotalAnnualRetirementIncomeBeforeTaxCalculation := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTaxCalculation.On("CalculateTraditional", &testModel).Return(test.totalAnnualRetirementIncomeBeforeTax)

			c := calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTaxCalculation,
			}

			actual := c.CalculateTraditional(&testModel)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.totalTaxesOwedAfterStandadDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestNetDistributionAfterTaxesCalculateTraditionalRetirement(t *testing.T) {
	testModel := calculator.Model{}

	for _, test := range netDistributionAfterTaxesTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &testModel).Return(test.totalTaxesOwedAfterStandadDeduction)

			mockTotalAnnualRetirementIncomeBeforeTaxCalculation := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTaxCalculation.On("CalculateTraditionalRetirement", &testModel).Return(test.totalAnnualRetirementIncomeBeforeTax)

			c := calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTaxCalculation,
			}

			actual := c.CalculateTraditionalRetirement(&testModel)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.totalTaxesOwedAfterStandadDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestNetDistributionAfterTaxesCalculateRoth(t *testing.T) {
	testModel := calculator.Model{}

	for _, test := range netDistributionAfterTaxesTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRoth", &testModel).Return(test.totalTaxesOwedAfterStandadDeduction)

			mockTotalAnnualRetirementIncomeBeforeTaxCalculation := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTaxCalculation.On("CalculateRoth", &testModel).Return(test.totalAnnualRetirementIncomeBeforeTax)

			c := calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTaxCalculation,
			}

			actual := c.CalculateRoth(&testModel)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.totalTaxesOwedAfterStandadDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestNetDistributionAfterTaxesCalculateRothRetirement(t *testing.T) {
	testModel := calculator.Model{}

	for _, test := range netDistributionAfterTaxesTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &testModel).Return(test.totalTaxesOwedAfterStandadDeduction)

			mockTotalAnnualRetirementIncomeBeforeTaxCalculation := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTaxCalculation.On("CalculateRothRetirement", &testModel).Return(test.totalAnnualRetirementIncomeBeforeTax)

			c := calculator.NetDistributionAfterTaxes{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTaxCalculation,
			}

			actual := c.CalculateRothRetirement(&testModel)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.totalTaxesOwedAfterStandadDeduction

			assert.Equal(t, expected, actual)
		})
	}
}
