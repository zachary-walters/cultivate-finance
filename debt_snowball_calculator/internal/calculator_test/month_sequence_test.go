package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var monthSequenceTests = []struct {
	name            string
	model           *calculator.Model
	debtPayoffMonth float64
}{
	{
		name:            "Test Case 0",
		model:           &calculator.Model{},
		debtPayoffMonth: 10,
	},
	{
		name:            "Test Case 1",
		model:           &calculator.Model{},
		debtPayoffMonth: 99999,
	},
	{
		name:            "Test Case 2",
		model:           &calculator.Model{},
		debtPayoffMonth: 0,
	},
}

func TestNewMonthSequence(t *testing.T) {
	actual := calculator.NewMonthSequence()
	expected := &calculator.MonthSequence{
		DebtPayoffMonthCalculation: calculator.NewDebtPayoffMonth(),
	}

	assert.Equal(t, expected, actual)
}

func TestNewMonthlySequenceCalculate(t *testing.T) {
	for _, test := range monthSequenceTests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoffMonth := new(MockCalculation)
			mockDebtPayoffMonth.On("Calculate", test.model).Return(test.debtPayoffMonth)

			c := &calculator.MonthSequence{
				DebtPayoffMonthCalculation: mockDebtPayoffMonth,
			}

			actual := c.Calculate(test.model)
			expected := func() []float64 {
				sequence := []float64{}

				for i := 1; i <= int(test.debtPayoffMonth)+2; i++ {
					sequence = append(sequence, float64(i))
				}

				return sequence
			}()

			assert.Equal(t, expected, actual)
		})

	}
}
