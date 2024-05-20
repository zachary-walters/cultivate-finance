package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalDisbursements struct {
	mock.Mock
}

func (m *MockTotalDisbursements) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalDisbursements) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func TestTotalDisburesmentsCalculateTraditional(t *testing.T) {
	tests := []struct {
		name                string
		model               calculator.Model
		balancesTraditional calculator.ChartData
	}{
		{
			name: "Test case 0",
			balancesTraditional: calculator.ChartData{
				AfterTaxIncome: map[int32]float64{
					0: 100,
					1: 200,
					3: 300,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			c := &calculator.TotalDisbursements{}

			actual := c.CalculateTraditional(&test.model)
			expected := float64(0)

			assert.Equal(t, expected, actual)
		})
	}
}
