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

var totalInterestTests = []struct {
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

func TestNewTotalInterest(t *testing.T) {
	actual := calculator.NewTotalInterest()
	expected := calculator.TotalInterest{
		TotalDisbursementsCalculation: calculator.NewTotalDisbursements(),
		TotalContributionsCalculation: calculator.NewTotalContributions(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalInterestCalculateTraditional(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalInterestTests {
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

func TestTotalInterestCalculateTraditionalRetirement(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalInterestTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursements := new(MockTotalDisbursements)
			mockTotalContributions := new(MockTotalContributions)

			mockTotalDisbursements.On("CalculateTraditionalRetirement", &model).Return(test.totalDisbursements)
			mockTotalContributions.On("CalculateTraditionalRetirement", &model).Return(test.totalContributions)

			c := &calculator.TotalInterest{
				TotalDisbursementsCalculation: mockTotalDisbursements,
				TotalContributionsCalculation: mockTotalContributions,
			}

			actual := c.CalculateTraditionalRetirement(&model)
			expected := test.totalDisbursements - test.totalContributions

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalInterestCalculateRoth(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalInterestTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursements := new(MockTotalDisbursements)
			mockTotalContributions := new(MockTotalContributions)

			mockTotalDisbursements.On("CalculateRothRetirement", &model).Return(test.totalDisbursements)
			mockTotalContributions.On("CalculateRothRetirement", &model).Return(test.totalContributions)

			c := &calculator.TotalInterest{
				TotalDisbursementsCalculation: mockTotalDisbursements,
				TotalContributionsCalculation: mockTotalContributions,
			}

			actual := c.CalculateRoth(&model)
			expected := c.CalculateRothRetirement(&model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalInterestCalculateRothRetirement(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalInterestTests {
		t.Run(test.name, func(t *testing.T) {
			mockTotalDisbursements := new(MockTotalDisbursements)
			mockTotalContributions := new(MockTotalContributions)

			mockTotalDisbursements.On("CalculateRothRetirement", &model).Return(test.totalDisbursements)
			mockTotalContributions.On("CalculateRothRetirement", &model).Return(test.totalContributions)

			c := &calculator.TotalInterest{
				TotalDisbursementsCalculation: mockTotalDisbursements,
				TotalContributionsCalculation: mockTotalContributions,
			}

			actual := c.CalculateRothRetirement(&model)
			expected := test.totalDisbursements - test.totalContributions

			assert.Equal(t, expected, actual)
		})
	}
}
