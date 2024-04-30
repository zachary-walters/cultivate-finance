package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalDisbursementsRothMatchingNet struct {
	mock.Mock
}

func (m *MockTotalDisbursementsRothMatchingNet) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursementsRothMatchingNet) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestTotalDisbursementsRothMatchingNet(t *testing.T) {
	tests := []struct {
		name                    string
		model                   calculator.Model
		balancesRothMatchingNet calculator.ChartData
	}{
		{
			name: "Test case 0",
			balancesRothMatchingNet: calculator.ChartData{
				Withdrawal: map[int]float64{
					0: 100,
					1: 200,
					3: 300,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesRothMatchingNet := new(MockBalancesRothMatchingNet)
			mockBalancesRothMatchingNet.On("Calculate", test.model).Return(test.balancesRothMatchingNet)

			c := &calculator.TotalDisbursementsRothMatchingNet{
				BalancesRothMatchingNetContributionsCalculation: mockBalancesRothMatchingNet,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				rothBalancesMatchingNet := c.BalancesRothMatchingNetContributionsCalculation.Calculate(test.model)

				var totalDisbursementsAfterTax float64

				for _, income := range rothBalancesMatchingNet.Withdrawal {
					totalDisbursementsAfterTax += income
				}

				return totalDisbursementsAfterTax
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
