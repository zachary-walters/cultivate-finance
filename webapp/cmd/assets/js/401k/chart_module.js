import { isDollarValue } from "/assets/js/401k/401k_calculations_module.js";

export const generateCharts = (calculations) => {
  let chartStatus = Chart.getChart("traditional-chart"); 
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }

  new Chart("traditional-chart", {
    type: "doughnut",
    data: {
      labels: ["Contributions", "Interest"],
      datasets: [{
        backgroundColor: [
          "rgba(0, 40, 255, 0.5)",
          "rgba(85, 167, 131, 0.5)",
        ],
        data: [calculations.traditional.TOTAL_CONTRIBUTIONS, calculations.traditional_retirement.TOTAL_INTEREST]
      }]
    },
    options: {
      animation: false,
      plugins: {
        title: {
          display: true,
          text: "Traditional" + " ( " + isDollarValue(calculations.traditional_retirement.TOTAL_DISBURSEMENTS) + " )",
        }
      }
    }
  });

  chartStatus = Chart.getChart("roth-chart"); 
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }
  new Chart("roth-chart", {
    type: "doughnut",
    data: {
      labels: ["Contributions", "Interest"],
      datasets: [{
        backgroundColor: [
          "rgba(0, 40, 255, 0.5)",
          "rgba(85, 167, 131, 0.5)",
        ],
        data: [calculations.roth.TOTAL_CONTRIBUTIONS, calculations.roth_retirement.TOTAL_INTEREST]
      }]
    },
    options: {
      animation: false,
      plugins: {
        title: {
          display: true,
          text: "Roth" + " ( " + isDollarValue(calculations.roth_retirement.TOTAL_DISBURSEMENTS) + " )",
        }
      }
    }
  });

  chartStatus = Chart.getChart("area-chart"); 
  if (chartStatus != undefined) {
    chartStatus.destroy();
  }
  new Chart('area-chart', { 
    type: 'line', 
    data: { 
      labels: Object.keys(calculations.traditional.BALANCES_TRADITIONAL.ending_balance), 
      datasets: [ 
        { 
          label: 'Traditional', 
          data: Object.values(calculations.traditional.BALANCES_TRADITIONAL.ending_balance), 
          backgroundColor: 'rgba(85, 167, 131, 0.5)', 
          borderColor: 'rgba(0, 255, 0, 0)', 
          borderWidth: 1, 
          fill: true, 
          pointStyle: false,
        }, 
        { 
          label: 'Roth', 
          data: Object.values(calculations.traditional.BALANCES_ROTH_MATCHING_NET_CONTRIBUTIONS.ending_balance), 
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
            text: 'Age' 
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
          position: 'top', 
        }, 
      } 
    } 
  }); 
}