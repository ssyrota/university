<<<<<<< Updated upstream
const inputMatrix = document.getElementById("input-matrix");
const outputMatrix = document.getElementById("output-matrix");
const determinantOutput = document.getElementById("determinant-output");
const errorMessage = document.getElementById("error-message");
const addRowBtn = document.getElementById("add-row");
const removeRowBtn = document.getElementById("remove-row");
const addColBtn = document.getElementById("add-column");
const removeColBtn = document.getElementById("remove-column");

let debounceTimer;

function debounce(func, delay) {
  return function (...args) {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => func.apply(this, args), delay);
  };
=======
import { api } from "./api.js";

const errorMessage = document.getElementById("error-message");

class Matrix {
  constructor(id) {
    this.matrix = document.getElementById(id);
  }
}


class InputMatrix {
  constructor(outputMatrix) {
    this.inputMatrix = new Matrix("input-matrix");
    this.outputMatrix = outputMatrix;
    this.initObservers();
  }

  initObservers(){
    const addRowBtn = document.getElementById("add-row");
    const removeRowBtn = document.getElementById("remove-row");
    const addColBtn = document.getElementById("add-column");
    const removeColBtn = document.getElementById("remove-column");
    addRowBtn.addEventListener('click', this.onAddRow);
    removeRowBtn.addEventListener('click', this.onRemoveRow);
    addColBtn.addEventListener('click', this.onAddColumn);
    removeColBtn.addEventListener('click', this.onRemoveColumn);

  }

  onAddRow() {}
  onRemoveRow() {}
  onAddColumn() {}
  onRemoveColumn() {}
}

class OutputMatrix {
  constructor() {
    this.outputMatrix = new Matrix("output-matrix");
  }
>>>>>>> Stashed changes
}
const outputMatrix = new OutputMatrix()
const inputMatrix = new InputMatrix(outputMatrix)


function getInputData() {
  const rows = inputMatrix.querySelectorAll("tbody tr");
  const a = [];
  const b = [];

  rows.forEach((row) => {
    const aRow = [];
    const aInputs = row.querySelectorAll(".a");
    aInputs.forEach((input) => {
      aRow.push(parseFloat(input.value) || 0);
    });
    a.push(aRow);
    const bVal = parseFloat(row.querySelector(".b").value) || 0;
    b.push(bVal);
  });

  return { a, b };
}

function updateOutput(data) {
  if (data.error) {
    errorMessage.textContent = data.error;
    outputMatrix
      .querySelectorAll("input")
      .forEach((input) => (input.value = ""));
  } else {
    errorMessage.textContent = "";
    data.b.forEach((value, index) => {
      const xInput = document.getElementById(`x${index + 1}`);
      if (xInput) xInput.value = value;
    });
  }
}

<<<<<<< Updated upstream
function updateDeterminant(data) {
  if (data.error) {
    errorMessage.textContent = data.error;
    determinantOutput.value = "";
  } else {
    errorMessage.textContent = "";
    console.log("updateDeterminant", { data, determinantOutput });
    determinantOutput.textContent = data.determinant;
  }
}

const solve = debounce(() => {
  const inputData = getInputData();
  fetch("/solve", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(inputData),
  })
    .then((response) => response.json())
    .then((data) => updateOutput(data))
    .catch((error) => {
      errorMessage.textContent = "An error occurred while solving the system.";
      console.error(error);
    });
}, 100);

const computeDeterminant = debounce(() => {
  const inputData = getInputData();
  fetch("/det", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(inputData),
  })
    .then((response) => response.json())
    .then((data) => updateDeterminant(data))
    .catch((error) => {
      errorMessage.textContent =
        "An error occurred while computing the determinant.";
      console.error(error);
    });
}, 100);

inputMatrix.addEventListener("input", () => {
  solve();
  computeDeterminant();
});
=======
inputMatrix.addEventListener('input', solve);
>>>>>>> Stashed changes

function addRow() {
  const tbody = inputMatrix.querySelector("tbody");
  const currentRowCount = tbody.querySelectorAll("tr").length;
  const newRowNumber = currentRowCount + 1;

  const tr = document.createElement("tr");

  const eqTd = document.createElement("td");
  eqTd.textContent = newRowNumber;
  tr.appendChild(eqTd);

  const colCount = inputMatrix.querySelectorAll("thead th").length - 2; // Exclude 'Equation' and '= b'
  for (let i = 1; i <= colCount; i++) {
    const td = document.createElement("td");
    const input = document.createElement("input");
    input.type = "number";
    input.classList.add("a");
    input.dataset.row = newRowNumber;
    input.dataset.col = i;
    td.appendChild(input);
    tr.appendChild(td);
  }

  const bTd = document.createElement("td");
  const bInput = document.createElement("input");
  bInput.type = "number";
  bInput.classList.add("b");
  bInput.dataset.row = newRowNumber;
  bTd.appendChild(bInput);
  tr.appendChild(bTd);

  tbody.appendChild(tr);
}

function removeRow() {
  const tbody = inputMatrix.querySelector("tbody");
  const rows = tbody.querySelectorAll("tr");
  if (rows.length > 1) {
    tbody.removeChild(rows[rows.length - 1]);
  }
}

function addColumn() {
  const thead = inputMatrix.querySelector("thead tr");
  const currentColCount = thead.querySelectorAll("th").length - 2; // Exclude 'Equation' and '= b'
  const newColNumber = currentColCount + 1;

  const th = document.createElement("th");
  th.textContent = `x${newColNumber}`;
  thead.insertBefore(th, thead.lastElementChild);

  const tbody = inputMatrix.querySelector("tbody");
  tbody.querySelectorAll("tr").forEach((tr, index) => {
    const td = document.createElement("td");
    const input = document.createElement("input");
    input.type = "number";
    input.classList.add("a");
    input.dataset.row = index + 1;
    input.dataset.col = newColNumber;
    td.appendChild(input);
    tr.insertBefore(td, tr.lastElementChild);
  });

  const outputThead = outputMatrix.querySelector("thead tr");
  const outputTh = document.createElement("th");
  outputTh.textContent = `x${newColNumber}`;
  outputThead.appendChild(outputTh);

  const outputTbody = outputMatrix.querySelector("tbody tr");
  const outputTd = document.createElement("td");
  const outputInput = document.createElement("input");
  outputInput.type = "number";
  outputInput.id = `x${newColNumber}`;
  outputInput.readOnly = true;
  outputTd.appendChild(outputInput);
  outputTbody.appendChild(outputTd);
}

function removeColumn() {
  const thead = inputMatrix.querySelector("thead tr");
  const headerCols = thead.querySelectorAll("th");
  if (headerCols.length > 3) {
    // Ensure at least one variable column remains
    thead.removeChild(headerCols[headerCols.length - 2]); // Remove the last variable column

    const tbody = inputMatrix.querySelector("tbody");
    tbody.querySelectorAll("tr").forEach((tr) => {
      tr.removeChild(tr.children[tr.children.length - 2]);
    });

    const outputThead = outputMatrix.querySelector("thead tr");
    const outputHeaderCols = outputThead.querySelectorAll("th");
    outputThead.removeChild(outputHeaderCols[outputHeaderCols.length - 1]); // Remove last th

    const outputTbody = outputMatrix.querySelector("tbody tr");
    outputTbody.removeChild(
      outputTbody.children[outputTbody.children.length - 1]
    ); // Remove last td
  }
}
<<<<<<< Updated upstream

addRowBtn.addEventListener("click", addRow);
removeRowBtn.addEventListener("click", removeRow);
addColBtn.addEventListener("click", addColumn);
removeColBtn.addEventListener("click", removeColumn);
=======
>>>>>>> Stashed changes
