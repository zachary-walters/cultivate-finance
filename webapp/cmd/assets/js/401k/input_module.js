import { calculateAll, isPercentValue, isDollarValue } from "/assets/js/401k/401k_calculations_module.js";
import { generateCharts } from "/assets/js/401k/chart_module.js";

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

export async function recalculate() {
  let all = await calculateAll(get401kCalculatorInput());
  document.getElementById("annual-growth-less-inflation").innerHTML = isPercentValue(all.traditional.ANNUAL_GROWTH_LESS_INFLATION);
  document.getElementById("annual-tax-savings-with-contributions").innerHTML = isDollarValue(all.traditional.ANNUAL_TAX_SAVINGS_WITH_CONTRIBUTION);
  document.getElementById("effective-tax-rate-on-gross").innerHTML = isPercentValue(all.traditional_retirement.EFFECTIVE_TAX_RATE_ON_GROSS, true);
  document.getElementById("equivalent-roth-contributions").innerHTML = isDollarValue(all.traditional.EQUIVALENT_ROTH_CONTRIBUTIONS);
  document.getElementById("income-after-standard-deduction").innerHTML = isDollarValue(all.traditional.INCOME_AFTER_STANDARD_DEDUCTION);
  document.getElementById("income-after-standard-deduction-and-contributions").innerHTML = isDollarValue(all.traditional.INCOME_AFTER_STANDARD_DEDUCTION_AND_CONTRIBUTIONS);
  document.getElementById("income-after-standard-deduction-retirement").innerHTML = isDollarValue(all.traditional_retirement.INCOME_AFTER_STANDARD_DEDUCTION);
  document.getElementById("net-distribution-after-taxes").innerHTML = isDollarValue(all.traditional_retirement.NET_DISTRIBUTION_AFTER_TAXES);
  document.getElementById("standard-deduction").innerHTML = isDollarValue(all.traditional.STANDARD_DEDUCTION);
  document.getElementById("standard-deduction-retirement").innerHTML = isDollarValue(all.traditional_retirement.STANDARD_DEDUCTION);
  document.getElementById("tax-rate-of-savings").innerHTML = isPercentValue(all.traditional.TAX_RATE_OF_SAVINGS, true);
  document.getElementById("total-taxes-owed-after-standard-deduction").innerHTML = isDollarValue(all.traditional_retirement.TOTAL_TAXES_OWED_AFTER_STANDARD_DEDUCTION);
  document.getElementById("decision").innerHTML = all.traditional.ROTH_OR_TRADITIONAL_DECISION;
  
  document.getElementById("decision").classList = all.traditional.ROTH_OR_TRADITIONAL_DECISION == "Traditional" ? ['traditional'] : ['roth']

  generateCharts(all);
}

