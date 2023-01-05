
/**
  takes the params as a JSON object { "key": value }
*/
async function sendPOST(params) {
  var response;
  var json;
  try {
    response = await fetch('http://localhost:5000/api/calc', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(params)
    })
    if (response.ok) { // if HTTP-status is 200-299
      // get the response body (the method explained below)
      json = await response.json();
    } else {
      alert("HTTP-Error: " + response.status);
    }
    console.log(JSON.stringify(json));
  } catch (error) {
    alert(error)
  }

  return json.result;
}


const calculator = {
  displayValue: '0',
  firstOperand: null,
  waitingForSecondOperand: false,
  operator: null,
};

function inputDigit(digit) {
  const { displayValue, waitingForSecondOperand } = calculator;

  if (waitingForSecondOperand === true) {
    calculator.displayValue = digit;
    calculator.waitingForSecondOperand = false;
  } else {
    calculator.displayValue = displayValue === '0' ? digit : displayValue + digit;
  }
}

function inputDecimal(dot) {
  // If the `displayValue` does not contain a decimal point
  if (!calculator.displayValue.includes(dot)) {
    // Append the decimal point
    calculator.displayValue += dot;
  }
}

async function handleOperator(nextOperator) {
  const { firstOperand, displayValue, operator } = calculator
  const inputValue = parseFloat(displayValue);

  if (operator && calculator.waitingForSecondOperand) {
    calculator.operator = nextOperator;
    return;
  }

  if (firstOperand == null) {
    calculator.firstOperand = inputValue;
  } else if (operator) {
    const currentValue = firstOperand || 0;
    const result = await performCalculation[operator](currentValue, inputValue);

    calculator.displayValue = String(result);
    calculator.firstOperand = result;
  }

  calculator.waitingForSecondOperand = true;
  calculator.operator = nextOperator;
}

async function sendCalculation(first, second, operator) {
  const result = await sendPOST({
    "first": first,
    "second": second,
    "operator": operator,
  })
  return result;
}

const performCalculation = {
  '/': async (firstOperand, secondOperand) => await sendCalculation(firstOperand, secondOperand, '/'),

  '*': async (firstOperand, secondOperand) => await sendCalculation(firstOperand, secondOperand, 'x'),

  '+': async (firstOperand, secondOperand) => await sendCalculation(firstOperand, secondOperand, '+'),

  '-': async (firstOperand, secondOperand) => await sendCalculation(firstOperand, secondOperand, '-'),

  '=': async (firstOperand, secondOperand) => secondOperand
};

function resetCalculator() {
  calculator.displayValue = '0';
  calculator.firstOperand = null;
  calculator.waitingForSecondOperand = false;
  calculator.operator = null;
}

function updateDisplay() {
  const display = document.getElementById("calc-screen");
  if (calculator !== null && display !== null) {
    display.value = calculator.displayValue;
  }
}

updateDisplay();

window.addEventListener("load", function() {
  document.getElementById("calc-keys")?.addEventListener('click', (event) => {
    const { target } = event;
    if (!target.matches('button')) {
      return;
    }

    if (target.classList.contains('operator')) {
      handleOperator(target.value);
      updateDisplay();
      return;
    }

    if (target.classList.contains('decimal')) {
      inputDecimal(target.value);
      updateDisplay();
      return;
    }

    if (target.classList.contains('all-clear')) {
      resetCalculator();
      updateDisplay();
      return;
    }

    inputDigit(target.value);
    updateDisplay();
  });
});
