package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalDisbursements struct {
	mock.Mock
}

func (m *MockTotalDisbursements) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalDisbursementsTests = []struct {
	name                    string
	model                   calculator.Model
	balancesTraditional     calculator.ChartData
	balancerRothMatchingNet calculator.ChartData
}{
	{
		name: "Test case 0",
		balancesTraditional: calculator.ChartData{
			AfterTaxIncome: map[int32]float64{
				0: 100,
				1: 200,
				3: 300,
			},
		},
		balancerRothMatchingNet: calculator.ChartData{
			AfterTaxIncome: map[int32]float64{
				0: 100,
				1: 200,
				3: 300,
			},
		},
	},
	{
		name: "Test case 1",
	},
}

func NewTotalDisbursements(t *testing.T) {
	actual := calculator.NewTotalDisbursements()
	expected := calculator.TotalDisbursements{
		BalancesTraditionalCalculation:                  calculator.NewBalancesTraditional(),
		BalancesRothMatchingNetContributionsCalculation: calculator.NewBalancesRothMatchingNetContributions(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalDisburesmentsCalculateTraditional(t *testing.T) {
	for _, test := range totalDisbursementsTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalDisbursements{}

			actual := c.CalculateTraditional(&test.model)
			expected := float64(0)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalDisburesmentsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalDisbursementsTests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesTraditional := new(MockBalancesTraditional)
			mockBalancesTraditional.On("Calculate", &test.model).Return(test.balancesTraditional)

			c := &calculator.TotalDisbursements{
				BalancesTraditionalCalculation: mockBalancesTraditional,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				var totalDisbursementsTraditional float64

				for _, income := range test.balancesTraditional.AfterTaxIncome {
					totalDisbursementsTraditional += income
				}

				return totalDisbursementsTraditional
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalDisburesmentsCalculateRoth(t *testing.T) {
	for _, test := range totalDisbursementsTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalDisbursements{}

			actual := c.CalculateRoth(&test.model)
			expected := float64(0)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalDisburesmentsCalculateRothRetirement(t *testing.T) {
	for _, test := range totalDisbursementsTests {
		t.Run(test.name, func(t *testing.T) {
			mockRothMatchingNetContributions := new(MockBalancesRothMatchingNet)
			mockRothMatchingNetContributions.On("Calculate", &test.model).Return(test.balancesTraditional)

			c := &calculator.TotalDisbursements{
				BalancesRothMatchingNetContributionsCalculation: mockRothMatchingNetContributions,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				var totalDisbursementsTraditional float64

				for _, income := range test.balancesTraditional.AfterTaxIncome {
					totalDisbursementsTraditional += income
				}

				return totalDisbursementsTraditional
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
