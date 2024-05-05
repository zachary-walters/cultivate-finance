package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeRoth struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeRoth) Calculate(model calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeRoth) CalculateRetirement(model calculator.Model) float64 {
	return m.Calculate(model)
}

func TestSocialSecurityTaxableIncomeRothCalculate(t *testing.T) {
	tests := []struct {
		name                                      string
		model                                     calculator.Model
		socialSecurityTaxableIncomeIndividualRoth float64
		socialSecurityTaxableIncomeJointRoth      float64
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
			mockSocialSecurityTaxableIncomeIndividualRoth := new(MockSocialSecurityTaxableIncomeIndividualRoth)
			mockSocialSecurityTaxableIncomeJointRoth := new(MockSocialSecurityTaxableIncomeJointRoth)

			mockSocialSecurityTaxableIncomeIndividualRoth.On("Calculate", test.model).Return(test.socialSecurityTaxableIncomeIndividualRoth)
			mockSocialSecurityTaxableIncomeJointRoth.On("Calculate", test.model).Return(test.socialSecurityTaxableIncomeJointRoth)

			c := &calculator.SocialSecurityTaxableIncomeRoth{
				SocialSecurityTaxableIncomeIndividualRothCalculation: mockSocialSecurityTaxableIncomeIndividualRoth,
				SocialSecurityTaxableIncomeJointRothCalculation:      mockSocialSecurityTaxableIncomeJointRoth,
			}

			actual := c.Calculate(test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividualRoth
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividualRoth
				case "married-joint":
					return test.socialSecurityTaxableIncomeJointRoth
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJointRoth
				default:
					return 0
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
