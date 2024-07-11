export const generateCharts = (calculations) => {
  let chartStatus = Chart.getChart("months-payoff-chart");
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }

  var snowball = calculations.SNOWBALL.filter(debt => !debt.invalid);

  new Chart("months-payoff-chart", {
    type: "bar",
    data: {
      labels: getLabels(snowball),
      datasets: [{
        data: getMonthPayoffData(snowball),
        borderWidth: 2,
        backgroundColor: "#8091F3",
      }],
    },
    options: {
      indexAxis: 'y',
      animation: false,
      bar: {
        borderWidth: 2,
      },
      responsive: true,
      plugins: {
        legend: {
          display: false,
        },
        title: {
          display: true,
          text: "Months Until Debt Payoff"
        },
      },
    },
  });

  chartStatus = Chart.getChart("payoff-over-time-chart"); 
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }
  new Chart('payoff-over-time-chart', { 
    type: 'line', 
    data: { 
      labels: calculations.MONTH_SEQUENCE, 
      datasets: [ 
        {  
          data: calculations.MONTHLY_SEQUENCE_BALANCES, 
          backgroundColor: 'rgba(0, 40, 255, 0.5)', 
          borderColor: 'rgba(0, 0, 255, 0)', 
          borderWidth: 1, 
          fill: true, 
          pointStyle: false,
        }, 
      ] 
    }, 
    options: { 
      animation: false,
      scales: { 
        y: { 
          beginAtZero: true, 
          stacked: false, 
          title: { 
            display: true, 
            text: 'Dollars' 
          } 
        }, 
        x: { 
          stacked: true,
          title: { 
            display: true, 
            text: 'Months' 
          } 
        } 
      }, 
      layout: { 
        padding: { 
          left: 20, 
          right: 20, 
          top: 20, 
          bottom: 20 
        } 
      }, 
      plugins: { 
        legend: {
          display: false,
        },
        title: {
          display: true,
          text: "Payoff Over Time"
        },
      } 
    } 
  }); 

  chartStatus = Chart.getChart("snowball-donut-chart"); 
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }

  new Chart("snowball-donut-chart", {
    type: "doughnut",
    data: {
      labels: ["Interest", "Initial Debt"],
      datasets: [{
        backgroundColor: [
          "rgba(0, 40, 255, 0.5)",
          "rgba(85, 167, 131, 0.5)",
        ],
        data: [calculations.TOTAL_INTEREST,calculations.TOTAL_BEGINNING_DEBT]
      }]
    },
    options: {
      animation: false,
      plugins: {
        title: {
          display: true,
          text: "Total Principal and Interest",
        }
      }
    }
  });
}

const getLabels = (snowball) => {
  let debtNames = [];
  for (var i = 0; i < snowball.length; i++) {
    debtNames.push(snowball[i].debt.name);
  }
  return debtNames
}

const getMonthPayoffData = (snowball) => {
  let maxMonths = [];
  for (var i = 0; i < snowball.length; i++) {
    let debt = snowball[i];
    maxMonths.push(debt.months[debt.months.length - 1]);
  }
  return maxMonths
}