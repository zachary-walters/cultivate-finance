package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockIncomeAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockIncomeAfterStandardDeduction) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockIncomeAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var incomeAfterStandardDeductionTests = []struct {
	name               string
	model              calculator.Model
	standardDeduction  float64
	totalTaxableIncome float64
}{
	{
		name: "Test Case Basic",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAnnualIncome: 600000,
			},
		},
		standardDeduction:  32984.0,
		totalTaxableIncome: 9503.0,
	},
	{
		name: "Test Case Negative",
		model: calculator.Model{
			Input: calculator.Input{
				CurrentAnnualIncome: -100000,
			},
		},
		standardDeduction:  -320984.0,
		totalTaxableIncome: -2947.0,
	},
}

func TestNewIncomeAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewIncomeAfterStandardDeduction()
	expected := calculator.IncomeAfterStandardDeduction{
		StandardDeductionCalculation:  calculator.NewStandardDeduction(),
		TotalTaxableIncomeCalculation: calculator.NewTotalTaxableIncome(),
	}

	assert.Equal(t, expected, actual)
}

func TestIncomeAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range incomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockStandardDeduction := new(MockStandardDeduction)
			mockStandardDeduction.On("CalculateTraditional", mock.Anything).Return(test.standardDeduction)

			c := &calculator.IncomeAfterStandardDeduction{
				StandardDeductionCalculation: mockStandardDeduction,
			}

			result := c.CalculateTraditional(&test.model)
			expected := func() float64 {
				if test.model.Input.CurrentAnnualIncome-test.standardDeduction <= 0 {
					return 0.0
				}
				return test.model.Input.CurrentAnnualIncome - test.standardDeduction
			}()
			assert.Equal(t, expected, result)
		})
	}
}

func TestIncomeAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range incomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockStandardDeduction := new(MockStandardDeduction)
			mockStandardDeduction.On("CalculateTraditionalRetirement", mock.Anything).Return(test.standardDeduction)

			mockTotalTaxableIncome := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateTraditionalRetirement", mock.Anything).Return(test.totalTaxableIncome)

			c := &calculator.IncomeAfterStandardDeduction{
				StandardDeductionCalculation:  mockStandardDeduction,
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
			}

			result := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				if test.totalTaxableIncome-test.standardDeduction <= 0 {
					return 0.0
				}
				return test.totalTaxableIncome - test.standardDeduction
			}()
			assert.Equal(t, expected, result)
		})
	}
}

func TestIncomeAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range incomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockStandardDeduction := new(MockStandardDeduction)
			mockStandardDeduction.On("CalculateRoth", mock.Anything).Return(test.standardDeduction)

			c := &calculator.IncomeAfterStandardDeduction{
				StandardDeductionCalculation: mockStandardDeduction,
			}

			result := c.CalculateRoth(&test.model)
			expected := func() float64 {
				if test.model.Input.CurrentAnnualIncome-test.standardDeduction <= 0 {
					return 0.0
				}
				return test.model.Input.CurrentAnnualIncome - test.standardDeduction
			}()
			assert.Equal(t, expected, result)
		})
	}
}

func TestIncomeAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range incomeAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockStandardDeduction := new(MockStandardDeduction)
			mockStandardDeduction.On("CalculateRothRetirement", mock.Anything).Return(test.standardDeduction)

			mockTotalTaxableIncome := new(MockStandardDeduction)
			mockTotalTaxableIncome.On("CalculateRothRetirement", mock.Anything).Return(test.totalTaxableIncome)

			c := &calculator.IncomeAfterStandardDeduction{
				StandardDeductionCalculation:  mockStandardDeduction,
				TotalTaxableIncomeCalculation: mockTotalTaxableIncome,
			}

			result := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				if test.totalTaxableIncome-test.standardDeduction <= 0 {
					return 0.0
				}
				return test.totalTaxableIncome - test.standardDeduction
			}()
			assert.Equal(t, expected, result)
		})
	}
}
