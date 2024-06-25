package test

import (
	"testing"

	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var snowballChartTests = []struct {
	name string
	model *calculator.Model
}{
	{
		name: "Test Case 0",
		model: &basicTestModel1,
	},
}

func TestSnowballCalculate(t *testing.T) {
	for _, test := range snowballChartTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.Snowball{}

			c.Calculate(test.model)
		})
	}
}

var basicTestModel1 = calculator.Model{
	Input: calculator.Input{
		ExtraMonthlyPayment:     100,
		OneTimeImmediatePayment: 400,
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