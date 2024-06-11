package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAbstractIncomePerBracketAfterStandardDeduction struct {
	mock.Mock
}

func (m *MockAbstractIncomePerBracketAfterStandardDeduction) CalculateTraditional(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeduction) CalculateTraditionalRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeduction) CalculateRoth(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

func (m *MockAbstractIncomePerBracketAfterStandardDeduction) CalculateRothRetirement(model *calculator.Model, taxRates []calculator.TaxRate) []float64 {
	args := m.Called(model, taxRates)
	return args.Get(0).([]float64)
}

var abstractIncomePerBracketAfterStandardDeductionTests = []struct {
	name                         string
	model                        calculator.Model
	taxRates                     []calculator.TaxRate
	incomeAfterStandardDeduction float64
	incomePerBracket             float64
}{
	{
		name: "Test Case 0",
		taxRates: []calculator.TaxRate{
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
		incomeAfterStandardDeduction: 2000.0,
		incomePerBracket:             3943.0,
	},
}

func TestNewAbstractIncomePerBracketAfterStandardDeduction(t *testing.T) {
	actual := calculator.NewAbstractIncomePerBracketAfterStandardDeduction()
	expected := calculator.AbstractIncomePerBracketAfterStandardDeduction{
		IncomeAfterStandardDeductionCalculation: calculator.NewIncomeAfterStandardDeduction(),
		AbstractIncomePerBracketCalculation:     calculator.NewAbstractIncomePerBracket(),
	}

	assert.Equal(t, expected, actual)
}

func TestAbstractIncomePerBracketAfterStandardDeductionCalculateTraditional(t *testing.T) {
	for _, test := range abstractIncomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)
			mockIncomeAfterStandardDeduction.On("CalculateTraditional", &test.model).Return(test.incomeAfterStandardDeduction)

			mockAbstractIncomePerBracket := new(MockAbstractIncomePerBracket)
			mockAbstractIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.AbstractIncomePerBracketAfterStandardDeduction{
				IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
				AbstractIncomePerBracketCalculation:     mockAbstractIncomePerBracket,
			}

			actual := c.CalculateTraditional(&test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAbstractIncomePerBracketAfterStandardDeductionCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range abstractIncomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)
			mockIncomeAfterStandardDeduction.On("CalculateTraditionalRetirement", &test.model).Return(test.incomeAfterStandardDeduction)

			mockAbstractIncomePerBracket := new(MockAbstractIncomePerBracket)
			mockAbstractIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.AbstractIncomePerBracketAfterStandardDeduction{
				IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
				AbstractIncomePerBracketCalculation:     mockAbstractIncomePerBracket,
			}

			actual := c.CalculateTraditionalRetirement(&test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAbstractIncomePerBracketAfterStandardDeductionCalculateRoth(t *testing.T) {
	for _, test := range abstractIncomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)
			mockIncomeAfterStandardDeduction.On("CalculateRoth", &test.model).Return(test.incomeAfterStandardDeduction)

			mockAbstractIncomePerBracket := new(MockAbstractIncomePerBracket)
			mockAbstractIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.AbstractIncomePerBracketAfterStandardDeduction{
				IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
				AbstractIncomePerBracketCalculation:     mockAbstractIncomePerBracket,
			}

			actual := c.CalculateRoth(&test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAbstractIncomePerBracketAfterStandardDeductionCalculateRothRetirement(t *testing.T) {
	for _, test := range abstractIncomePerBracketAfterStandardDeductionTests {
		t.Run(test.name, func(t *testing.T) {
			mockIncomeAfterStandardDeduction := new(MockIncomeAfterStandardDeduction)
			mockIncomeAfterStandardDeduction.On("CalculateRothRetirement", &test.model).Return(test.incomeAfterStandardDeduction)

			mockAbstractIncomePerBracket := new(MockAbstractIncomePerBracket)
			mockAbstractIncomePerBracket.On("Calculate", test.taxRates, mock.Anything, test.incomeAfterStandardDeduction).Return(test.incomePerBracket)

			c := calculator.AbstractIncomePerBracketAfterStandardDeduction{
				IncomeAfterStandardDeductionCalculation: mockIncomeAfterStandardDeduction,
				AbstractIncomePerBracketCalculation:     mockAbstractIncomePerBracket,
			}

			actual := c.CalculateRothRetirement(&test.model, test.taxRates)
			expected := []float64{}

			for range test.taxRates {
				expected = append(expected, test.incomePerBracket)
			}

			assert.Equal(t, expected, actual)
		})
	}
}
