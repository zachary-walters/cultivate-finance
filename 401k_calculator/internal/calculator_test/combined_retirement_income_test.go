package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockCombinedRetirementIncome struct {
	mock.Mock
}

func (m *MockCombinedRetirementIncome) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCombinedRetirementIncome) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCombinedRetirementIncome) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCombinedRetirementIncome) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var combinedRetirementIncomeTests = []struct {
	name                 string
	adjustedGrossIncome  float64
	halfOfSocialSecurity float64
	model                calculator.Model
}{
	{
		name:                 "Test Case 0",
		adjustedGrossIncome:  145943,
		halfOfSocialSecurity: 903405,
	},
	{
		name:                 "Test Case 1",
		adjustedGrossIncome:  0,
		halfOfSocialSecurity: 903405,
	},
	{
		name:                 "Test Case 1",
		adjustedGrossIncome:  145943,
		halfOfSocialSecurity: 0,
	},
	{
		name:                 "Test Case 1",
		adjustedGrossIncome:  math.MaxFloat64,
		halfOfSocialSecurity: math.MaxFloat64,
	},
	{
		name:                 "Test Case 1",
		adjustedGrossIncome:  -math.MaxFloat64,
		halfOfSocialSecurity: -math.MaxFloat64,
	},
}

func TestNewCombinedRetirementIncome(t *testing.T) {
	actual := calculator.NewCombinedRetirementIncome()
	expected := calculator.CombinedRetirementIncome{
		AdjustedGrossIncomeCalculation:  calculator.NewAdjustedGrossIncome(),
		HalfOfSocialSecurityCalculation: calculator.NewHalfOfSocialSecurity(),
	}

	assert.Equal(t, expected, actual)
}

func TestCombinedRetirementIncomeCalculateTraditional(t *testing.T) {
	for _, test := range combinedRetirementIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.CombinedRetirementIncome{}

			actual := c.CalculateTraditional(&test.model)

			assert.Zero(t, actual)
		})
	}
}

func TestCombinedRetirementIncomeCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range combinedRetirementIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateTraditionalRetirement", &test.model).Return(test.adjustedGrossIncome)

			mockHalfOfSocialSecurity := new(MockHalfOfSocialSecurity)
			mockHalfOfSocialSecurity.On("CalculateTraditionalRetirement", &test.model).Return(test.halfOfSocialSecurity)

			c := &calculator.CombinedRetirementIncome{
				AdjustedGrossIncomeCalculation:  mockAdjustedGrossIncome,
				HalfOfSocialSecurityCalculation: mockHalfOfSocialSecurity,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := test.adjustedGrossIncome + test.halfOfSocialSecurity

			assert.Equal(t, actual, expected)
		})
	}
}

func TestCombinedRetirementIncomeCalculateRoth(t *testing.T) {
	for _, test := range combinedRetirementIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.CombinedRetirementIncome{}

			actual := c.CalculateRoth(&test.model)

			assert.Zero(t, actual)
		})
	}
}

func TestCombinedRetirementIncomeCalculateRothRetirement(t *testing.T) {
	for _, test := range combinedRetirementIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateRothRetirement", &test.model).Return(test.adjustedGrossIncome)

			mockHalfOfSocialSecurity := new(MockHalfOfSocialSecurity)
			mockHalfOfSocialSecurity.On("CalculateRothRetirement", &test.model).Return(test.halfOfSocialSecurity)

			c := &calculator.CombinedRetirementIncome{
				AdjustedGrossIncomeCalculation:  mockAdjustedGrossIncome,
				HalfOfSocialSecurityCalculation: mockHalfOfSocialSecurity,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := test.adjustedGrossIncome + test.halfOfSocialSecurity

			assert.Equal(t, actual, expected)
		})
	}
}
