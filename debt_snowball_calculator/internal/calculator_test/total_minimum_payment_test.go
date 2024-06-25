package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var testTotalMinimumPayment = []struct {
	name  string
	model *calculator.Model
}{
	{
		name: "Test Case 0",
		model: &calculator.Model{
			Input: calculator.Input{
				Debts: []calculator.Debt{
					{
						MinimumPayment: 1,
					},
					{
						MinimumPayment: 2,
					},
					{
						MinimumPayment: 3,
					},
					{
						MinimumPayment: 4,
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
						MinimumPayment: -1,
					},
					{
						MinimumPayment: -2,
					},
					{
						MinimumPayment: -3,
					},
					{
						MinimumPayment: -4,
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

func TestNewMinimumPayment(t *testing.T) {
	actual := calculator.NewTotalMinimumPayment()
	expected := &calculator.TotalMinimumPayment{}

	assert.Equal(t, expected, actual)
}

func TestTotalMinimumPaymentCalculate(t *testing.T) {
	for _, test := range testTotalMinimumPayment {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalMinimumPayment{}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				total := 0.0
				for _, debt := range test.model.Input.Debts {
					total += debt.MinimumPayment
				}

				return c.SanitizeToZero(total)
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
