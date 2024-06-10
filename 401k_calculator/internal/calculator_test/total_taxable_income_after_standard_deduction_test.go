package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxableIncomeAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockTotalTaxableIncomeAfterStandardDeduction) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncomeAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncomeAfterStandardDeduction) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncomeAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxableIncomeAfterStandardDeductionTests = []struct {
	name               string
	model              calculator.Model
	totalTaxableIncome float64
	standardDeduction  float64
}{
	{
		name:               "Test Case 0",
		totalTaxableIncome: 10000,
		standardDeduction:  200,
	},
}

func TestNewTotalTaxableIncomeAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewTotalTaxableIncomeAfterStandardDeduction()
	expected := calculator.TotalTaxableIncomeAfterStandardDeduction{
		TotalTaxableIncomeCalculation: calculator.NewTotalTaxableIncome(),
		StandardDeductionCalculation:  calculator.NewStandardDeduction(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxableIncomeAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxableIncomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxableIncome := new(MockTotalTaxableIncome)
			mockStandardDeduction := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateTraditional", &test.model).Return(test.totalTaxableIncome)
			mockStandardDeduction.On("CalculateTraditional", &test.model).Return(test.standardDeduction)

			c := &calculator.TotalTaxableIncomeAfterStandardDeduction{
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
				StandardDeductionCalculation:  mockStandardDeduction,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := test.totalTaxableIncome - test.standardDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxableIncomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxableIncome := new(MockTotalTaxableIncome)
			mockStandardDeduction := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateTraditionalRetirement", &test.model).Return(test.totalTaxableIncome)
			mockStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.standardDeduction)

			c := &calculator.TotalTaxableIncomeAfterStandardDeduction{
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
				StandardDeductionCalculation:  mockStandardDeduction,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := test.totalTaxableIncome - test.standardDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range totalTaxableIncomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxableIncome := new(MockTotalTaxableIncome)
			mockStandardDeduction := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateRoth", &test.model).Return(test.totalTaxableIncome)
			mockStandardDeduction.On("CalculateRoth", &test.model).Return(test.standardDeduction)

			c := &calculator.TotalTaxableIncomeAfterStandardDeduction{
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
				StandardDeductionCalculation:  mockStandardDeduction,
			}

			actual := c.CalculateRoth(&test.model)
			expected := test.totalTaxableIncome - test.standardDeduction

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxableIncomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalTaxableIncome := new(MockTotalTaxableIncome)
			mockStandardDeduction := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateRothRetirement", &test.model).Return(test.totalTaxableIncome)
			mockStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.standardDeduction)

			c := &calculator.TotalTaxableIncomeAfterStandardDeduction{
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
				StandardDeductionCalculation:  mockStandardDeduction,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := test.totalTaxableIncome - test.standardDeduction

			assert.Equal(t, expected, actual)
		})
	}
}
