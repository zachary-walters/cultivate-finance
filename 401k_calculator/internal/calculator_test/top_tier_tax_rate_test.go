package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTopTierTaxRate struct {
	mock.Mock
}

func (m *MockTopTierTaxRate) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTopTierTaxRate) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTopTierTaxRate) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTopTierTaxRate) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var topTierTaxRateTests = []struct {
	name                          string
	model                         calculator.Model
	taxOnTraditionalIRAWithdrawal float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			Input: calculator.Input{
				YearlyWithdrawal: 10000,
			},
		},
		taxOnTraditionalIRAWithdrawal: 3546455,
	},
	{
		name: "Test Case 1",
		model: calculator.Model{
			Input: calculator.Input{
				YearlyWithdrawal: 10000,
			},
		},
	},
	{
		name:                          "Test Case 2",
		taxOnTraditionalIRAWithdrawal: 3546455,
	},
	{
		name: "Test Case 3",
	},
}

func TestNewTopTierTaxRate(t *testing.T) {
	actual := calculator.NewTopTierTaxRate()
	expected := calculator.TopTierTaxRate{
		TaxOnTraditionalIRAWithdrawalCalculation: calculator.NewTaxOnTraditionalIRAWithdrawal(),
	}

	assert.Equal(t, expected, actual)
}

func TestTopTierTaxRateCalculateTraditional(t *testing.T) {
	for _, test := range topTierTaxRateTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TopTierTaxRate{}

			assert.Zero(t, c.CalculateTraditional(&test.model))
		})
	}
}

func TestTopTierTaxRateCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range topTierTaxRateTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateTraditional", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := &calculator.TopTierTaxRate{
				TaxOnTraditionalIRAWithdrawalCalculation: mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				if test.model.Input.YearlyWithdrawal == 0 {
					return 0.0
				}
				return test.taxOnTraditionalIRAWithdrawal / test.model.Input.YearlyWithdrawal
			}()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestTopTierTaxRateCalculateRoth(t *testing.T) {
	for _, test := range topTierTaxRateTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TopTierTaxRate{}

			assert.Zero(t, c.CalculateRoth(&test.model))
		})
	}
}

func TestTopTierTaxRateCalculateRothRetirement(t *testing.T) {
	for _, test := range topTierTaxRateTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TopTierTaxRate{}

			assert.Zero(t, c.CalculateRothRetirement(&test.model))
		})
	}
}
