package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var debtPayoffMonthTests = []struct {
	name     string
	model    *calculator.Model
	snowball calculator.DebtSequences
}{
	{
		name:  "Test Case 0",
		model: &calculator.Model{},
		snowball: calculator.DebtSequences{
			{
				Months: []float64{1},
			},
		},
	},
	{
		name:  "Test Case 1",
		model: &calculator.Model{},
	},
	{
		name:  "Test Case 2",
		model: &calculator.Model{},
		snowball: calculator.DebtSequences{
			{
				Months: []float64{1},
			},
			{
				Months: []float64{2, 5, 6, 89},
			},
		},
	},
	{
		name:  "Test Case 2",
		model: &calculator.Model{},
		snowball: calculator.DebtSequences{
			{
				Months: []float64{0},
			},
			{
				Months: []float64{},
			},
		},
	},
}

func TestNewDebtPayoffMonth(t *testing.T) {
	actual := calculator.NewDebtPayoffMonth()
	expected := &calculator.DebtPayoffMonth{
		SnowballCalculation: calculator.NewSnowball(),
	}

	assert.Equal(t, expected, actual)
}

func TestDebtPayoffMonthCalculateSnowball(t *testing.T) {
	for _, test := range debtPayoffMonthTests {
		t.Run(test.name, func(t *testing.T) {
			mockSnowballCalculation := new(MockSnowballCalculation)
			mockSnowballCalculation.On("CalculateSnowball", test.model).Return(test.snowball)

			c := &calculator.DebtPayoffMonth{
				SnowballCalculation: mockSnowballCalculation,
			}

			actual := c.CalculateSnowball(test.model)
			expected := func() float64 {
				if len(test.snowball) <= 0 {
					return 0.0
				}

				lastDebtSequence := test.snowball[len(test.snowball)-1]

				if len(lastDebtSequence.Months) <= 0 {
					return 0
				}

				lastMonth := lastDebtSequence.Months[len(lastDebtSequence.Months)-1]

				return lastMonth
			}()

			assert.Equal(t, actual, expected)
		})
	}
}
