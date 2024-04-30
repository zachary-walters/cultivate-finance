package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalContributionsRoth struct {
	mock.Mock
}

func (m *MockTotalContributionsRoth) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalContributionsRoth) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestTotalContributionsRothCalculate(t *testing.T) {
	tests := []struct {
		name                                 string
		balancesRothMatchingNetContributions calculator.ChartData
	}{
		{
			name: "Test Case 0",
			balancesRothMatchingNetContributions: calculator.ChartData{
				Contribution: map[int]float64{
					0: 100,
					1: 200,
				},
			},
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockRothMatchingNet := new(MockBalancesRothMatchingNet)
			mockRothMatchingNet.On("Calculate", model).Return(test.balancesRothMatchingNetContributions)

			c := &calculator.TotalContributionsRoth{
				BalancesRothMatchingNetContributionsCalculation: mockRothMatchingNet,
			}

			actual := c.Calculate(model)
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
