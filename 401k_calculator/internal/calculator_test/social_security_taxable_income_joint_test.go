package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeJoint struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeJoint) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeJoint) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeJoint) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeJoint) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var socialSecurityTaxableIncomeJointTests = []struct {
	name                string
	model               calculator.Model
	adjustedGrossIncome float64
}{
	{
		name: "Test Case 0",
		model: calculator.Model{
			Input: calculator.Input{
				SocialSecurity: 100000,
			},
			SocialSecurityTaxRatesJoint: []calculator.TaxRate{
				{
					Cap:  25000,
					Rate: 0.0,
				},
				{
					Cap:  34000,
					Rate: 0.5,
				},
				{
					Cap:  -1,
					Rate: 0.85,
				},
			},
		},
		adjustedGrossIncome: 0,
	},
	{
		name: "Test Case 1",
		model: calculator.Model{
			Input: calculator.Input{
				SocialSecurity: 100000,
			},
			SocialSecurityTaxRatesJoint: []calculator.TaxRate{
				{
					Cap:  25000,
					Rate: 0.0,
				},
				{
					Cap:  34000,
					Rate: 0.5,
				},
				{
					Cap:  -1,
					Rate: 0.85,
				},
			},
		},
		adjustedGrossIncome: 26000,
	},
	{
		name: "Test Case 2",
		model: calculator.Model{
			Input: calculator.Input{
				SocialSecurity: 100000,
			},
			SocialSecurityTaxRatesJoint: []calculator.TaxRate{
				{
					Cap:  25000,
					Rate: 0.0,
				},
				{
					Cap:  34000,
					Rate: 0.5,
				},
				{
					Cap:  -1,
					Rate: 0.85,
				},
			},
		},
		adjustedGrossIncome: 35000,
	},
	{
		name: "Test Case 3",
	},
}

func TestNewSocialSecurityTaxableIncomeJoint(t *testing.T) {
	actual := calculator.NewSocialSecurityTaxableIncomeJoint()
	expected := calculator.SocialSecurityTaxableIncomeJoint{
		AdjustedGrossIncomeCalculation: calculator.NewAdjustedGrossIncome(),
	}

	assert.Equal(t, expected, actual)
}

func TestSocialSecurityTaxableIncomeJointCalculateTraditional(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateTraditional", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeJoint{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesJoint) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesJoint {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesJoint[len(test.model.SocialSecurityTaxRatesJoint)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeJointCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateTraditionalRetirement", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeJoint{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesJoint) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesJoint {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesJoint[len(test.model.SocialSecurityTaxRatesJoint)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeJointCalculateRoth(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateRoth", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeJoint{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesJoint) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesJoint {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesJoint[len(test.model.SocialSecurityTaxRatesJoint)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeJointCalculateRothRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeJointTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateRothRetirement", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeJoint{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesJoint) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesJoint {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesJoint[len(test.model.SocialSecurityTaxRatesJoint)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
