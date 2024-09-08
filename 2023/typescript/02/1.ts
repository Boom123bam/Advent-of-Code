const input = await Deno.readTextFile("02/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
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
      // console.log(set);
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

function isImpossible(gameData: setData[]) {
  let impossible = false;
  gameData.forEach(
    (set) =>
      (impossible =
        set.red > 12 || set.green > 13 || set.blue > 14 ? true : impossible),
  );
  return impossible;
}

const formatted = formatData(input);

const sum = formatted.reduce((sum, gameData, index) => {
  return isImpossible(gameData) ? sum : sum + index + 1;
}, 0);

console.log(sum);
