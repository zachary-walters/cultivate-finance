package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncome struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncome) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncome) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncome) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncome) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var socialSecurityTaxableIncomeTests = []struct {
	name                                  string
	model                                 calculator.Model
	socialSecurityTaxableIncomeIndividual float64
	socialSecurityTaxableIncomeJoint      float64
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
	{
		name: "Test Case 4",
	},
}

func TestNewSocialSecurityTaxableIncome(t *testing.T) {
	actual := calculator.NewSocialSecurityTaxableIncome()
	expected := calculator.SocialSecurityTaxableIncome{
		SocialSecurityTaxableIncomeIndividualCalculation: calculator.NewSocialSecurityTaxableIncomeIndividual(),
		SocialSecurityTaxableIncomeJointCalculation:      calculator.NewSocialSecurityTaxableIncomeJoint(),
	}

	assert.Equal(t, expected, actual)
}

func TestSocialSecurityTaxableIncomeTraditionalCalculateTraditional(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockSocialSecurityTaxableIncomeIndividual := new(MockSocialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint := new(MockSocialSecurityTaxableIncomeJoint)

			mockSocialSecurityTaxableIncomeIndividual.On("CalculateTraditional", &test.model).Return(test.socialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint.On("CalculateTraditional", &test.model).Return(test.socialSecurityTaxableIncomeJoint)

			c := &calculator.SocialSecurityTaxableIncome{
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncomeIndividual,
				SocialSecurityTaxableIncomeJointCalculation:      mockSocialSecurityTaxableIncomeJoint,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-joint":
					return test.socialSecurityTaxableIncomeJoint
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJoint
				default:
					return test.socialSecurityTaxableIncomeIndividual
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeTraditionalCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockSocialSecurityTaxableIncomeIndividual := new(MockSocialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint := new(MockSocialSecurityTaxableIncomeJoint)

			mockSocialSecurityTaxableIncomeIndividual.On("CalculateTraditionalRetirement", &test.model).Return(test.socialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint.On("CalculateTraditionalRetirement", &test.model).Return(test.socialSecurityTaxableIncomeJoint)

			c := &calculator.SocialSecurityTaxableIncome{
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncomeIndividual,
				SocialSecurityTaxableIncomeJointCalculation:      mockSocialSecurityTaxableIncomeJoint,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-joint":
					return test.socialSecurityTaxableIncomeJoint
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJoint
				default:
					return test.socialSecurityTaxableIncomeIndividual
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeTraditionalCalculateRoth(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockSocialSecurityTaxableIncomeIndividual := new(MockSocialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint := new(MockSocialSecurityTaxableIncomeJoint)

			mockSocialSecurityTaxableIncomeIndividual.On("CalculateRoth", &test.model).Return(test.socialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint.On("CalculateRoth", &test.model).Return(test.socialSecurityTaxableIncomeJoint)

			c := &calculator.SocialSecurityTaxableIncome{
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncomeIndividual,
				SocialSecurityTaxableIncomeJointCalculation:      mockSocialSecurityTaxableIncomeJoint,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-joint":
					return test.socialSecurityTaxableIncomeJoint
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJoint
				default:
					return test.socialSecurityTaxableIncomeIndividual
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeTraditionalCalculateRothRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockSocialSecurityTaxableIncomeIndividual := new(MockSocialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint := new(MockSocialSecurityTaxableIncomeJoint)

			mockSocialSecurityTaxableIncomeIndividual.On("CalculateRothRetirement", &test.model).Return(test.socialSecurityTaxableIncomeIndividual)
			mockSocialSecurityTaxableIncomeJoint.On("CalculateRothRetirement", &test.model).Return(test.socialSecurityTaxableIncomeJoint)

			c := &calculator.SocialSecurityTaxableIncome{
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncomeIndividual,
				SocialSecurityTaxableIncomeJointCalculation:      mockSocialSecurityTaxableIncomeJoint,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				switch test.model.Input.RetirementFilingStatus {
				case "single":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-seperate":
					return test.socialSecurityTaxableIncomeIndividual
				case "married-joint":
					return test.socialSecurityTaxableIncomeJoint
				case "head-of-household":
					return test.socialSecurityTaxableIncomeJoint
				default:
					return test.socialSecurityTaxableIncomeIndividual
				}
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
