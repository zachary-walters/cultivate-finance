import { calculateAll } from "/assets/js/debt_snowball/calculations_module.js";

export async function recalculate(debtInputs, otherInputs) {
  let debts = [];
  let debt = {};

  for (var i = 0; i < debtInputs.length; i++) {
    if (debtInputs[i].classList.contains("debt-name")) {
      debt.name = debtInputs[i].value;
    } else if (debtInputs[i].classList.contains("debt-amount")) {
      debt.amount = sanitizeToZero(debtInputs[i].value);
    } else if (debtInputs[i].classList.contains("debt-min-payment")) {
      debt.minimum_payment = sanitizeToZero(debtInputs[i].value);
    } else if (debtInputs[i].classList.contains("debt-interest")) {
      debt.interest = sanitizeToZero(debtInputs[i].value);
      debts.push(debt);
      debt = {};
    }
  }

  let all = await calculateAll(debts, sanitizeToZero(otherInputs[0].value), sanitizeToZero(otherInputs[1].value));
  
  console.log(all)
}

export const sanitizeToZero = (input) => {
  if (input === undefined || input == null || input === NaN || input === '') {
    return parseFloat(0);
  }

  if (input < 0) {
    return parseFloat(0);
  }

  return parseFloat(input)
}