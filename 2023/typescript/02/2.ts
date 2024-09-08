const input = await Deno.readTextFile("02/input.txt");
if (lines[lines.length - 1] == "") {
  lines.pop();
}

interface setData {
  red: number;
  blue: number;
  green: number;
}

function formatData(input: string) {
  let i = input.replaceAll(/^.*: /gm, "");
  const lines = i.split("\n");

  const result: setData[][] = [];

  lines.forEach((line) => {
    const gameData: setData[] = [];
    const sets = line.split("; ");

    sets.forEach((set) => {
      const setData: setData = {
        red: 0,
        blue: 0,
        green: 0,
      };
      const parts = set.split(", ");
      parts.forEach((part) => {
        const partParts = part.split(" ");
        const number = Number(partParts[0]);
        const color: keyof setData = partParts[1] as keyof setData;
        setData[color] = number;
      });

      gameData.push(setData);
    });

    result.push(gameData);
  });
  return result;
}

function findLeastPossible(gameData: setData[]) {
  const result: setData = {
    red: 0,
    blue: 0,
    green: 0,
  };
  gameData.forEach((set) => {
    result.red = set.red > result.red ? set.red : result.red;
    result.blue = set.blue > result.blue ? set.blue : result.blue;
    result.green = set.green > result.green ? set.green : result.green;
  });
  return result;
}

const formatted = formatData(input);

const sum = formatted.reduce((sum, gameData) => {
  const leastPossible = findLeastPossible(gameData);
  const power = leastPossible.red * leastPossible.green * leastPossible.blue;
  return sum + power;
}, 0);

console.log(sum);
