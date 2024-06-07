package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockSocialSecurityTaxableIncomeIndividual struct {
	mock.Mock
}

func (m *MockSocialSecurityTaxableIncomeIndividual) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeIndividual) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeIndividual) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockSocialSecurityTaxableIncomeIndividual) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

var socialSecurityTaxableIncomeIndividualTests = []struct {
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
			SocialSecurityTaxRatesIndividual: []calculator.TaxRate{
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
			SocialSecurityTaxRatesIndividual: []calculator.TaxRate{
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
			SocialSecurityTaxRatesIndividual: []calculator.TaxRate{
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

func TestNewSocialSecurityTaxableIncomeIndividual(t *testing.T) {
	actual := calculator.NewSocialSecurityTaxableIncomeIndividual()
	expected := calculator.SocialSecurityTaxableIncomeIndividual{
		AdjustedGrossIncomeCalculation: calculator.NewAdjustedGrossIncome(),
	}

	assert.Equal(t, expected, actual)
}

func TestSocialSecurityTaxableIncomeIndividualCalculateTraditional(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeIndividualTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateTraditional", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeIndividual{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateTraditional(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesIndividual) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesIndividual {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesIndividual[len(test.model.SocialSecurityTaxRatesIndividual)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeIndividualCalculateTraditionalRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeIndividualTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateTraditionalRetirement", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeIndividual{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateTraditionalRetirement(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesIndividual) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesIndividual {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesIndividual[len(test.model.SocialSecurityTaxRatesIndividual)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeIndividualCalculateRoth(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeIndividualTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateRoth", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeIndividual{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateRoth(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesIndividual) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesIndividual {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesIndividual[len(test.model.SocialSecurityTaxRatesIndividual)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}

func TestSocialSecurityTaxableIncomeIndividualCalculateRothRetirement(t *testing.T) {
	for _, test := range socialSecurityTaxableIncomeIndividualTests {
		t.Run(test.name, func(t *testing.T) {
			mockAdjustedGrossIncome := new(MockAdjustedGrossIncome)
			mockAdjustedGrossIncome.On("CalculateRothRetirement", &test.model).Return(test.adjustedGrossIncome)

			c := &calculator.SocialSecurityTaxableIncomeIndividual{
				AdjustedGrossIncomeCalculation: mockAdjustedGrossIncome,
			}

			actual := c.CalculateRothRetirement(&test.model)
			expected := func() float64 {
				if len(test.model.SocialSecurityTaxRatesIndividual) == 0 {
					return 0.0
				}

				for _, taxRate := range test.model.SocialSecurityTaxRatesIndividual {
					if taxRate.Cap > test.adjustedGrossIncome {
						return test.model.Input.SocialSecurity * taxRate.Rate
					}
				}

				return test.model.SocialSecurityTaxRatesIndividual[len(test.model.SocialSecurityTaxRatesIndividual)-1].Rate * test.model.Input.SocialSecurity
			}()

			assert.Equal(t, expected, actual)
		})
	}
}
