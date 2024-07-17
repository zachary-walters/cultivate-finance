package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var validSnowballAvalancheTests = []struct {
	name       string
	model      *calculator.Model
	snowball   calculator.DebtSequences
	validDebts []calculator.Debt
}{
	{
		name:  "Test Case 0",
		model: &calculator.Model{},
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
		validDebts: []calculator.Debt{
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
	{
		name:  "Test Case 0",
		model: &calculator.Model{},
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
		validDebts: []calculator.Debt{
			{
				Name: "d1",
			},
		},
	},
}

func TestNewValidSnowballAvalanche(t *testing.T) {
	actual := calculator.NewValidSnowballAvalanche()
	expected := &calculator.ValidSnowballAvalanche{
		SnowballAvalancheCalculation: calculator.NewSnowballAvalanche(),
		ValidDebtsCalculation:        calculator.NewValidDebts(),
	}

	assert.Equal(t, expected, actual)
}

func TestValidSnowballAvalancheCalculateSnowball(t *testing.T) {
	for _, test := range validSnowballAvalancheTests {
		t.Run(test.name, func(t *testing.T) {
			mockSnowball := new(MockSnowballCalculation)
			mockSnowball.On("CalculateSnowball", test.model).Return(test.snowball)

			mockValidDebts := new(MockValidDebtsCalculation)
			mockValidDebts.On("CalculateSnowball", test.model).Return(test.validDebts)

			c := &calculator.ValidSnowballAvalanche{
				SnowballAvalancheCalculation: mockSnowball,
				ValidDebtsCalculation:        mockValidDebts,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() calculator.DebtSequences {
				debtSequences := calculator.DebtSequences{}

				for _, debtSequence := range test.snowball {
					for _, debt := range test.validDebts {
						if debt == debtSequence.Debt {
							debtSequences = append(debtSequences, debtSequence)
						}
					}
				}

				return debtSequences
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestValidSnowballAvalancheCalculateAvalanche(t *testing.T) {
	for _, test := range validSnowballAvalancheTests {
		t.Run(test.name, func(t *testing.T) {
			mockSnowball := new(MockSnowballCalculation)
			mockSnowball.On("CalculateAvalanche", test.model).Return(test.snowball)

			mockValidDebts := new(MockValidDebtsCalculation)
			mockValidDebts.On("CalculateAvalanche", test.model).Return(test.validDebts)

			c := &calculator.ValidSnowballAvalanche{
				SnowballAvalancheCalculation: mockSnowball,
				ValidDebtsCalculation:        mockValidDebts,
			}

			actual := c.CalculateAvalanche(test.model)
			expected := func() calculator.DebtSequences {
				debtSequences := calculator.DebtSequences{}

				for _, debtSequence := range test.snowball {
					for _, debt := range test.validDebts {
						if debt == debtSequence.Debt {
							debtSequences = append(debtSequences, debtSequence)
						}
					}
				}

				return debtSequences
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
