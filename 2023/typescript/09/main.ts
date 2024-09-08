let input = await Deno.readTextFile("09/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

const lines = input
  .split("\n")
  .map((lineString) => lineString.split(" ").map(Number));

function extrapolate(line: number[], sum = line[line.length - 1]) {
  const newLine = line.reduce((result: number[], currentNum, index) => {
    if (index == 0) return result;
    const prevNum = line[index - 1];
    return [...result, currentNum - prevNum];
  }, []);

  // if (!newLine.filter((num) => num != 0).length) return sum;
  if (newLine.every((num) => num == 0)) return sum;

  sum += newLine[newLine.length - 1];
  return extrapolate(newLine, sum);
}

function extrapolate2(line: number[], result = line[0]) {
  let currentLine = line;
  let positive = false;

  function execute() {
    currentLine = currentLine.reduce((result: number[], currentNum, index) => {
      if (index == 0) return result;
      const prevNum = currentLine[index - 1];
      return [...result, currentNum - prevNum];
    }, []);
    if (currentLine.every((num) => num == 0)) return result;

    result += currentLine[0] * (positive ? 1 : -1);
    positive = !positive;
    execute();
  }
  execute();
  return result;
}

const extrapolatedVals = lines.map((l) => extrapolate(l));
const result = extrapolatedVals.reduce((a, b) => a + b);
console.log(result);
const extrapolatedVals2 = lines.map((l) => extrapolate2(l));
const result2 = extrapolatedVals2.reduce((a, b) => a + b);
console.log(result2);
