package test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalContributions struct {
	mock.Mock
}

func (m *MockTotalContributions) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalContributionsTests = []struct {
	name                                 string
	balancesTraditional                  calculator.ChartData
	balancesRothMatchingNetContributions calculator.ChartData
}{
	{
		name: "Test Case 0",
		balancesTraditional: calculator.ChartData{
			Contribution: map[int32]float64{
				0: 100,
				1: 200,
			},
		},
		balancesRothMatchingNetContributions: calculator.ChartData{
			Contribution: map[int32]float64{
				0: 100,
				1: 200,
			},
		},
	},
	{
		name: "Test Case 1",
	},
	{
		name: "Test Case 2",
		balancesTraditional: calculator.ChartData{
			Contribution: func() map[int32]float64 {
				m := map[int32]float64{}
				for i := 0; i <= 10000; i++ {
					m[int32(i)] = math.Round(rand.Float64())
				}
				return m
			}(),
		},
		balancesRothMatchingNetContributions: calculator.ChartData{
			Contribution: func() map[int32]float64 {
				m := map[int32]float64{}
				for i := 0; i <= 10000; i++ {
					m[int32(i)] = math.Round(rand.Float64())
				}
				return m
			}(),
		},
	},
}

func TestNewTotalContributions(t *testing.T) {
	actual := calculator.NewTotalContributions()
	expected := calculator.TotalContributions{
		BalancesTraditionalCalculation:                  calculator.NewBalancesTraditional(),
		BalancesRothMatchingNetContributionsCalculation: calculator.NewBalancesRothMatchingNetContributions(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalContributionsCalculateTraditional(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesTraditional := new(MockBalancesTraditional)
			mockBalancesTraditional.On("Calculate", &model).Return(test.balancesTraditional)

			c := &calculator.TotalContributions{
				BalancesTraditionalCalculation: mockBalancesTraditional,
			}

			actual := c.CalculateTraditional(&model)
			expected := func() float64 {
				total := 0.0
				for _, contribution := range test.balancesTraditional.Contribution {
					total += contribution
				}

				return total
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalContributionsCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range totalContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalContributions{}

			actual := c.CalculateTraditionalRetirement(&calculator.Model{})

			assert.Zero(t, actual)
		})
	}
}

func TestTotalContributionsCalculateRoth(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesRothMatchingNet := new(MockBalancesRothMatchingNet)
			mockBalancesRothMatchingNet.On("Calculate", &model).Return(test.balancesRothMatchingNetContributions)

			c := &calculator.TotalContributions{
				BalancesRothMatchingNetContributionsCalculation: mockBalancesRothMatchingNet,
			}

			actual := c.CalculateRoth(&model)
			expected := func() float64 {
				total := 0.0
				for _, contribution := range test.balancesRothMatchingNetContributions.Contribution {
					total += contribution
				}

				return total
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalContributionsCalculateRothRetirement(t *testing.T) {
	for _, test := range totalContributionsTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalContributions{}

			actual := c.CalculateRothRetirement(&calculator.Model{})

			assert.Zero(t, actual)
		})
	}
}
