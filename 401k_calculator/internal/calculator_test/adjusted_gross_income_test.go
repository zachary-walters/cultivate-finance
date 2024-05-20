package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAdjustedGrossIncome struct {
	mock.Mock
}

func (m *MockAdjustedGrossIncome) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncome) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncome) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncome) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var adjustedGrossIncomeTests = []struct {
	name  string
	model calculator.Model
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			Input: calculator.Input{
				PensionIncome:    9492875,
				WorkIncome:       5293857,
				RentalNetIncome:  40394,
				AnnuityIncome:    5075,
				YearlyWithdrawal: 9405097426,
			},
		},
	},
	{
		name: "Test Case 1",
	},
	{
		name: "Test Case 2",
		model: calculator.Model{
			Input: calculator.Input{
				PensionIncome:    -1,
				WorkIncome:       -1,
				RentalNetIncome:  -1,
				AnnuityIncome:    -1,
				YearlyWithdrawal: -1,
			},
		},
	},
	{
		name: "Test Case 3",
		model: calculator.Model{
			Input: calculator.Input{
				PensionIncome:    math.Inf(1),
				WorkIncome:       math.Inf(1),
				RentalNetIncome:  math.Inf(1),
				AnnuityIncome:    math.Inf(1),
				YearlyWithdrawal: math.Inf(1),
			},
		},
	},
	{
		name: "Test Case 4",
		model: calculator.Model{
			Input: calculator.Input{
				PensionIncome:    math.Inf(-1),
				WorkIncome:       math.Inf(-1),
				RentalNetIncome:  math.Inf(-1),
				AnnuityIncome:    math.Inf(-1),
				YearlyWithdrawal: math.Inf(-1),
			},
		},
	},
}

func TestNewAdjustedGrossIncome(t *testing.T) {
	actual := calculator.NewAdjustedGrossIncome()
	expected := calculator.AdjustedGrossIncome{}

	assert.Equal(t, expected, actual)
}

func TestAdjustedGrossIncomeTraditionalCalculateTraditional(t *testing.T) {
	for _, test := range adjustedGrossIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AdjustedGrossIncome{}

			actual := c.CalculateTraditional(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAdjustedGrossIncomeTraditionalCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range adjustedGrossIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AdjustedGrossIncome{}

			actual := c.CalculateTraditional(&test.model)
			expected := test.model.Input.WorkIncome +
				test.model.Input.PensionIncome +
				test.model.Input.RentalNetIncome +
				test.model.Input.AnnuityIncome +
				test.model.Input.YearlyWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAdjustedGrossIncomeTraditionalCalculateRoth(t *testing.T) {
	for _, test := range adjustedGrossIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AdjustedGrossIncome{}

			actual := c.CalculateRoth(&test.model)
			expected := c.CalculateRothRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAdjustedGrossIncomeTraditionalCalculateRothRetirement(t *testing.T) {
	for _, test := range adjustedGrossIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AdjustedGrossIncome{}

			actual := c.CalculateRoth(&test.model)
			expected := test.model.Input.WorkIncome +
				test.model.Input.PensionIncome +
				test.model.Input.RentalNetIncome +
				test.model.Input.AnnuityIncome

			assert.Equal(t, expected, actual)
		})
	}
}
