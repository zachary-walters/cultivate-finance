package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var snowballQATests = []struct {
	name     string
	model    *calculator.Model
	expected calculator.DebtSequences
}{
	{
		name:     "Test Case 0",
		model:    &snowballQATestModel0,
		expected: snowballQATestModel0Answer,
	},
	{
		name:     "Test Case 1",
		model:    &snowballQATestModel1,
		expected: snowballQATestModel1Answer,
	},
}

func TestSnowballQACalculate(t *testing.T) {
	for _, test := range snowballQATests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.Snowball{}

			actual := c.Calculate(test.model)

			assert.Equal(t, actual[0], test.expected[0])
			assert.Equal(t, actual[len(actual)-1], test.expected[len(test.expected)-1])
		})
	}
}

var snowballQATestModel0 = calculator.Model{
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

var snowballQATestModel0Answer = calculator.DebtSequences{
	{
		Debt: calculator.Debt{
			Name:           "debt0",
			Amount:         1000,
			AnnualInterest: 19.49,
			MinimumPayment: 50,
		},
		Months:   []float64{1, 2, 3, 4, 5},
		Payments: []float64{550, 150, 150, 150, 7.507879127317551},
		Balances: []float64{450, 304.8725, 157.38788752083335, 7.507879127317547, 0},
	},
	{
		Debt: calculator.Debt{
			Name:           "debt7",
			Amount:         10000,
			AnnualInterest: 16,
			MinimumPayment: 350,
		},
		Months:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		Payments: []float64{350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 986.3932294567666, 1394.7511709901887},
		Balances: []float64{9650, 9424, 9194.986666666666, 8962.919822222222, 8727.758753185186, 8489.462203227655, 8247.988365937357, 8003.294877483188, 7755.338809182965, 7504.076659972071, 7249.4643487716985, 6991.457206755321, 6730.009969512059, 6465.076769105553, 6196.61112602696, 5924.565941040652, 5648.893486921194, 5369.545400080144, 5086.472672081212, 4799.625641042295, 4508.953982922859, 4214.406702695163, 3915.9321253977655, 3613.4778870697355, 3306.9909255639986, 2996.4174712381855, 2681.7030375213612, 2362.7924113549793, 1394.7511709901887, 0},
	},
}

var snowballQATestModel1 = calculator.Model{
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

var snowballQATestModel1Answer = calculator.DebtSequences{
	{
		Debt: calculator.Debt{
			Name:           "debt0",
			Amount:         1000,
			AnnualInterest: 19.49,
			MinimumPayment: 50,
		},
		Months:   []float64{1},
		Payments: []float64{1000},
		Balances: []float64{0},
	},
	{
		Debt: calculator.Debt{
			Name:           "debt7",
			Amount:         10000,
			AnnualInterest: 16,
			MinimumPayment: 350,
		},
		Months:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17},
		Payments: []float64{350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 2289.978281398712, 2450, 1804.510020443504},
		Balances: []float64{9650, 9424, 9194.986666666666, 8962.919822222222, 8727.758753185186, 8489.462203227655, 8247.988365937357, 8003.294877483188, 7755.338809182965, 7504.076659972071, 7249.4643487716985, 6991.457206755321, 6730.009969512059, 6465.076769105553, 4230.766467542931, 1804.510020443504, 0},
	},
}
