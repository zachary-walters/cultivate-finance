package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockTotalTaxableIncome struct {
	mock.Mock
}

func (m *MockTotalTaxableIncome) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncome) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncome) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockTotalTaxableIncome) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var totalTaxableIncomeTests = []struct {
	name                        string
	adjustedGrossIncome         float64
	socialSecurityTaxbaleIncome float64
}{
	{
		name:                        "Test Case 0",
		adjustedGrossIncome:         10000,
		socialSecurityTaxbaleIncome: 500,
	},
}

func TestNewTotalTaxableIncome(t *testing.T) {
	actual := calculator.NewTotalTaxableIncome()
	expected := calculator.TotalTaxableIncome{
		AdjustedGrossIncomeCalculation:                   calculator.NewAdjustedGrossIncome(),
		SocialSecurityTaxableIncomeIndividualCalculation: calculator.NewSocialSecurityTaxableIncomeIndividual(),
	}

	assert.Equal(t, expected, actual)
}

func TestTotalTaxableIncomeCalculateTraditional(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockSocialSecurityTaxableIncome := new(MockSocialSecurityTaxableIncome)

			mockAdjustedGrossIncome.On("CalculateTraditional", &model).Return(test.adjustedGrossIncome)
			mockSocialSecurityTaxableIncome.On("CalculateTraditional", &model).Return(test.socialSecurityTaxbaleIncome)

			c := &calculator.TotalTaxableIncome{
				AdjustedGrossIncomeCalculation:                   mockAdjustedGrossIncome,
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncome,
			}

			actual := c.CalculateTraditional(&model)
			expected := test.adjustedGrossIncome + test.socialSecurityTaxbaleIncome

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeCalculateTraditionalRetirement(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockSocialSecurityTaxableIncome := new(MockSocialSecurityTaxableIncome)

			mockAdjustedGrossIncome.On("CalculateTraditionalRetirement", &model).Return(test.adjustedGrossIncome)
			mockSocialSecurityTaxableIncome.On("CalculateTraditionalRetirement", &model).Return(test.socialSecurityTaxbaleIncome)

			c := &calculator.TotalTaxableIncome{
				AdjustedGrossIncomeCalculation:                   mockAdjustedGrossIncome,
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncome,
			}

			actual := c.CalculateTraditionalRetirement(&model)
			expected := test.adjustedGrossIncome + test.socialSecurityTaxbaleIncome

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeCalculateRoth(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockSocialSecurityTaxableIncome := new(MockSocialSecurityTaxableIncome)

			mockAdjustedGrossIncome.On("CalculateRoth", &model).Return(test.adjustedGrossIncome)
			mockSocialSecurityTaxableIncome.On("CalculateRoth", &model).Return(test.socialSecurityTaxbaleIncome)

			c := &calculator.TotalTaxableIncome{
				AdjustedGrossIncomeCalculation:                   mockAdjustedGrossIncome,
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncome,
			}

			actual := c.CalculateRoth(&model)
			expected := test.adjustedGrossIncome + test.socialSecurityTaxbaleIncome

			assert.Equal(t, expected, actual)
		})
	}
}

func TestTotalTaxableIncomeCalculateRothRetirement(t *testing.T) {
	model := calculator.Model{}

	for _, test := range totalTaxableIncomeTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockSocialSecurityTaxableIncome := new(MockSocialSecurityTaxableIncome)

			mockAdjustedGrossIncome.On("CalculateRothRetirement", &model).Return(test.adjustedGrossIncome)
			mockSocialSecurityTaxableIncome.On("CalculateRothRetirement", &model).Return(test.socialSecurityTaxbaleIncome)

			c := &calculator.TotalTaxableIncome{
				AdjustedGrossIncomeCalculation:                   mockAdjustedGrossIncome,
				SocialSecurityTaxableIncomeIndividualCalculation: mockSocialSecurityTaxableIncome,
			}

			actual := c.CalculateRothRetirement(&model)
			expected := test.adjustedGrossIncome + test.socialSecurityTaxbaleIncome

			assert.Equal(t, expected, actual)
		})
	}
}
