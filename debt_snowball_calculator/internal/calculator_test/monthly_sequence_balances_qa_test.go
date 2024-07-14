package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachary-walters/cultivate-finance/debt_snowball_calculator/internal/calculator"
)

var monthlySequenceBalancesQATests = []struct {
	name               string
	model              *calculator.Model
	expected           []float64
	debtPayoff         float64
	snowball           calculator.DebtSequences
	totalBeginningDebt float64
}{
	{
		name:               "Test Case 0",
		model:              &calculator.Model{},
		expected:           monthlySequenceBalancesQATestExpected1,
		debtPayoff:         30,
		snowball:           monthlySequenceBalancesQATestDebtSequence1,
		totalBeginningDebt: 38000,
	},
}

func TestMonthlySequenceBalancesQA(t *testing.T) {
	for _, test := range monthlySequenceBalancesQATests {
		t.Run(test.name, func(t *testing.T) {
			mockDebtPayoff := new(MockCalculation)
			mockSnowball := new(MockSnowballCalculation)
			mockTotalBeginningDebt := new(MockCalculation)

			mockDebtPayoff.On("CalculateSnowball", test.model).Return(test.debtPayoff)
			mockSnowball.On("CalculateSnowball", test.model).Return(test.snowball)
			mockTotalBeginningDebt.On("CalculateSnowball", test.model).Return(test.totalBeginningDebt)

			c := &calculator.MonthlySequenceBalances{
				DebtPayoffMonthCalculation:    mockDebtPayoff,
				SnowballCalculation:           mockSnowball,
				TotalBeginningDebtCalculation: mockTotalBeginningDebt,
			}

			actual := c.CalculateSnowball(test.model)

			assert.Equal(t, test.expected, actual)
		})
	}
}

var monthlySequenceBalancesQATestDebtSequence1 = calculator.DebtSequences{
	calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt0", Amount: 1000, AnnualInterest: 19.49, MinimumPayment: 50}, Months: []float64{
			1, 2, 3, 4, 5}, Payments: []float64{
			550, 150, 150, 150, 7.507879127317551}, Balances: []float64{
			450, 304.8725, 157.38788752083335, 7.507879127317547, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt1", Amount: 2000, AnnualInterest: 24.49, MinimumPayment: 100}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, Payments: []float64{
			100, 100, 100, 100, 242.49212087268245, 250, 250, 250, 250, 250, 250, 75.31534439032339}, Balances: []float64{
			1900, 1836.735, 1772.1788667916665, 1706.3052504981063, 1493.6871159125292, 1269.0686971364441, 1039.8661907971705, 805.9860433076894, 567.3327918081939, 323.80902520101273, 75.3153443903234, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt2", Amount: 3000, AnnualInterest: 26.99, MinimumPayment: 150}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, Payments: []float64{
			150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 324.6846556096766, 400, 400, 400, 400, 25.337132508502293}, Balances: []float64{
			2850, 2760.7275, 2669.4471126875, 2576.113677330363, 2480.6810174563184, 2383.1019180072735, 2283.3281019797873, 2181.3102065401495, 2076.9977586022483, 1970.3391498561439, 1861.2816112349917, 1571.1575821522545, 1197.4988681041623, 815.435946812605, 424.77979364966524, 25.33713250850229, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt3", Amount: 4000, AnnualInterest: 24.29, MinimumPayment: 200}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, Payments: []float64{
			200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 574.6628674914978, 600, 452.6703899244595}, Balances: []float64{
			3800, 3672.87, 3543.1666769166663, 3410.8379424019213, 3275.830653752707, 3138.0905925690845, 2997.562442980337, 2854.1897694303307, 2707.914994013216, 2558.679373350367, 2406.4229749992674, 2251.0846533848776, 2092.6020252438097, 1930.911444571453, 1765.9479770619869, 1597.6453740310167, 1043.689377442723, 452.6703899244595, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt4", Amount: 5000, AnnualInterest: 29.99, MinimumPayment: 200}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}, Payments: []float64{
			200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 347.3296100755405, 800, 800, 800, 434.4871454082128}, Balances: []float64{
			4800, 4714.961666666667, 4627.798083652779, 4538.456137426734, 4446.881387061258, 4353.018031059563, 4256.80887335246, 4158.1952884456605, 4057.117185696065, 3953.5129726952523, 3847.319517737861, 3738.472111351993, 3626.9044268681982, 3512.5484800030126, 3395.3345874324214, 3275.191324330003, 3152.045480843884, 3025.822017485974, 2745.432396825632, 1994.0519948096328, 1223.8933442465836, 434.4871454082128, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt5", Amount: 6000, AnnualInterest: 15, MinimumPayment: 200}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}, Payments: []float64{
			200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 565.5128545917871, 1000, 1000, 178.0324410088822}, Balances: []float64{
			5800, 5670, 5538.375, 5405.1046875, 5270.16849609375, 5133.545602294922, 4995.214922323608, 4855.155108852653, 4713.344547713311, 4569.761354559727, 4424.383371491724, 4277.18816363537, 4128.153015680812, 3977.254928376822, 3824.4706149815324, 3669.7764976688018, 3513.1487038896616, 3354.5630626882826, 3193.995100971886, 3031.4200397340346, 2866.81279023071, 2700.147950108594, 2161.3180342107676, 1175.834509638402, 178.03244100888213, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt6", Amount: 7000, AnnualInterest: 18, MinimumPayment: 200}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, Payments: []float64{
			200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 200, 1021.9675589911178, 1200, 1200, 563.6067705432334}, Balances: []float64{
			6800, 6699, 6596.485, 6492.432275, 6386.818759125, 6279.6210405118745, 6170.815356119552, 6060.377586461345, 5948.283250258265, 5834.507499012139, 5719.025111497322, 5601.810488169782, 5482.837645492328, 5362.080210174713, 5239.511413327334, 5115.104084527244, 4988.830645795152, 4860.663105482079, 4730.57305206431, 4598.531647845275, 4464.509622562954, 4328.477266901398, 4190.40442590492, 4050.2604922934934, 3908.0143996778957, 2929.3375432970793, 1755.2776064465354, 563.6067705432334, 0}}, calculator.DebtSequence{
		Debt: calculator.Debt{
			Name: "debt7", Amount: 10000, AnnualInterest: 16, MinimumPayment: 350}, Months: []float64{
			1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, Payments: []float64{
			350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 350, 986.3932294567666, 1394.7511709901887}, Balances: []float64{
			9650, 9424, 9194.986666666666, 8962.919822222222, 8727.758753185186, 8489.462203227655, 8247.988365937357, 8003.294877483188, 7755.338809182965, 7504.076659972071, 7249.4643487716985, 6991.457206755321, 6730.009969512059, 6465.076769105553, 6196.61112602696, 5924.565941040652, 5648.893486921194, 5369.545400080144, 5086.472672081212, 4799.625641042295, 4508.953982922859, 4214.406702695163, 3915.9321253977655, 3613.4778870697355, 3306.9909255639986, 2996.4174712381855, 2681.7030375213612, 2362.7924113549793, 1394.7511709901887, 0}}}

var monthlySequenceBalancesQATestExpected1 = []float64{
	38000,
	36050,
	35083.16666666667,
	34099.82529423611,
	33099.677671506666,
	32081.826182586745,
	31045.90808480681,
	29991.584253490277,
	28918.508880521018,
	27826.329337274263,
	26714.68603464671,
	25583.21228012319,
	24431.170205449595,
	23258.00595090137,
	22063.307779044157,
	20846.6555124799,
	19607.62035410622,
	18346.607694892617,
	17063.26397566094,
	15756.47322194304,
	14423.629323431236,
	13064.169739963108,
	11677.519065113369,
	10267.654585513452,
	8839.57288900163,
	7393.037766250776,
	5925.755014535265,
	4436.980643967897,
	2926.3991818982126,
	1394.7511709901887,
	0,
}