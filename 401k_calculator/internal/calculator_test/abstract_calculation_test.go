package test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zachary-walters/rothvtrad/401k_calculator/internal/calculator"
)

type MockCalculation struct {
	mock.Mock
}

func (m *MockCalculation) CalculateTraditional(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCalculation) CalculateTraditionalRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCalculation) CalculateRoth(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

func (m *MockCalculation) CalculateRothRetirement(model *calculator.Model) float64 {
	args := m.Called(model)
	return args.Get(0).(float64)
}

type MockSequenceCalculation struct{ mock.Mock }

func (m *MockSequenceCalculation) CalculateTraditional(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockSequenceCalculation) CalculateTraditionalRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockSequenceCalculation) CalculateRoth(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

func (m *MockSequenceCalculation) CalculateRothRetirement(model *calculator.Model) []float64 {
	args := m.Called(model)
	return args.Get(0).([]float64)
}

type MockChartCalculation struct{ mock.Mock }

func (m *MockChartCalculation) Calculate(model *calculator.Model) calculator.ChartData {
	args := m.Called(model)
	return args.Get(0).(calculator.ChartData)
}

type MockDecisionCalculation struct{ mock.Mock }

func (m *MockDecisionCalculation) Calculate(model *calculator.Model) string {
	args := m.Called(model)
	return args.Get(0).(string)
}

func TestNewModel(t *testing.T) {
	input := calculator.Input{}

	model := calculator.NewModel(input)

	assert.Equal(t, input, model.Input)

	input = calculator.Input{
		CurrentAge:                8975,
		RetirementAge:             392,
		RetirementFilingStatus:    "status",
		CurrentFilingStatus:       "status1",
		AnnualContributionsPreTax: 534532,
		AnnuityIncome:             334225,
	}

	model = calculator.NewModel(input)

	assert.Equal(t, input, model.Input)
	assert.Equal(t, calculator.Constants.InflationRate, model.InflationRate)
	assert.Equal(t, calculator.Constants.SingleTaxRates, model.SingleTaxRates)
	assert.Equal(t, calculator.Constants.MarriedJointTaxRates, model.MarriedJointTaxRates)
	assert.Equal(t, calculator.Constants.MarriedSeparateTaxRates, model.MarriedSeparateTaxRates)
	assert.Equal(t, calculator.Constants.HeadOfHouseholdTaxRates, model.HeadOfHouseholdTaxRates)
	assert.Equal(t, calculator.Constants.SocialSecurityTaxRatesIndividual, model.SocialSecurityTaxRatesIndividual)
	assert.Equal(t, calculator.Constants.SocialSecurityTaxRatesJoint, model.SocialSecurityTaxRatesJoint)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_SINGLE, model.STANDARD_DEDUCTION_SINGLE)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_MARRIED_JOINT, model.STANDARD_DEDUCTION_MARRIED_JOINT)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_MARRIED_SEPERATE, model.STANDARD_DEDUCTION_MARRIED_SEPERATE)
	assert.Equal(t, calculator.Constants.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD, model.STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD)
}

func TestCalculateSynchronous(t *testing.T) {
	model := &calculator.Model{}
	datakey := "testKey"

	t.Run("Test with Calculation type", func(t *testing.T) {
		calculation := &MockCalculation{}
		calculation.On("CalculateTraditional", model).Return(1.0)
		calculation.On("CalculateTraditionalRetirement", model).Return(2.0)
		calculation.On("CalculateRoth", model).Return(3.0)
		calculation.On("CalculateRothRetirement", model).Return(4.0)

		expected := calculator.CalculationData{
			Datakey:                    datakey,
			TraditionalValue:           calculation.CalculateTraditional(model),
			TraditionalRetirementValue: calculation.CalculateTraditionalRetirement(model),
			RothValue:                  calculation.CalculateRoth(model),
			RothRetirementValue:        calculation.CalculateRothRetirement(model),
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with SequenceCalculation type", func(t *testing.T) {
		calculation := &MockSequenceCalculation{}
		calculation.On("CalculateTraditional", model).Return([]float64{1.0})
		calculation.On("CalculateTraditionalRetirement", model).Return([]float64{2.0})
		calculation.On("CalculateRoth", model).Return([]float64{3.0})
		calculation.On("CalculateRothRetirement", model).Return([]float64{4.0})

		expected := calculator.CalculationData{
			Datakey:                    datakey,
			TraditionalValue:           calculation.CalculateTraditional(model),
			TraditionalRetirementValue: calculation.CalculateTraditionalRetirement(model),
			RothValue:                  calculation.CalculateRoth(model),
			RothRetirementValue:        calculation.CalculateRothRetirement(model),
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with ChartCalculation type", func(t *testing.T) {
		calculation := &MockChartCalculation{}
		calculation.On("Calculate", model).Return(calculator.ChartData{})
		expected := calculator.CalculationData{
			Datakey:                    datakey,
			TraditionalValue:           calculation.Calculate(model),
			TraditionalRetirementValue: nil,
			RothValue:                  nil,
			RothRetirementValue:        nil,
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with DecisionCalculation type", func(t *testing.T) {
		calculation := &MockDecisionCalculation{}
		calculation.On("Calculate", model).Return("investing is important")
		expected := calculator.CalculationData{
			Datakey:                    datakey,
			TraditionalValue:           calculation.Calculate(model),
			TraditionalRetirementValue: nil,
			RothValue:                  nil,
			RothRetirementValue:        nil,
		}
		result := calculator.CalculateSynchronous(model, calculation, datakey)
		assert.Equal(t, expected, result)
	})
}

func TestCalculateAsync(t *testing.T) {
	model := &calculator.Model{}
	datakey := "testKey"

	t.Run("Test with Calculation type", func(t *testing.T) {
		calculation := &MockCalculation{}
		calculation.On("CalculateTraditional", model).Return(1.0)
		calculation.On("CalculateTraditionalRetirement", model).Return(2.0)
		calculation.On("CalculateRoth", model).Return(3.0)
		calculation.On("CalculateRothRetirement", model).Return(4.0)

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsync(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})
}

func TestCalculateAsyncWasm(t *testing.T) {
	model := &calculator.Model{}
	datakey := "testKey"

	t.Run("Test with Calculation type", func(t *testing.T) {
		calculation := &MockCalculation{}
		calculation.On("CalculateTraditional", model).Return(1.0)
		calculation.On("CalculateTraditionalRetirement", model).Return(2.0)
		calculation.On("CalculateRoth", model).Return(3.0)
		calculation.On("CalculateRothRetirement", model).Return(4.0)

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})

	t.Run("Test with SequenceCalculation type", func(t *testing.T) {
		calculation := &MockSequenceCalculation{}
		calculation.On("CalculateTraditional", model).Return([]float64{1.0})
		calculation.On("CalculateTraditionalRetirement", model).Return([]float64{2.0})
		calculation.On("CalculateRoth", model).Return([]float64{3.0})
		calculation.On("CalculateRothRetirement", model).Return([]float64{4.0})

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})

	t.Run("Test with ChartCalculation type", func(t *testing.T) {
		calculation := &MockChartCalculation{}
		calculation.On("Calculate", model).Return(calculator.ChartData{})

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})

	t.Run("Test with Decision type", func(t *testing.T) {
		calculation := &MockDecisionCalculation{}
		calculation.On("Calculate", model).Return("investing is important")

		expected := calculator.CalculateSynchronousWasm(model, calculation, datakey)

		wg := &sync.WaitGroup{}
		ch := make(chan calculator.CalculationData)
		wg.Add(1)

		go calculator.CalculateAsyncWasm(wg, ch, datakey, calculation, model)

		result := <-ch
		close(ch)

		assert.Equal(t, expected, result)
	})
}

func TestTranslateChartMap(t *testing.T) {
	t.Run("Test with non-empty map", func(t *testing.T) {
		m := map[int32]float64{
			1: 1.1,
			2: 2.2,
			3: 3.3,
		}
		expected := map[string]interface{}{
			"1": 1.1,
			"2": 2.2,
			"3": 3.3,
		}
		result := calculator.TranslateChartMap(m)
		assert.Equal(t, expected, result)
	})

	t.Run("Test with empty map", func(t *testing.T) {
		m := map[int32]float64{}
		expected := map[string]interface{}{}
		result := calculator.TranslateChartMap(m)
		assert.Equal(t, expected, result)
	})
}
