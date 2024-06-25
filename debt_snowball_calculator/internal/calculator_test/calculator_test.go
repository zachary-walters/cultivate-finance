package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

type MockCalculation struct {
	mock.Mock
}

func (m *MockCalculation) Calculate(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

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

var basicTestModel = calculator.Model{
	Input: calculator.Input{
		ExtraMonthlyPayment:     1000,
		OneTimeImmediatePayment: 1000,
		Debts: []calculator.Debt{
			{
				Name:           "debt0",

				Amount:         1000,
				MinimumPayment: 50,
				AnnualInterest: 19.49,
			},
			{
				Name:           "debt1",

				Amount:         2000,
				MinimumPayment: 100,
				AnnualInterest: 24.49,
			},
			{
				Name:           "debt2",

				Amount:         3000,
				MinimumPayment: 150,
				AnnualInterest: 26.99,
			},
			{
				Name:           "debt3",

				Amount:         4000,
				MinimumPayment: 200,
				AnnualInterest: 24.29,
			},
			{
				Name:           "debt4",

				Amount:         5000,
				MinimumPayment: 200,
				AnnualInterest: 29.99,
			},
			{
				Name:           "debt5",

				Amount:         6000,
				MinimumPayment: 200,
				AnnualInterest: 15.00,
			},
			{
				Name:           "debt6",

				Amount:         7000,
				MinimumPayment: 200,
				AnnualInterest: 18.00,
			},
			{
				Name:           "debt7",

				Amount:         10000,
				MinimumPayment: 350,
				AnnualInterest: 16.00,
			},
		},
	},
}
