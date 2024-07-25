import { generateCharts } from "/assets/js/debt_snowball/chart_module.js";
import { calculateAll } from "/assets/js/debt_snowball/calculations_module.js";
import { generateOrderTable } from "/assets/js/debt_snowball/table_module.js"

export async function recalculate(debtInputs, otherInputs, uuid) {
  let debts = [];
  let debt = {};
  let uuids = [];

  for (var i = 0; i < debtInputs.length; i++) {
    if (debtInputs[i].classList.contains("debt-name")) {
      debt.name = debtInputs[i].value;
    } else if (debtInputs[i].classList.contains("debt-amount")) {
      debt.amount = sanitizeToZero(debtInputs[i].value);
    } else if (debtInputs[i].classList.contains("debt-min-payment")) {
      debt.minimum_payment = sanitizeToZero(debtInputs[i].value);
    } else if (debtInputs[i].classList.contains("debt-interest")) {
      debt.interest = sanitizeToZero(debtInputs[i].value);
      debt.id = debtInputs[i].getAttribute("uuid");
      uuids.push(debt.id);
      debts.push(debt);
      debt = {};
    }
  }

  let calculations = await calculateAll(debts, sanitizeToZero(otherInputs[0].value), sanitizeToZero(otherInputs[1].value));

  document.getElementById("debt-payoff-month-snowball").innerHTML = `${calculations.DEBT_PAYOFF_MONTH.snowball}(${(calculations.DEBT_PAYOFF_MONTH.snowball / 12).toFixed(1)}yr)`;
  document.getElementById("total-beginning-debt-snowball").innerHTML = isDollarValue(calculations.TOTAL_BEGINNING_DEBT.snowball);
  document.getElementById("total-interest-snowball").innerHTML = isDollarValue(calculations.TOTAL_INTEREST.snowball);
  document.getElementById("total-minimum-payment-snowball").innerHTML = isDollarValue(calculations.TOTAL_MINIMUM_PAYMENT.snowball);
  document.getElementById("total-payments-snowball").innerHTML = isDollarValue(calculations.TOTAL_PAYMENTS.snowball);

  document.getElementById("debt-payoff-month-avalanche").innerHTML = `${calculations.DEBT_PAYOFF_MONTH.avalanche}(${(calculations.DEBT_PAYOFF_MONTH.avalanche / 12).toFixed(1)}yr)`;
  document.getElementById("total-beginning-debt-avalanche").innerHTML = isDollarValue(calculations.TOTAL_BEGINNING_DEBT.avalanche);
  document.getElementById("total-interest-avalanche").innerHTML = isDollarValue(calculations.TOTAL_INTEREST.avalanche);
  document.getElementById("total-minimum-payment-avalanche").innerHTML = isDollarValue(calculations.TOTAL_MINIMUM_PAYMENT.avalanche);
  document.getElementById("total-payments-avalanche").innerHTML = isDollarValue(calculations.TOTAL_PAYMENTS.avalanche);

  if (uuid != null) {
    validateInput(calculations.VALID_DEBTS.snowball, uuids);
  }

  document.getElementById("decision").innerHTML = calculations.DECISION.snowball.choice;
  document.getElementById("decision").classList = calculations.DECISION.snowball.choice == "Snowball" ? ['snowball'] : ['avalanche']

  generateOrderTable(calculations);
  generateCharts(calculations);
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

const validateInput = (validDebts, uuids) => {
  let validDebtIds = validDebts.map(a => a.id);

  uuids.forEach((u) => {
    var inputAlert = document.getElementById(`input-validation-alert-${u}`); 
    if (validDebtIds.includes(u)) {
      inputAlert.style.visibility = "hidden";
    } else {
      inputAlert.style.visibility = "visible"; 
    }
  })
}
