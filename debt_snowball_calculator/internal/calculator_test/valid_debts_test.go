package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var validDebtsTests = []struct {
	name      string
	model     calculator.Model
	snowball  calculator.DebtSequences
	avalanche calculator.DebtSequences
}{
	{
		name: "Test Case 0 - snowball",
		model: calculator.Model{
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
			},
		},
	},
	{
		name: "Test Case 1 - snowball",
		model: calculator.Model{
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
	{
		name: "Test Case 2 - avalanche",
		model: calculator.Model{
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
		avalanche: calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name: "d1",
				},
			},
		},
	},
	{
		name: "Test Case 3 - avalanche",
		model: calculator.Model{
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
		avalanche: calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name: "d1",
				},
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

func TestNewValidDebts(t *testing.T) {
	actual := calculator.NewValidDebts()
	expected := &calculator.ValidDebts{
		SnowballAvalancheCalculation: calculator.NewSnowballAvalanche(),
	}

	assert.Equal(t, expected, actual)
}

func TestValidDebtsCalculate(t *testing.T) {
	for _, test := range validDebtsTests {
		t.Run(test.name, func(t *testing.T) {
			mockSnowball := new(MockSnowballCalculation)
			mockSnowball.On("CalculateSnowball", test.model).Return(test.snowball)
			mockSnowball.On("CalculateAvalanche", test.model).Return(test.avalanche)

			c := &calculator.ValidDebts{
				SnowballAvalancheCalculation: mockSnowball,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() []calculator.Debt {
				snowballDebts := []calculator.Debt{}

				validDebts := []calculator.Debt{}

				for _, debtSequence := range test.snowball {
					snowballDebts = append(snowballDebts, debtSequence.Debt)
				}

				for _, avalancheSequence := range test.avalanche {
					for _, debt := range snowballDebts {
						if avalancheSequence.Debt == debt {
							validDebts = append(validDebts, debt)
						}
					}
				}

				return validDebts
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestValidDebtsCalculateAvalanche(t *testing.T) {
	for _, test := range validDebtsTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.NewValidDebts()

			actual := c.CalculateAvalanche(test.model)
			expected := c.CalculateSnowball(test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
