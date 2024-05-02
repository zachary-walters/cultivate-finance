package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAdjustedGrossIncomeRoth struct {
	mock.Mock
}

func (m *MockAdjustedGrossIncomeRoth) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAdjustedGrossIncomeRoth) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestAdjustedGrossIncomeRothCalculate(t *testing.T) {
	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					PensionIncome:   9492875,
					WorkIncome:      5293857,
					RentalNetIncome: 40394,
					AnnuityIncome:   5075,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AdjustedGrossIncomeRoth{}

			actual := c.Calculate(test.model)
			expected := test.model.Input.WorkIncome +
				test.model.Input.PensionIncome +
				test.model.Input.RentalNetIncome +
				test.model.Input.AnnuityIncome

			assert.Equal(t, expected, actual)
		})
	}
}
