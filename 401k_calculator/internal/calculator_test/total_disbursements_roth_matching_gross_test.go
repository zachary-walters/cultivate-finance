package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalDisbursementsRothMatchingGross struct {
	mock.Mock
}

func (m *MockTotalDisbursementsRothMatchingGross) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalDisbursementsRothMatchingGross(t *testing.T) {
	tests := []struct {
		name                      string
		model                     calculator.Model
		balancesRothMatchingGross calculator.ChartData
	}{
		{
			name: "Test case 0",
			balancesRothMatchingGross: calculator.ChartData{
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
			mockBalancesRothMatchingGross := new(MockBalancesRothMatchingGross)
			mockBalancesRothMatchingGross.On("Calculate", test.model).Return(test.balancesRothMatchingGross)

			c := &calculator.TotalDisbursementsRothMatchingGross{
				BalancesRothMatchingGrossContributionsCalculation: mockBalancesRothMatchingGross,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				rothBalancesMatchingGross := c.BalancesRothMatchingGrossContributionsCalculation.Calculate(test.model)

				var totalDisbursementsAfterTax float64

				for _, income := range rothBalancesMatchingGross.Withdrawal {
					totalDisbursementsAfterTax += income
				}

				return totalDisbursementsAfterTax
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
