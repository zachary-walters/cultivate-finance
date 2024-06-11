package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateTests = []struct {
	name                                                      string
	model                                                     calculator.Model
	taxesOwedPerBracketAfterStandardDeductionAndContributions []float64
}{
	{
		name:  "Test Case 0",
		model: calculator.Model{},
		taxesOwedPerBracketAfterStandardDeductionAndContributions: []float64{
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		},
	},
	{
		name: "Test Case 1",
	},
}

func TestNewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			for _, value := range test.taxesOwedPerBracketAfterStandardDeductionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{}

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparate.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			for _, value := range test.taxesOwedPerBracketAfterStandardDeductionAndContributions {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeparateCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparateTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedSeparate{}

		actual := c.CalculateRothRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}
