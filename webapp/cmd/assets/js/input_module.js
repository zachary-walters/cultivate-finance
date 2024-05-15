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
    work_income: parseFloat(document.getElementById('work-income').value),
    qualified_dividends: parseFloat(document.getElementById('qualified-dividends').value),
    other_long_term_capital_gains: parseFloat(document.getElementById('other-long-term-capital-gains').value),
    pension_income: parseFloat(document.getElementById('pension-income').value),
    rental_net_income: parseFloat(document.getElementById('rental-net-income').value),
    annuity_income: parseFloat(document.getElementById('annuity-income').value),
    social_security: parseFloat(document.getElementById('social-security').value),
    other_taxable_income: parseFloat(document.getElementById('other-taxable-income').value),
    yearly_withdrawal: parseFloat(document.getElementById('yearly-withdrawal').value),
  }
};
