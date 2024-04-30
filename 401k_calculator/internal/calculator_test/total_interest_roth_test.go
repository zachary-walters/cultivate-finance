package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalInterestRoth struct {
	mock.Mock
}

func (m *MockTotalInterestRoth) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestMockTotalInterestRothCalculate(t *testing.T) {
	tests := []struct {
		name                   string
		totalDisbursementsRoth float64
		totalContributions     float64
	}{
		{
			name:                   "Test Case 0",
			totalDisbursementsRoth: 100000,
			totalContributions:     1,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursementsRoth := new(MockTotalDisbursementsRothMatchingNet)
			mockTotalContributionsRoth := new(MockTotalContributionsRoth)

			mockTotalDisbursementsRoth.On("Calculate", model).Return(test.totalDisbursementsRoth)
			mockTotalContributionsRoth.On("Calculate", model).Return(test.totalContributions)

			c := &calculator.TotalInterestRoth{
				TotalDisbursementsRothMatchingNetCalculation: mockTotalDisbursementsRoth,
				TotalContributionsRothCalculation:            mockTotalContributionsRoth,
			}

			actual := c.Calculate(model)
			expected := test.totalDisbursementsRoth - test.totalContributions

			assert.Equal(t, expected, actual)
		})
	}
}
