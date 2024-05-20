package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalInterest struct {
	mock.Mock
}

func (m *MockTotalInterest) Calculate(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestMockTotalInterestCalculateTraditional(t *testing.T) {
	tests := []struct {
		name               string
		totalDisbursements float64
		totalContributions float64
	}{
		{
			name:               "Test Case 0",
			totalDisbursements: 100000,
			totalContributions: 1,
		},
	}

	model := calculator.Model{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursements := new(MockTotalDisbursements)
			mockTotalContributions := new(MockTotalContributions)

			mockTotalDisbursements.On("CalculateTraditionalRetirement", &model).Return(test.totalDisbursements)
			mockTotalContributions.On("CalculateTraditionalRetirement", &model).Return(test.totalContributions)

			c := &calculator.TotalInterest{
				TotalDisbursementsCalculation: mockTotalDisbursements,
				TotalContributionsCalculation: mockTotalContributions,
			}

			actual := c.CalculateTraditional(&model)
			expected := c.CalculateTraditionalRetirement(&model)

			assert.Equal(t, expected, actual)
		})
	}
}
