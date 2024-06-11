package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionSingle struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionSingle) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionSingle) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionSingle) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionSingle) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedPerBracketAfterStandardDeductionSingleTests = []struct {
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

func TestNewTotalTaxesOwedAfterStandardDeductionSingle(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionSingle()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionSingle{
		TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionSingle(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionSingleCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionSingle{
			TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsSingle.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsSingle.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionSingleCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionSingle{
			TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsSingle.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsSingle.CalculateTraditionalRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionSingleCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionSingle{
			TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsSingle.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsSingle.CalculateRoth(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionSingleCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionSingle{
			TaxesOwedPerBracketAfterStandardDeductionSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsSingle.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRothRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsSingle.CalculateRothRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
