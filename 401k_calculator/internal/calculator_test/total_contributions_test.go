package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalContributions struct {
	mock.Mock
}

func (m *MockTotalContributions) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributions) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalContributionsCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                string
		balancesTraditional calculator.ChartData
	}{
		{
			name: "Test Case 0",
			balancesTraditional: calculator.ChartData{
				Contribution: map[int32]float64{
					0: 100,
					1: 200,
				},
			},
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesTraditional := new(MockBalancesTraditional)
			mockBalancesTraditional.On("Calculate", model).Return(test.balancesTraditional)

			c := &calculator.TotalContributions{
				BalancesTraditionalCalculation: mockBalancesTraditional,
			}

			actual := c.CalculateTraditional(model)
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
