package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalAnnualRetirementIncomeBeforeTax struct {
	mock.Mock
}

func (m *MockTotalAnnualRetirementIncomeBeforeTax) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTax) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTax) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTax) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalAnnualRetirementIncomeBeforeTaxCalculateTraditional(t *testing.T) {
	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					AnnuityIncome:             123452,
					OtherLongTermCapitalGains: 41325,
					OtherTaxableIncome:        234,
					PensionIncome:             21341,
					QualifiedDividends:        64536,
					RentalNetIncome:           2345,
					SocialSecurity:            65347,
					WorkIncome:                643653,
					YearlyWithdrawal:          12341234,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.TotalAnnualRetirementIncomeBeforeTax{}

			actual := c.CalculateTraditional(&test.model)
			expected := test.model.Input.AnnuityIncome +
				test.model.Input.OtherLongTermCapitalGains +
				test.model.Input.OtherTaxableIncome +
				test.model.Input.PensionIncome +
				test.model.Input.QualifiedDividends +
				test.model.Input.RentalNetIncome +
				test.model.Input.SocialSecurity +
				test.model.Input.WorkIncome +
				test.model.Input.YearlyWithdrawal

			assert.Equal(t, expected, actual)
		})
	}
}
