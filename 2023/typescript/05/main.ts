let input = await Deno.readTextFile("05/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`;

function remap(line: number[], num: number) {
  const [to, from, range] = line;

  if (num >= from && num < from + range) {
    return to + num - from;
  }
  return null;
}

function getLocation(seedNum: number) {
  maps.forEach((map) => {
    for (const line of map) {
      const remapped = remap(line, seedNum);
      if (remapped) {
        seedNum = remapped;
        break;
      }
    }
  });
  return seedNum;
}

function findMin(seeds: number[]) {
  const result = seeds?.map(getLocation);

  return result?.reduce((min, num) => (num < min ? num : min), Infinity);
}

let sections = testInput.split("\n\n");

const firstLine = sections.shift();

const seeds = firstLine?.split(": ")[1].split(" ").map(Number);

let maps = sections.map((section) =>
  section
    .split("map:\n")[1]
    .split("\n")
    .map((line) => line.split(" ").map(Number)),
);

if (seeds) console.log(findMin(seeds));

const pairs = firstLine
  ?.split(":")[1]
  .match(/\d+ \d+/g)
  ?.map((pair) => {
    const [start, range] = pair.split(" ").map(Number);
    return { start, end: start + range };
  });

let p2Min = Infinity;
pairs?.forEach((pair) => {
  const { start, end } = pair;
  for (let i = start; i < end; i++) {
    const res = getLocation(i);
    if (res < p2Min) p2Min = res;
  }
});

console.log(p2Min);
