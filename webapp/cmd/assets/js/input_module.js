export const get401kCalculatorInput = (datakey) => {
  return {
    datakey: datakey,
    current_age: parseInt(document.getElementById('current-age').value),
    current_filing_status: document.getElementById('current-filing-status').value,
    current_annual_income: parseFloat(document.getElementById('current-annual-income').value),
    annual_contributions_pre_tax: parseFloat(document.getElementById('annual-contributions-pre-tax').value),
    annual_investment_growth: parseFloat(document.getElementById('annual-investment-growth').value),
    retirement_age: parseInt(document.getElementById('retirement-age').value),
    retirement_filing_status: document.getElementById('retirement-filing-status').value,
    yearly_withdrawal: parseFloat(document.getElementById('yearly-withdrawal').value),
  }
};
