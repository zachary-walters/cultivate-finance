package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var totalPaymentsTests = []struct {
	name                    string
	model                   calculator.Model
	monthlySequencePayments []float64
}{
	{
		name:                    "Test Case 0",
		model:                   calculator.Model{},
		monthlySequencePayments: []float64{1, 2, 3, 4},
	},
	{
		name:                    "Test Case 0",
		model:                   calculator.Model{},
		monthlySequencePayments: []float64{math.Inf(-1)},
	},
}

func TestNewTotalPayments(t *testing.T) {
	actual := calculator.NewTotalPayments()
	expected := &calculator.TotalPayments{
		MonthlySequencePaymentsCalculation: calculator.NewMonthlySequencePayments(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalPaymentsCalculateSnowball(t *testing.T) {
	for _, test := range totalPaymentsTests {
		t.Run(test.name, func(t *testing.T) {
			mockMonthlySequencePayments := new(MockSequenceCalculation)
			mockMonthlySequencePayments.On("CalculateSnowball", test.model).Return(test.monthlySequencePayments)

			c := &calculator.TotalPayments{
				MonthlySequencePaymentsCalculation: mockMonthlySequencePayments,
			}

			actual := c.CalculateSnowball(test.model)

			expected := 0.0
			for _, p := range test.monthlySequencePayments {
				expected += p
			}

			expected = c.SanitizeToZero(expected)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalPaymentsCalculateAvalanche(t *testing.T) {
	for _, test := range totalPaymentsTests {
		t.Run(test.name, func(t *testing.T) {
			mockMonthlySequencePayments := new(MockSequenceCalculation)
			mockMonthlySequencePayments.On("CalculateAvalanche", test.model).Return(test.monthlySequencePayments)

			c := &calculator.TotalPayments{
				MonthlySequencePaymentsCalculation: mockMonthlySequencePayments,
			}

			actual := c.CalculateAvalanche(test.model)

			expected := 0.0
			for _, p := range test.monthlySequencePayments {
				expected += p
			}

			expected = c.SanitizeToZero(expected)

			assert.Equal(t, expected, actual)
		})
	}
}
