package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockHalfOfSocialSecurity struct {
	mock.Mock
}

func (m *MockHalfOfSocialSecurity) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockHalfOfSocialSecurity) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestMockHalfOfSocialSecurityCalculate(t *testing.T) {
	tests := []struct {
		name  string
		model calculator.Model
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					SocialSecurity: 345987,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.HalfOfSocialSecurity{}

			actual := c.Calculate(test.model)
			expected := test.model.Input.SocialSecurity * 0.5

			assert.Equal(t, expected, actual)
		})
	}
}
