import { wasmBrowserInstantiate } from "/assets/js/instantiateWasm.js";

const go = new Go(); 
const importObject = go.importObject;

export const calculateAll = async (debts, oneTimePayment, extraMonthlyPayment) => {
  const wasmModule = await wasmBrowserInstantiate("/assets/debtsnowball.wasm", importObject);
  go.run(wasmModule.instance); 

  return window.calculateAll(debts, debts.length, oneTimePayment, extraMonthlyPayment);
}

