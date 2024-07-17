package test

import (
	"math"
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
	{
		name:     "Test Case 2",
		model:    &snowballQATestModel2,
		expected: snowballQATestModel2Answer,
	},
}

func TestSnowballQACalculateSnowball(t *testing.T) {
	const float64EqualityThreshold = 1e-9
	for _, test := range snowballQATests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.SnowballAvalanche{
				MaxYear: 1000,
			}

			actual := c.CalculateSnowball(test.model)

			firstActual := actual[0]
			lastActual := actual[len(actual)-1]

			firstExpected := test.expected[0]
			lastExpected := test.expected[len(test.expected)-1]

			assert.Equal(t, firstActual.Months, firstExpected.Months)
			assert.Equal(t, lastActual.Months, lastExpected.Months)

			for idx := range firstActual.Balances {
				assert.True(t, math.Abs(firstActual.Balances[idx]-firstExpected.Balances[idx]) <= float64EqualityThreshold)
			}

			for idx := range lastActual.Balances {
				assert.True(t, math.Abs(lastActual.Balances[idx]-lastExpected.Balances[idx]) <= float64EqualityThreshold)
			}

			for idx := range firstActual.Payments {
				assert.True(t, math.Abs(firstActual.Payments[idx]-firstExpected.Payments[idx]) <= float64EqualityThreshold)
			}

			for idx := range lastActual.Payments {
				assert.True(t, math.Abs(lastActual.Payments[idx]-lastExpected.Payments[idx]) <= float64EqualityThreshold)
			}

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

var snowballQATestModel2 = calculator.Model{
	Input: calculator.Input{
		ExtraMonthlyPayment:     1000,
		OneTimeImmediatePayment: 500,
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

var snowballQATestModel2Answer = calculator.DebtSequences{
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
		Months:   []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18},
		Payments: []float64{350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 1652.6910388331376, 2450, 2450, 9.02361678711440},
		Balances: []float64{9650, 9424, 9194.986666666668, 8962.919822222224, 8727.758753185188, 8489.462203227657, 8247.988365937359, 8003.29487748319, 7755.3388091829665, 7504.076659972073, 7249.4643487717, 6991.457206755323, 6730.009969512061, 6465.076769105555, 4876.550873342716, 2458.9048849872856, 9.023616787116117, 0},
	},
}
