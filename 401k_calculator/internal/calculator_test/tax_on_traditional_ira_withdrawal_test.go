package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxOnTraditionalIRAWithdrawal struct {
	mock.Mock
}

func (m *MockTaxOnTraditionalIRAWithdrawal) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxOnTraditionalIRAWithdrawal) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxOnTraditionalIRAWithdrawal) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTaxOnTraditionalIRAWithdrawal) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var taxOnTraditionalIRAWithdrawalTests = []struct {
	name                                            string
	model                                           calculator.Model
	totalTaxesOwedAfterStandardDeductionTraditional float64
	totalTaxesOwedAfterStandardDeductionRoth        float64
}{
	{
		name: "Test Case 0",
		totalTaxesOwedAfterStandardDeductionTraditional: 934830,
		totalTaxesOwedAfterStandardDeductionRoth:        543,
	},
}

func TestNewTaxOnTraditionalIRAWithdrawal(t *testing.T) {
	actual := calculator.NewTaxOnTraditionalIRAWithdrawal()
	expected := calculator.TaxOnTraditionalIRAWithdrawal{
		TotalTaxesOwedAfterStandardDeductionCalculation: calculator.NewTotalTaxesOwedAfterStandardDeduction(),
	}

	assert.Equal(t, expected, actual)
}

func TestTaxOnTraditionalIRAWithdrawalCalculateTraditional(t *testing.T) {
	for _, test := range taxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionTraditional)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionRoth)

			c := &calculator.TaxOnTraditionalIRAWithdrawal{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxOnTraditionalIRAWithdrawalCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range taxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionTraditional)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionRoth)

			c := &calculator.TaxOnTraditionalIRAWithdrawal{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := test.totalTaxesOwedAfterStandardDeductionTraditional - test.totalTaxesOwedAfterStandardDeductionRoth

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxOnTraditionalIRAWithdrawalCalculateRoth(t *testing.T) {
	for _, test := range taxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionTraditional)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionRoth)

			c := &calculator.TaxOnTraditionalIRAWithdrawal{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateRoth(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTaxOnTraditionalIRAWithdrawalCalculateRothRetirement(t *testing.T) {
	for _, test := range taxOnTraditionalIRAWithdrawalTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxesOwedAfterStandardDeduction := new(MockTotalTaxesOwedAfterStandardDeduction)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionTraditional)
			mockTotalTaxesOwedAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.totalTaxesOwedAfterStandardDeductionRoth)

			c := &calculator.TaxOnTraditionalIRAWithdrawal{
				TotalTaxesOwedAfterStandardDeductionCalculation: mockTotalTaxesOwedAfterStandardDeduction,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
