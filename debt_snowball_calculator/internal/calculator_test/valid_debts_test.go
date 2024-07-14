package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var validDebtsTests = []struct {
	name      string
	model     *calculator.Model
	snowball  calculator.DebtSequences
	avalanche calculator.DebtSequences
}{
	{
		name: "Test Case 0 - snowball",
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
		},
	},
	{
		name: "Test Case 1 - snowball",
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
		avalanche: calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name: "d1",
				},
				Invalid: true,
			},
		},
	},
	{
		name: "Test Case 3 - avalanche",
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
		SnowballCalculation: calculator.NewSnowball(),
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
				SnowballCalculation: mockSnowball,
			}

			actual := c.Calculate(test.model)
			expected := func() []calculator.Debt {
				validDebts := []calculator.Debt{}
				invalidDebts := []calculator.Debt{}

				for _, debtSequence := range test.snowball {
					if debtSequence.Invalid {
						invalidDebts = append(invalidDebts, debtSequence.Debt)
					}
				}

				for _, debtSequence := range test.avalanche {
					if debtSequence.Invalid {
						invalidDebts = append(invalidDebts, debtSequence.Debt)
					}
				}

				for _, debt := range test.model.Input.Debts {
					if func(s []calculator.Debt, d calculator.Debt) bool {
						for _, a := range s {
							if a == d {
								return true
							}
						}
						return false
					}(invalidDebts, debt) {
						continue
					}

					validDebts = append(validDebts, debt)
				}

				return validDebts
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
