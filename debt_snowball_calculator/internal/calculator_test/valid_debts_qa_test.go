package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var validDebtsQATests = []struct {
	name      string
	model     *calculator.Model
	snowball  calculator.DebtSequences
	avalanche calculator.DebtSequences
}{
	{
		model: &calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Name: "d1",
					},
					{
						Name: "d2",
					},
					{
						Name: "d3",
					},
				},
			},
		},
		snowball: calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name: "d1",
				},
				Invalid: true,
			},
			{
				Debt: calculator.Debt{
					Name: "d2",
				},
			},
			{
				Debt: calculator.Debt{
					Name: "d3",
				},
			},
		},
	},
}

var validDebtQAAnswer0 = []calculator.Debt{
	{
		Name: "d2",
	},
	{
		Name: "d3",
	},
}

func TestValidDebtsCalculateQA(t *testing.T) {
	for _, test := range validDebtsQATests {
		t.Run(test.name, func(t *testing.T) {
			mockSnowball := new(MockSnowballCalculation)
			mockSnowball.On("CalculateSnowball", test.model).Return(test.snowball)
			mockSnowball.On("CalculateAvalanche", test.model).Return(test.avalanche)

			c := calculator.ValidDebts{
				SnowballAvalancheCalculation: mockSnowball,
			}

			actual := c.CalculateSnowball(test.model)
			expected := validDebtQAAnswer0

			assert.Equal(t, expected, actual)
		})
	}
}
