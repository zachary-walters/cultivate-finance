package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeTraditional struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeTraditional) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeTraditional) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestSocialSecurityTaxableIncomeTraditionalCalculate(t *testing.T) {
	tests := []struct {
		name                                             string
		model                                            calculator.Model
		socialSecurityTaxableIncomeIndividualTraditional float64
		socialSecurityTaxableIncomeJointTraditional      float64
	}{
		{
			name: "Test Case 0",
			model: calculator.Model{
				Input: calculator.Input{
					RetirementFilingStatus: "single",
				},
			},
		},
		{
			name: "Test Case 1",
			model: calculator.Model{
				Input: calculator.Input{
					RetirementFilingStatus: "married-seperate",
				},
			},
		},
		{
			name: "Test Case 2",
			model: calculator.Model{
				Input: calculator.Input{
					RetirementFilingStatus: "married-joint",
				},
			},
		},
		{
			name: "Test Case 3",
			model: calculator.Model{
				Input: calculator.Input{
					RetirementFilingStatus: "head-of-household",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockSocialSecurityTaxableIncomeIndividualTraditional := new(MockSocialSecurityTaxableIncomeIndividualTraditional)
			mockSocialSecurityTaxableIncomeJointTraditional := new(MockSocialSecurityTaxableIncomeJointTraditional)

			mockSocialSecurityTaxableIncomeIndividualTraditional.On("Calculate", test.model).Return(test.socialSecurityTaxableIncomeIndividualTraditional)
			mockSocialSecurityTaxableIncomeJointTraditional.On("Calculate", test.model).Return(test.socialSecurityTaxableIncomeJointTraditional)

			c := &calculator.SocialSecurityTaxableIncomeTraditional{
				SocialSecurityTaxableIncomeIndividualTraditionalCalculation: mockSocialSecurityTaxableIncomeIndividualTraditional,
				SocialSecurityTaxableIncomeJointTraditionalCalculation:      mockSocialSecurityTaxableIncomeJointTraditional,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividualTraditional
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividualTraditional
				case "married-joint":
					return test.socialSecurityTaxableIncomeJointTraditional
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJointTraditional
				default:
					return 0
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
