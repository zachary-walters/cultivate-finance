export const generateOrderTable = (calculations) => {
  let table = `<table class="fixed_headers">
                <thead>
                  <tr>
                    <th class="snowball-header">Snowball</th>
                    <th class="avalanche-header">Avalanche</th>
                  </tr>
                </thead>
                <tbody>`;

  for (var i = 0; i < calculations.VALID_SNOWBALL.snowball.length; i++) {
    let snowball = calculations.VALID_SNOWBALL.snowball[i];
    let avalanche = calculations.VALID_SNOWBALL.avalanche[i];
    table = table + `<tr>
                       <td>${snowball.debt.name} (mo: ${snowball.months.slice(-1)[0]})</td>
                      <td>${avalanche.debt.name} (mo: ${avalanche.months.slice(-1)[0]})</td>
                    </tr>`;
  }

  table = table + `</tbody>
                </table>`;

  document.getElementById("debt-table").innerHTML = table;
}