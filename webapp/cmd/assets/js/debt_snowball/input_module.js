import { generateCharts } from "/assets/js/debt_snowball/chart_module.js";
import { calculateAll } from "/assets/js/debt_snowball/calculations_module.js";

export async function recalculate(debtInputs, otherInputs, uuid) {
  let debts = [];
  let debt = {};
  let amounts = [];

  for (var i = 0; i < debtInputs.length; i++) {
    if (debtInputs[i].classList.contains("debt-name")) {
      debt.name = debtInputs[i].value;
    } else if (debtInputs[i].classList.contains("debt-amount")) {
      debt.amount = sanitizeToZero(debtInputs[i].value);
      amounts.push({"uuid": debtInputs[i].getAttribute("uuid"), "amount": debt.amount})
    } else if (debtInputs[i].classList.contains("debt-min-payment")) {
      debt.minimum_payment = sanitizeToZero(debtInputs[i].value);
    } else if (debtInputs[i].classList.contains("debt-interest")) {
      debt.interest = sanitizeToZero(debtInputs[i].value);
      debts.push(debt);
      debt = {};
    }
  }

  let calculations = await calculateAll(debts, sanitizeToZero(otherInputs[0].value), sanitizeToZero(otherInputs[1].value));

  document.getElementById("debt-payoff-month").innerHTML = calculations.DEBT_PAYOFF_MONTH;
  document.getElementById("total-beginning-debt").innerHTML = isDollarValue(calculations.TOTAL_BEGINNING_DEBT);
  document.getElementById("total-interest").innerHTML = isDollarValue(calculations.TOTAL_INTEREST);
  document.getElementById("total-minimum-payment").innerHTML = isDollarValue(calculations.TOTAL_MINIMUM_PAYMENT);
  document.getElementById("total-payments").innerHTML = isDollarValue(calculations.TOTAL_PAYMENTS);

  if (uuid != null) {
    validateInput(calculations.SNOWBALL, amounts);
  }

  generateCharts(calculations)
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

const isDollarValue = (s) => {
  return "$" + s.toLocaleString(undefined, {maximumFractionDigits: 2});
}

const validateInput = (snowball, amounts) => {
  var valid = new Array();
  amounts.forEach((x) => {
    if (x.amount > 0) {
      valid.push(x);
    } 
  })

  valid.sort((a, b) => a.amount - b.amount);

  for (var i = 0; i < valid.length; i++) {
    var inputAlert = document.getElementById(`input-validation-alert-${valid[i].uuid}`);
    if (snowball[i].invalid) {
      inputAlert.style.visibility = "visible";
    } else {
      inputAlert.style.visibility = "hidden";
    }
  }
}