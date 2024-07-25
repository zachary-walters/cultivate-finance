package test

import (
	"sync"
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
	expected := calculator.Model{
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
	expected = calculator.Model{
		Input: i,
	}

	assert.Equal(t, expected, actual)
	assert.Equal(t, expected.Input.Datakey, actual.Input.Datakey)
	assert.Equal(t, expected.Input.Debts[0].Amount, actual.Input.Debts[0].Amount)
	assert.Equal(t, expected.Input.Debts[0].AnnualInterest, actual.Input.Debts[0].AnnualInterest)
	assert.Equal(t, 9940.0, actual.Input.Debts[0].AnnualInterest)
	assert.Equal(t, 1337.0, actual.Input.Debts[0].Amount)
}

func TestCalculateSynchronous(t *testing.T) {
	model := calculator.Model{}
	datakey := "testKey"

	t.Run("Test with Calculation type", func(t *testing.T) {
		calculation := &MockCalculation{}
		calculation.On("CalculateSnowball", model).Return(1.0)
		calculation.On("CalculateAvalanche", model).Return(2.0)

		expected := calculator.CalculationData{
			Datakey:   datakey,
			Value:     1.0,
			Avalanche: 2.0,
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with SequenceCalculation type", func(t *testing.T) {
		calculation := &MockSequenceCalculation{}
		calculation.On("CalculateSnowball", model).Return([]float64{1.0, 100, 3.9})
		calculation.On("CalculateAvalanche", model).Return([]float64{2.0, 11111, 33333})

		expected := calculator.CalculationData{
			Datakey:   datakey,
			Value:     []float64{1.0, 100, 3.9},
			Avalanche: []float64{2.0, 11111, 33333},
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with SnowballCalculation type", func(t *testing.T) {
		calculation := &MockSnowballCalculation{}
		calculation.On("CalculateSnowball", model).Return(calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name:   "12345",
					Amount: 1337,
				},
			},
			{
				Debt: calculator.Debt{
					Name: "hello",
				},
			},
		})
		calculation.On("CalculateAvalanche", model).Return(calculator.DebtSequences{
			{
				Debt: calculator.Debt{
					Name:   "p923094",
					Amount: 210983120938,
				},
			},
			{
				Debt: calculator.Debt{
					Name: "uarstyn",
				},
			},
		})

		expected := calculator.CalculationData{
			Datakey: datakey,
			Value: calculator.DebtSequences{
				{
					Debt: calculator.Debt{
						Name:   "12345",
						Amount: 1337,
					},
				},
				{
					Debt: calculator.Debt{
						Name: "hello",
					},
				},
			},
			Avalanche: calculator.DebtSequences{
				{
					Debt: calculator.Debt{
						Name:   "p923094",
						Amount: 210983120938,
					},
				},
				{
					Debt: calculator.Debt{
						Name: "uarstyn",
					},
				},
			},
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})
}

func TestCalculateAsync(t *testing.T) {
	model := calculator.Model{}
	datakey := "testKey"

	t.Run("Test with Calculation type", func(t *testing.T) {
		calculation := &MockCalculation{}
		calculation.On("CalculateSnowball", model).Return(1.0)
		calculation.On("CalculateAvalanche", model).Return(2.0)

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsync(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})
}
