package calculator

type constants struct {
	InflationRate                        float64
	SingleTaxRates                       []TaxRate
	MarriedJointTaxRates                 []TaxRate
	MarriedSeperateTaxRates              []TaxRate
	HeadOfHouseholdTaxRates              []TaxRate
	STANDARD_DEDUCTION_SINGLE            float64
	STANDARD_DEDUCTION_MARRIED_JOINT     float64
	STANDARD_DEDUCTION_MARRIED_SEPERATE  float64
	STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD float64
}

var Constants = constants{
	InflationRate: 0.03,
	SingleTaxRates: []TaxRate{
		{
			Cap:  10275,
			Rate: 0.10,
		},
		{
			Cap:  41775,
			Rate: 0.12,
		},
		{
			Cap:  89075,
			Rate: 0.22,
		},
		{
			Cap:  170050,
			Rate: 0.24,
		},
		{
			Cap:  215950,
			Rate: 0.32,
		},
		{
			Cap:  539900,
			Rate: 0.35,
		},
		{
			Cap:  -1,
			Rate: 0.37,
		},
	},
	MarriedJointTaxRates: []TaxRate{
		{
			Cap:  20550,
			Rate: 0.10,
		},
		{
			Cap:  83550,
			Rate: 0.12,
		},
		{
			Cap:  178150,
			Rate: 0.22,
		},
		{
			Cap:  340100,
			Rate: 0.24,
		},
		{
			Cap:  431900,
			Rate: 0.32,
		},
		{
			Cap:  647850,
			Rate: 0.35,
		},
		{
			Cap:  -1,
			Rate: 0.37,
		},
	},
	MarriedSeperateTaxRates: []TaxRate{
		{
			Cap:  10275,
			Rate: 0.10,
		},
		{
			Cap:  41775,
			Rate: 0.12,
		},
		{
			Cap:  89075,
			Rate: 0.22,
		},
		{
			Cap:  170050,
			Rate: 0.24,
		},
		{
			Cap:  215950,
			Rate: 0.32,
		},
		{
			Cap:  323925,
			Rate: 0.35,
		},
		{
			Cap:  -1,
			Rate: 0.37,
		},
	},
	HeadOfHouseholdTaxRates: []TaxRate{
		{
			Cap:  14650,
			Rate: 0.10,
		},
		{
			Cap:  55900,
			Rate: 0.12,
		},
		{
			Cap:  89050,
			Rate: 0.22,
		},
		{
			Cap:  170050,
			Rate: 0.24,
		},
		{
			Cap:  215950,
			Rate: 0.32,
		},
		{
			Cap:  539900,
			Rate: 0.35,
		},
		{
			Cap:  -1,
			Rate: 0.37,
		},
	},
	STANDARD_DEDUCTION_SINGLE:            13850,
	STANDARD_DEDUCTION_MARRIED_JOINT:     27700,
	STANDARD_DEDUCTION_MARRIED_SEPERATE:  13850,
	STANDARD_DEDUCTION_HEAD_OF_HOUSEHOLD: 20800,
}
