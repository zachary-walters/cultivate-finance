import { wasmBrowserInstantiate } from "/assets/js/instantiateWasm.js";

const go = new Go(); 
const importObject = go.importObject;

export const calculate = async (input) => {
  const wasmModule = await wasmBrowserInstantiate("/assets/401k.wasm", importObject);
  go.run(wasmModule.instance);

  return window.calculate({
    datakey: input.datakey,
    current_age: input.current_age,
    current_filing_status: input.CurrentFilingStatus,
    current_annual_income: input.current_annual_income,
    annual_contributions_pre_tax: input.annual_contributions_pre_tax,
    annual_investment_growth: input.annual_investment_growth,
    retirement_age: input.retirement_age,
    retirement_filing_status: input.retirement_filing_status,
    yearly_withdrawal: input.yearly_withdrawal,
  });

};

export const calculateAll = async (input) => {
  const wasmModule = await wasmBrowserInstantiate("/assets/401k.wasm", importObject);
  go.run(wasmModule.instance); 

  return window.calculateAll({
    current_age: input.current_age,
    current_filing_status: input.current_filing_status,
    current_annual_income: input.current_annual_income,
    annual_contributions_pre_tax: input.annual_contributions_pre_tax,
    annual_investment_growth: input.annual_investment_growth,
    retirement_age: input.retirement_age,
    retirement_filing_status: input.retirement_filing_status,
    work_income: input.work_income,
    qualified_dividends: input.qualified_dividends,
    other_long_term_capital_gains: input.other_long_term_capital_gains,
    pension_income: input.pension_income,
    rental_net_income: input.rental_net_income,
    annuity_income: input.annuity_income,
    social_security: input.social_security,
    other_taxable_income: input.other_taxable_income,
    yearly_withdrawal: input.yearly_withdrawal,
  });
}

export const isDollarValue = (s) => {
  return "$" + s.toLocaleString(undefined, {maximumFractionDigits: 2});
}

export const isPercentValue = (s, convert) => {
  s = convert ? parseFloat(s).toFixed(4) * 100 : s;

  return s + "%";
}