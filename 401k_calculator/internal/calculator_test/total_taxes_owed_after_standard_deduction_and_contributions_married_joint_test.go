package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointTests = []struct {
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

func TestNewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint{
		TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

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

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint{}

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributions)

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

func TestTotalTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedAfterStandardDeductionAndContributionsMarriedJointTests {
		c := &calculator.TotalTaxesOwedAfterStandardDeductionAndContributionsMarriedJoint{}

		actual := c.CalculateRothRetirement(&test.model)
		expected := 0.0

		assert.Equal(t, expected, actual)
	}
}
