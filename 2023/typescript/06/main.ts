let input = await Deno.readTextFile("06/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

const [times, distances] = input
  .split("\n")
  .map((line) => line.split(/: */)[1].split(/ +/g).map(Number));

function getNumWaysToWin(t: number, recordDist: number) {
  let ways = 0;
  for (let holdTime = 0; holdTime <= t; holdTime++) {
    const remainingTime = t - holdTime;
    const dist = holdTime * remainingTime;
    if (dist > recordDist) ways++;
  }
  return ways;
}

const p1 = times.reduce((acc, time, idx) => {
  const distance = distances[idx];
  return acc * getNumWaysToWin(time, distance);
}, 1);

console.log(p1);

const time = Number(times.reduce((res, num) => res + num.toString(), ""));
const distance = Number(
  distances.reduce((res, num) => res + num.toString(), ""),
);

const p2 = getNumWaysToWin(time, distance);
console.log(p2);
