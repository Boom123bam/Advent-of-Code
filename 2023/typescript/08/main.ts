let input = await Deno.readTextFile("08/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

let [directions, rest] = input.split("\n\n");

const instructionsMap = new Map(
  rest.split("\n").map((line) => {
    const [label, left, right] = line.match(/[A-Z1-9]+/g) as string[];
    // instructions[label] = { left, right };
    return [label, [left, right]];
  }),
);

function getDirection(index: number) {
  return directions[index % directions.length];
}

function getSteps(from: string, to: RegExp) {
  let current = from;
  let step = 0;
  for (; !current.match(to); step++) {
    const direction = getDirection(step);
    current = instructionsMap.get(current)?.[direction == "L" ? 0 : 1] ?? "";
  }

  return step;
}

console.log(getSteps("AAA", /ZZZ/g));

// part2

function gcd(a: number, b: number): number {
  return b === 0 ? a : gcd(b, a % b);
}

function lcm(a: number, b: number) {
  return (a * b) / gcd(a, b);
}

const keysEndingWithA = Array.from(instructionsMap.keys()).filter(
  (key) => key[2] == "A",
);

const stepsForEachKey = keysEndingWithA.map((key) => getSteps(key, /..Z/g));

console.log(stepsForEachKey.reduce(lcm));
