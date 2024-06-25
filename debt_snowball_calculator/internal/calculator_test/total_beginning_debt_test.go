package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var totalBeginningDebtTests = []struct {
	name  string
	model *calculator.Model
}{
	{
		name: "Test Case 0",
		model: &calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Amount: 1,
					},
					{
						Amount: 2,
					},
					{
						Amount: 3,
					},
					{
						Amount: 4,
					},
				},
			},
		},
	},
	{
		name: "Test Case 1",
		model: &calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						Amount: -1,
					},
					{
						Amount: -2,
					},
					{
						Amount: -3,
					},
					{
						Amount: -4,
					},
				},
			},
		},
	},
	{
		name:  "Test Case 2",
		model: &calculator.Model{},
	},
}

func TestNewTotalBeginningDebt(t *testing.T) {
	actual := calculator.NewTotalBeginningDebt()
	expected := &calculator.TotalBeginningDebt{}

	assert.Equal(t, expected, actual)
}

func TestTotalBeginningDebtCalculate(t *testing.T) {
	for _, test := range totalBeginningDebtTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalBeginningDebt{}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				total := 0.0
				for _, debt := range test.model.Input.Debts {
					total += debt.Amount
				}

				return c.SanitizeToZero(total)
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
