package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedSeparate) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests = []struct {
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

func TestNewTotalTaxesOwedAfterStandardDeductionMarriedSeparate(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedSeparate()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
		TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.CalculateTraditionalRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.CalculateRoth(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRothRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedSeparate.CalculateRothRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
