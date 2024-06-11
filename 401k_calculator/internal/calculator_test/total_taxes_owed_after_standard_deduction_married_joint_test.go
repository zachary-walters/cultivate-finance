package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxesOwedAfterStandardDeductionMarriedJoint struct {
	mock.Mock
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxesOwedAfterStandardDeductionMarriedJoint) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxesOwedPerBracketAfterStandardDeductionMarriedJointTests = []struct {
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

func TestNewTotalTaxesOwedAfterStandardDeductionMarriedJoint(t *testing.T) {
	actual := calculator.NewTotalTaxesOwedAfterStandardDeductionMarriedJoint()
	expected := calculator.TotalTaxesOwedAfterStandardDeductionMarriedJoint{
		TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: calculator.NewTaxesOwedPerBracketAfterStandardDeductionMarriedJoint(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateTraditional(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.On("CalculateTraditional", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditional(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.CalculateTraditional(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateTraditionalRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.CalculateTraditionalRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateRoth(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.On("CalculateRoth", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRoth(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.CalculateRoth(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}

func TestTotalTaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculateRothRetirement(t *testing.T) {
	for _, test := range totalTaxesOwedPerBracketAfterStandardDeductionMarriedJointTests {
		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionMarriedJoint)

		c := &calculator.TotalTaxesOwedAfterStandardDeductionMarriedJoint{
			TaxesOwedPerBracketAfterStandardDeductionMarriedJointCalculation: mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint,
		}

		mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.On("CalculateRothRetirement", &test.model).Return(test.taxesOwedPerBracketAfterStandardDeductions)

		actual := c.CalculateRothRetirement(&test.model)
		expected := func() float64 {
			expected := 0.0
			taxesOwedPerBracketAfterStandardDudections := mockTaxesOwedPerBracketAfterStandardDeductionsMarriedJoint.CalculateRothRetirement(&test.model)
			for _, value := range taxesOwedPerBracketAfterStandardDudections {
				expected += value
			}

			return expected
		}()

		assert.Equal(t, expected, actual)
	}
}
