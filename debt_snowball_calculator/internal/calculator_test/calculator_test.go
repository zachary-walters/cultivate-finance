package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

func TestSanitizeToZero(t *testing.T) {
	c := calculator.AbstractCalculation{}

	actual := c.SanitizeToZero(1)
	expected := 1.0
	assert.Equal(t, expected, actual)

	actual = c.SanitizeToZero(1.0)
	expected = 1.0
	assert.Equal(t, expected, actual)

	actual = c.SanitizeToZero(-1)
	expected = 0.0
	assert.Equal(t, expected, actual)

	actual = c.SanitizeToZero(-1.0)
	expected = 0.0
	assert.Equal(t, expected, actual)

	actual = c.SanitizeToZero("not a valid input")
	expected = 0.0
	assert.Equal(t, expected, actual)
}

func TestNewModel(t *testing.T) {
	i := calculator.Input{}

	actual := calculator.NewModel(i)
	expected := &calculator.Model{
		Input: i,
	}

	assert.Equal(t, expected, actual)

	i = calculator.Input{
		Debts: []calculator.Debt{
			{
				Amount:         1337.0,
				AnnualInterest: 9940.0,
			},
		},
		Datakey: "fakeDK",
	}

	actual = calculator.NewModel(i)
	expected = &calculator.Model{
		Input: i,
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, expected.Input.Datakey, actual.Input.Datakey)
	assert.Equal(t, expected.Input.Debts[0].Amount, actual.Input.Debts[0].Amount)
	assert.Equal(t, expected.Input.Debts[0].AnnualInterest, actual.Input.Debts[0].AnnualInterest)
	assert.Equal(t, 9940.0, actual.Input.Debts[0].AnnualInterest)
	assert.Equal(t, 1337.0, actual.Input.Debts[0].Amount)
}
