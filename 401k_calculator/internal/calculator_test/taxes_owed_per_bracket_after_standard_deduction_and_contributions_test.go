package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTaxesOwedPerBracketAfterStandardDeductionAndContributions struct {
	mock.Mock
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) Calculate(model calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockTaxesOwedPerBracketAfterStandardDeductionAndContributions) CalculateRetirement(model calculator.Model) []float64 {
	return m.Calculate(model)
}

func TestTaxesOwedPerBracketAfterStandardDeductionAndContributions(t *testing.T) {
	tests := []struct {
		name                                                                     string
		model                                                                    calculator.Model
		taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle          []float64
		taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint    []float64
		taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate []float64
		taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold []float64
	}{
		{
			name: "Test Case 0",
			taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "single",
				},
			},
		},
		{
			name: "Test Case 1",
			taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-joint",
				},
			},
		},
		{
			name: "Test Case 2",
			taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "married-seperate",
				},
			},
		},
		{
			name: "Test Case 3",
			taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold: []float64{1, 2, 3, 4, 5},
			model: calculator.Model{
				Input: calculator.Input{
					CurrentFilingStatus: "head-of-household",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold := new(MockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsSingle)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate)
			mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.On("Calculate", test.model).Return(test.taxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold)

			c := calculator.TaxesOwedPerBracketAfterStandardDeductionAndContributions{
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsSingleCalculation:          mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJointCalculation:    mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperateCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate,
				TaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHouseholdCalculation: mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold,
			}

			actual := c.Calculate(test.model)
			expected := func() []float64 {
				switch test.model.Input.CurrentFilingStatus {
				case "single":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsSingle.Calculate(test.model)
				case "married-joint":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedJoint.Calculate(test.model)
				case "married-seperate":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsMarriedSeperate.Calculate(test.model)
				case "head-of-household":
					return mockTaxesOwedPerBracketAfterStandardDeductionAndContributionsHeadOfHousehold.Calculate(test.model)
				default:
					return nil
				}
			}()

			assert.Equal(t, expected, actual)
			assert.NotEmpty(t, actual)
		})
	}
}
