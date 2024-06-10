package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal struct {
	mock.Mock
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalTests = []struct {
	name                                 string
	model                                calculator.Model
	totalAnnualRetirementIncomeBeforeTax float64
	taxOnTraditionalIRAWithdrawal        float64
}{
	{
		name: "Test Case 0",
	},
}

func TestNewTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal(t *testing.T) {
	actual := calculator.NewTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal()
	expected := calculator.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal{
		TotalAnnualRetirementIncomeBeforeTaxCalculation: calculator.NewTotalAnnualRetirementIncomeBeforeTax(),
		TaxOnTraditionalIRAWithdrawalCalculation:        calculator.NewTaxOnTraditionalIRAWithdrawal(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculateTraditional(t *testing.T) {
	for _, test := range totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalAnnualRetirementIncomeBeforeTax := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTax.On("CalculateTraditional", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTax)

			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateTraditional", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := &calculator.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal{
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTax,
				TaxOnTraditionalIRAWithdrawalCalculation:        mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.taxOnTraditionalIRAWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalAnnualRetirementIncomeBeforeTax := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTax.On("CalculateTraditionalRetirement", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTax)

			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateTraditionalRetirement", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := &calculator.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal{
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTax,
				TaxOnTraditionalIRAWithdrawalCalculation:        mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.taxOnTraditionalIRAWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculateRoth(t *testing.T) {
	for _, test := range totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalAnnualRetirementIncomeBeforeTax := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTax.On("CalculateRoth", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTax)

			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateRoth", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := &calculator.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal{
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTax,
				TaxOnTraditionalIRAWithdrawalCalculation:        mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateRoth(&test.model)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.taxOnTraditionalIRAWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalCalculateRothRetirement(t *testing.T) {
	for _, test := range totalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalAnnualRetirementIncomeBeforeTax := new(MockTotalAnnualRetirementIncomeBeforeTax)
			mockTotalAnnualRetirementIncomeBeforeTax.On("CalculateRothRetirement", &test.model).Return(test.totalAnnualRetirementIncomeBeforeTax)

			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateRothRetirement", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := &calculator.TotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal{
				TotalAnnualRetirementIncomeBeforeTaxCalculation: mockTotalAnnualRetirementIncomeBeforeTax,
				TaxOnTraditionalIRAWithdrawalCalculation:        mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := test.totalAnnualRetirementIncomeBeforeTax - test.taxOnTraditionalIRAWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}
