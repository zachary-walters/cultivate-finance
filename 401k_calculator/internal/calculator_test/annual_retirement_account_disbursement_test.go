package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAnnualRetirementAccountDisbursement struct {
	mock.Mock
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateTraditional(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateTraditionalRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateRoth(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateRothRetirement(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestNewAnnualRetirementAccountDisbursement(t *testing.T) {
	actual := calculator.NewAnnualRetirementAccountDisbursement()
	expected := calculator.AnnualRetirementAccountDisbursement{
		EffectiveTaxRateOnGrossCalculation:       calculator.NewEffectiveTaxRateOnGross(),
		TaxOnTraditionalIRAWithdrawalCalculation: calculator.NewTaxOnTraditionalIRAWithdrawal(),
	}

	assert.Equal(t, actual, expected)
}

func TestAnnualRetirementAccountDisbursementCalculateTraditional(t *testing.T) {
	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case 0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AnnualRetirementAccountDisbursement{}

			actual := c.CalculateTraditional(test.model)
			expected := float64(0)

			assert.Equal(t, expected, actual)
		})
	}
}
