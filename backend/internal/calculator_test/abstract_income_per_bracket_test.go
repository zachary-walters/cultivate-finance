package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockAbstractIncomePerBracket struct {
	mock.Mock
}

func (m *MockAbstractIncomePerBracket) Calculate(taxRates []calculator.TaxRate, bracketSequence int, income float64) float64 {
	args := m.Called(taxRates, bracketSequence, income)
	return args.Get(0).(float64)
}

func TestIncomePerBracketCalculate(t *testing.T) {
	calc := &calculator.AbstractIncomePerBracket{}

	taxRates := []calculator.TaxRate{
		{Cap: 10000.0},
		{Cap: 30000.0},
		{Cap: 60000.0},
		{Cap: 100000.0},
	}

	tests := []struct {
		name            string
		taxRates        []calculator.TaxRate
		bracketSequence int
		income          float64
		expected        float64
	}{
		{
			name:            "Test Case Bracket 0",
			taxRates:        taxRates,
			bracketSequence: 0,
			income:          1000000.0,
			expected:        10000.0,
		},
		{
			name:            "Test Case Bracket 1",
			taxRates:        taxRates,
			bracketSequence: 1,
			income:          1000000.0,
			expected:        20000.0,
		},
		{
			name:            "Test Case Bracket 2",
			taxRates:        taxRates,
			bracketSequence: 2,
			income:          1000000.0,
			expected:        30000.0,
		},
		{
			name:            "Test Case Bracket 3",
			taxRates:        taxRates,
			bracketSequence: 3,
			income:          1000000.0,
			expected:        40000.0,
		},
		{
			name:            "Test Case Bracket 0 income lower than bracket",
			taxRates:        taxRates,
			bracketSequence: 0,
			income:          10,
			expected:        10,
		},
		{
			name:            "Test Case Bracket 1 overflow",
			taxRates:        taxRates,
			bracketSequence: 1,
			income:          10,
			expected:        0,
		},
		{
			name:            "Test Case Bracket 2 overflow",
			taxRates:        taxRates,
			bracketSequence: 2,
			income:          10,
			expected:        0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Calculate(tt.taxRates, tt.bracketSequence, tt.income)
			assert.Equal(t, tt.expected, result)
		})
	}

}

// These tests are more for QA than function logic
// These tests will test against known expected inputs and outputs
func TestIncomePerBracketCalculateQA(t *testing.T) {
	// 2023 tax Rates
	singleFilerTaxRates := []calculator.TaxRate{
		{
			Cap:  10275,
			Rate: 0.10,
		},
		{
			Cap:  41775,
			Rate: 0.12,
		},
		{
			Cap:  89075,
			Rate: 0.22,
		},
		{
			Cap:  170050,
			Rate: 0.24,
		},
		{
			Cap:  215950,
			Rate: 0.32,
		},
		{
			Cap:  539900,
			Rate: 0.35,
		},
		{
			Cap:  math.Inf(1),
			Rate: 0.37,
		},
	}

	marriedJointFilerTaxRates := []calculator.TaxRate{
		{
			Cap:  20550,
			Rate: 0.10,
		},
		{
			Cap:  83550,
			Rate: 0.12,
		},
		{
			Cap:  178150,
			Rate: 0.22,
		},
		{
			Cap:  340100,
			Rate: 0.24,
		},
		{
			Cap:  431900,
			Rate: 0.32,
		},
		{
			Cap:  647850,
			Rate: 0.35,
		},
		{
			Cap:  math.Inf(1),
			Rate: 0.37,
		},
	}

	incomeAfterStandardDeduction100000 := 100000.0 - 13850.0
	incomeAfterStandardDeduction10000000 := 10000000.0 - 27700.0
	tests := []struct {
		name            string
		taxRates        []calculator.TaxRate
		income          float64
		bracketSequence int
		expected        float64
	}{
		{
			name:            "QA Test Case 2023 Single Filer Rates at $100000 income bracket 0",
			taxRates:        singleFilerTaxRates,
			income:          incomeAfterStandardDeduction100000,
			bracketSequence: 0,
			expected:        10275,
		},
		{
			name:            "QA Test Case 2023 Single Filer Rates at $100000 income bracket 1",
			taxRates:        singleFilerTaxRates,
			income:          incomeAfterStandardDeduction100000,
			bracketSequence: 1,
			expected:        31500,
		},
		{
			name:            "QA Test Case 2023 Single Filer Rates at $100000 income bracket 2",
			taxRates:        singleFilerTaxRates,
			income:          incomeAfterStandardDeduction100000,
			bracketSequence: 2,
			expected:        44375,
		},
		{
			name:            "QA Test Case 2023 Single Filer Rates at $100000 income bracket 3",
			taxRates:        singleFilerTaxRates,
			income:          incomeAfterStandardDeduction100000,
			bracketSequence: 3,
			expected:        0,
		},
		{
			name:            "QA Test Case 2023 Single Filer Rates at $100000 income bracket 4",
			taxRates:        singleFilerTaxRates,
			income:          incomeAfterStandardDeduction100000,
			bracketSequence: 4,
			expected:        0,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 0",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 0,
			expected:        20550,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 1",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 1,
			expected:        63000,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 2",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 2,
			expected:        94600,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 3",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 3,
			expected:        161950,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 4",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 4,
			expected:        91800,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 5",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 5,
			expected:        215950,
		},
		{
			name:            "QA Test Case 2023 Married Joint Filer Rates at $10000000 income bracket 6",
			taxRates:        marriedJointFilerTaxRates,
			income:          incomeAfterStandardDeduction10000000,
			bracketSequence: 6,
			expected:        9324450,
		},
	}

	calc := &calculator.AbstractIncomePerBracket{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calc.Calculate(test.taxRates, test.bracketSequence, test.income)
			assert.Equal(t, test.expected, actual)
		})
	}
}
