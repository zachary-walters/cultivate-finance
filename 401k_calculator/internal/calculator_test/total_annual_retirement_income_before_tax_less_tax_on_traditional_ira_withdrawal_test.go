package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal struct {
	mock.Mock
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalAnnualRetirementIncomeBeforeTaxLessTaxOnTraditionalIRAWithdrawal) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}
