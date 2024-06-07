package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockRothOrTraditionalDecision struct {
	mock.Mock
}

func (m *MockRothOrTraditionalDecision) Calculate(model *calculator.Model) string {
	args := m.Called(model)
	return args.Get(0).(string)
}

func (m *MockRothOrTraditionalDecision) CalculateRetirement(model *calculator.Model) string {
	return m.Calculate(model)
}

func TestNewRothOrTraditionalDecision(t *testing.T) {
	actual := calculator.NewRothOrTraditionalDecision()
	expected := calculator.RothOrTraditionalDecision{
		TotalDisbursementsCalculation: calculator.NewTotalDisbursements(),
	}

	assert.Equal(t, expected, actual)
}

func TestRothOrTraditionalDecisionCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                                    string
		totalDisbursementsRothRetirement        float64
		totalDisbursementsTraditionalRetirement float64
	}{
		{
			name:                                    "Test Case Higher TaxRateOfSavings",
			totalDisbursementsRothRetirement:        1,
			totalDisbursementsTraditionalRetirement: 0,
		},
		{
			name:                                    "Test Case Higher EffectiveTaxRateOnGross",
			totalDisbursementsRothRetirement:        0,
			totalDisbursementsTraditionalRetirement: 1,
		},
		{
			name:                                    "Test Case Equal",
			totalDisbursementsRothRetirement:        0,
			totalDisbursementsTraditionalRetirement: 0,
		},
	}

	testModel := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursments := new(MockTotalDisbursements)
			mockTotalDisbursments.On("CalculateTraditionalRetirement", &testModel).Return(test.totalDisbursementsTraditionalRetirement)
			mockTotalDisbursments.On("CalculateRothRetirement", &testModel).Return(test.totalDisbursementsRothRetirement)

			c := calculator.RothOrTraditionalDecision{
				TotalDisbursementsCalculation: mockTotalDisbursments,
			}

			actual := c.Calculate(&testModel)
			expected := func() string {
				if test.totalDisbursementsTraditionalRetirement >= test.totalDisbursementsRothRetirement {
					return "Traditional"
				}

				return "Roth"
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
