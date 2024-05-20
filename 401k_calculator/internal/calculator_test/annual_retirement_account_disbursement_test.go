package test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockAnnualRetirementAccountDisbursement struct {
	mock.Mock
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockAnnualRetirementAccountDisbursement) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var annualRetirementAccountDisbursementTests = []struct {
	name                          string
	model                         calculator.Model
	effectiveTaxRateOnGross       float64
	taxOnTraditionalIRAWithdrawal float64
}{
	{
		name: "Test Case 0",
	},
	{
		name: "Test Case 1",
		model: calculator.Model{
			Input: calculator.Input{
				YearlyWithdrawal: 10000,
			},
		},
		effectiveTaxRateOnGross:       3333,
		taxOnTraditionalIRAWithdrawal: 4444,
	},
	{
		name: "Test Case 2",
		model: calculator.Model{
			Input: calculator.Input{
				YearlyWithdrawal: math.MaxFloat64,
			},
		},
		effectiveTaxRateOnGross:       math.MaxFloat64,
		taxOnTraditionalIRAWithdrawal: math.MaxFloat64,
	},
	{
		name: "Test Case 3",
		model: calculator.Model{
			Input: calculator.Input{
				YearlyWithdrawal: -math.MaxFloat64,
			},
		},
		effectiveTaxRateOnGross:       -math.MaxFloat64,
		taxOnTraditionalIRAWithdrawal: -math.MaxFloat64,
	},
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
	for _, test := range annualRetirementAccountDisbursementTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AnnualRetirementAccountDisbursement{}

			actual := c.CalculateTraditional(&test.model)
			expected := float64(0)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualRetirementAccountDisbursementCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range annualRetirementAccountDisbursementTests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateTraditionalRetirement", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := calculator.AnnualRetirementAccountDisbursement{
				TaxOnTraditionalIRAWithdrawalCalculation: mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				return test.model.Input.YearlyWithdrawal - test.taxOnTraditionalIRAWithdrawal
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualRetirementAccountDisbursementCalculateRoth(t *testing.T) {
	for _, test := range annualRetirementAccountDisbursementTests {
		t.Run(test.name, func(t *testing.T) {
			c := calculator.AnnualRetirementAccountDisbursement{}

			actual := c.CalculateRoth(&test.model)
			expected := float64(0)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAnnualRetirementAccountDisbursementCalculateRothRetirement(t *testing.T) {
	for _, test := range annualRetirementAccountDisbursementTests {
		t.Run(test.name, func(t *testing.T) {
			mockEffectiveTaxRateOnGross := new(MockEffectiveTaxRateOnGross)
			mockEffectiveTaxRateOnGross.On("CalculateRothRetirement", &test.model).Return(test.effectiveTaxRateOnGross)

			mockTaxOnTraditionalIRAWithdrawal := new(MockTaxOnTraditionalIRAWithdrawal)
			mockTaxOnTraditionalIRAWithdrawal.On("CalculateRothRetirement", &test.model).Return(test.taxOnTraditionalIRAWithdrawal)

			c := calculator.AnnualRetirementAccountDisbursement{
				EffectiveTaxRateOnGrossCalculation:       mockEffectiveTaxRateOnGross,
				TaxOnTraditionalIRAWithdrawalCalculation: mockTaxOnTraditionalIRAWithdrawal,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				return (test.model.Input.YearlyWithdrawal - test.taxOnTraditionalIRAWithdrawal) * (1 - test.effectiveTaxRateOnGross)
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
