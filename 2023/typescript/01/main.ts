let input = await Deno.readTextFile("01/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

const nums = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
};

function isDigit(char: string) {
  return /[1-9]/.test(char);
}

function findFirstIntIndex(line: string) {
  const res = line.split("").findIndex((char) => isDigit(char));
  return res == -1 ? null : res;
}
function findLastIntIndexFromRight(line: string) {
  const res = line
    .split("")
    .reverse()
    .findIndex((char) => isDigit(char));
  return res == -1 ? null : res;
}

function concatFirstLast(line: string) {
  let smallestIndex = Infinity;
  let smallestIndexFromRight = Infinity;
  let firstNum: string = "";
  let lastNum: string = "";
  const numKeys = Object.keys(nums) as (keyof typeof nums)[];
  numKeys.forEach((key) => {
    const res = line.replaceAll(key, nums[key]);
    const firstIntIndex = findFirstIntIndex(res);
    const lastIntIndexFromRight = findLastIntIndexFromRight(res);

    if (firstIntIndex != null && firstIntIndex < smallestIndex) {
      smallestIndex = firstIntIndex;
      firstNum = res[firstIntIndex];
    }
    if (
      lastIntIndexFromRight != null &&
      lastIntIndexFromRight < smallestIndexFromRight
    ) {
      smallestIndexFromRight = lastIntIndexFromRight;
      lastNum = res[res.length - 1 - lastIntIndexFromRight];
    }
  });
  if (firstNum && lastNum) return Number(firstNum + lastNum);
  throw Error("invalid line:" + line);
}

const lines = input.split("\n");
const res = lines.reduce((prev, line) => prev + concatFirstLast(line), 0);
console.log(res);
