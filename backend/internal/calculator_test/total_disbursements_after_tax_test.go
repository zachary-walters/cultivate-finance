package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/backend/internal/calculator"
)

type MockTotalDisbursementsAfterTax struct {
	mock.Mock
}

func (m *MockTotalDisbursementsAfterTax) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalDisburesmentsAfterTaxCalculate(t *testing.T) {
	tests := []struct {
		name                string
		model               calculator.Model
		balancesTraditional calculator.ChartData
	}{
		{
			name: "Test case 0",
			balancesTraditional: calculator.ChartData{
				AfterTaxIncome: map[int]float64{
					0: 100,
					1: 200,
					3: 300,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockBalancesTraditional := new(MockBalancesTraditional)
			mockBalancesTraditional.On("Calculate", test.model).Return(test.balancesTraditional)

			c := &calculator.TotalDisbursementsAfterTax{
				BalancesTraditionalCalculation: mockBalancesTraditional,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				traditionalBalances := c.BalancesTraditionalCalculation.Calculate(test.model)

				var totalDisbursementsAfterTax float64

				for _, income := range traditionalBalances.AfterTaxIncome {
					totalDisbursementsAfterTax += income
				}

				return totalDisbursementsAfterTax
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
