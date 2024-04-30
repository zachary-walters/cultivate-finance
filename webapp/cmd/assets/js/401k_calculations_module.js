import { wasmBrowserInstantiate } from "/assets/js/instantiateWasm.js";

const go = new Go(); 
const importObject = go.importObject;

export const calculate = async (input) => {
  const wasmModule = await wasmBrowserInstantiate("/assets/tinygoBin.wasm", importObject);
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
  const wasmModule = await wasmBrowserInstantiate("/assets/tinygoBin.wasm", importObject);
  go.run(wasmModule.instance); 

  return window.calculateAll({
    current_age: input.current_age,
    current_filing_status: input.current_filing_status,
    current_annual_income: input.current_annual_income,
    annual_contributions_pre_tax: input.annual_contributions_pre_tax,
    annual_investment_growth: input.annual_investment_growth,
    retirement_age: input.retirement_age,
    retirement_filing_status: input.retirement_filing_status,
    yearly_withdrawal: input.yearly_withdrawal,
  });
}

export const isDollarValue = (str) => {
  return "$" + str.toLocaleString();
}

export const isPercentValue = (str, convert) => {
  str = convert ? parseFloat(str).toFixed(4) * 100 : str;

  return str + "%";
}