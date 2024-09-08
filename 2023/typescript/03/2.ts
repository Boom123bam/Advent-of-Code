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

function getNum(row: number, col: number) {
  let offSet = 0;
  let string = "";
  while (
    typeof lines[row][col + offSet - 1] !== "undefined" &&
    isDigit(lines[row][col + offSet - 1])
  ) {
    offSet -= 1;
  }
  while (
    typeof lines[row][col + offSet] !== "undefined" &&
    isDigit(lines[row][col + offSet])
  ) {
    string += lines[row][col + offSet];
    offSet += 1;
  }
  return Number(string);
}

function findAdjacents(row: number, col: number) {
  let prevWasDigit = false;
  const adjacents = [];
  for (let offsetY = -1; offsetY <= 1; offsetY++) {
    for (let offsetX = -1; offsetX <= 1; offsetX++) {
      // check if it exists
      if (
        typeof lines[row + offsetY] === "undefined" ||
        typeof lines[row + offsetY][col + offsetX] === "undefined"
      )
        continue;
      const char = lines[row + offsetY][col + offsetX];
      // console.log(char);
      if (!prevWasDigit && isDigit(char)) {
        adjacents.push([row + offsetY, col + offsetX]);
      }
      // console.log(char, prevWasDigit);

      prevWasDigit = isDigit(char);
    }
    prevWasDigit = false;
  }
  // console.log(" ");

  return adjacents.map((coord) => getNum(coord[0], coord[1]));
}

function findGearRatio(row: number, col: number) {
  const adjs = findAdjacents(row, col);
  // console.log(adjs);
  return adjs.length == 2 ? adjs[0] * adjs[1] : 0;
}

// console.log(adjs);

// const adjacentMap = Array(lines.length)
//   .fill(null)
//   .map(() => Array(lines[0].length).fill(false));

// // console.log(adjacentMap[0]);

// updateAdjacentMap();
// // printMap();
// const numsData = getNumsData();

// let sum = 0;
// numsData.forEach((numData) => {
//   if (numData.hasAdjacent) {
//     sum += numData.number;
//   }
// });

// console.log(sum);

// const lines = testInput.split("\n");
const lines = input.split("\n");

let sum = 0;

lines.forEach((line, row) => {
  line.split("").forEach((char, col) => {
    sum += char == "*" ? findGearRatio(row, col) : 0;
  });
});

console.log(sum);
