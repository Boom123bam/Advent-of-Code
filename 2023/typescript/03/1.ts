const input = await Deno.readTextFile("03/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

interface NumData {
  number: number;
  row: number;
  start: number;
  end: number;
  hasAdjacent: boolean;
}

function isSymbol(char: string) {
  // return /[$&+,:;=?@#|'<>^*()%!/-]/.test(char);
  return !isDigit(char) && char != ".";
}
function isDigit(char: string) {
  return /[0-9]/.test(char);
}

function addAdjacents(row: number, col: number) {
  // replaces adjacents with true in map
  for (let offsetX = -1; offsetX <= 1; offsetX++) {
    for (let offsetY = -1; offsetY <= 1; offsetY++) {
      if (
        typeof adjacentMap[row + offsetY] === "undefined" ||
        typeof adjacentMap[row + offsetY][col + offsetX] === "undefined"
      )
        continue;
      adjacentMap[row + offsetY][col + offsetX] = true;
    }
  }
}

function printMap() {
  adjacentMap.forEach((row) => {
    let line = "";
    row.forEach((hasAdj) => (line += hasAdj ? "#" : "."));
    console.log(line);
  });
}

function updateAdjacentMap() {
  lines.forEach((line, rowIndex) => {
    line.split("").forEach((char, colIndex) => {
      if (isSymbol(char)) {
        addAdjacents(rowIndex, colIndex);
      }
    });
  });
}

function getNumsData() {
  const numDataList: NumData[] = [];
  lines.forEach((line, rowIndex) => {
    // gen list of numdata
    let start: number | null = null,
      end: number | null = null,
      hasAdjacent: boolean = false;

    line.split("").forEach((char, colIndex) => {
      if (isDigit(char) && start == null) {
        start = colIndex;
      } else if (start != null && !isDigit(char)) {
        end = colIndex;
        // add to list and reset
        numDataList.push({
          start,
          end,
          row: rowIndex,
          number: Number(line.substring(start, end)),
          hasAdjacent,
        });
        start = null;
        end = null;
        hasAdjacent = false;
      }

      if (adjacentMap[rowIndex][colIndex] && start != null) {
        // is adjacent
        hasAdjacent = true;
      }
    });

    if (start != null) {
      end = line.length;
      numDataList.push({
        start,
        end,
        row: rowIndex,
        number: Number(line.substring(start, end)),
        hasAdjacent,
      });
    }
  });
  return numDataList;
}

const lines = input.split("\n");

const adjacentMap = Array(lines.length)
  .fill(null)
  .map(() => Array(lines[0].length).fill(false));

// console.log(adjacentMap[0]);

updateAdjacentMap();
// printMap();
const numsData = getNumsData();

let sum = 0;
numsData.forEach((numData) => {
  if (numData.hasAdjacent) {
    sum += numData.number;
  }
});

console.log(sum);
