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

func (m *MockHalfOfSocialSecurity) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockHalfOfSocialSecurity) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockHalfOfSocialSecurity) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockHalfOfSocialSecurity) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var testHalfOfSocialSecurityTests = []struct {
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
	{
		name: "Test Case 1",
	},
}

func TestNewHalfOfSocialSecurity(t *testing.T) {
	actual := calculator.NewHalfOfSocialSecurity()
	expected := calculator.HalfOfSocialSecurity{}

	assert.Equal(t, expected, actual)
}

func TestMockHalfOfSocialSecurityCalculateTraditional(t *testing.T) {
	for _, test := range testHalfOfSocialSecurityTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.HalfOfSocialSecurity{}

			actual := c.CalculateTraditional(&test.model)
			expected := c.CalculateTraditionalRetirement(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockHalfOfSocialSecurityCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range testHalfOfSocialSecurityTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.HalfOfSocialSecurity{}

			actual := c.CalculateTraditional(&test.model)
			expected := test.model.Input.SocialSecurity * 0.5

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockHalfOfSocialSecurityCalculateRoth(t *testing.T) {
	for _, test := range testHalfOfSocialSecurityTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.HalfOfSocialSecurity{}

			actual := c.CalculateRoth(&test.model)
			expected := c.CalculateTraditional(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}

func TestMockHalfOfSocialSecurityCalculateRothRetirement(t *testing.T) {
	for _, test := range testHalfOfSocialSecurityTests {
		t.Run(test.name, func(t *testing.T) {
			c := &calculator.HalfOfSocialSecurity{}

			actual := c.CalculateRothRetirement(&test.model)
			expected := c.CalculateTraditional(&test.model)

			assert.Equal(t, expected, actual)
		})
	}
}
