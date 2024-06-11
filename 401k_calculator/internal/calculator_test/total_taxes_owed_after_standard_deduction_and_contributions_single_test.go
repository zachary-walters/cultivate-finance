package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsSingle) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedAfterStandardDeductionAndContributionsSingleTests = []struct {
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

func TestNewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsSingle()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

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

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsSingleTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{}

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsSingleTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

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

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsSingleTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsSingle{}

		actual := c.CalculateRothRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}
