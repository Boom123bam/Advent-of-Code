let input = await Deno.readTextFile("10/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

class Pipe {
  char: string;
  coord: Coord;

  neighbours: { connected: boolean; coord: Coord | null }[] = [];

  constructor(coord: Coord) {
    this.char = lines[coord.y][coord.x];
    this.coord = coord;

    const [topCoord, rightCoord, bottomCoord, leftCoord] =
      getNeighborsCoords(coord);
    const [topM, rightM, bottomM, leftM] = getPipeMappings(this.char);
    this.neighbours = [
      {
        coord: topCoord,
        connected: topM && !(topCoord == null),
      },
      {
        coord: rightCoord,
        connected: rightM && !(rightCoord == null),
      },
      {
        coord: bottomCoord,
        connected: bottomM && !(bottomCoord == null),
      },
      {
        coord: leftCoord,
        connected: leftM && !(leftCoord == null),
      },
    ];
  }

  getNextPipe(directionToReject: number) {
    for (const [directionIndex, neighbour] of this.neighbours.entries()) {
      if (
        !neighbour.connected ||
        directionIndex == directionToReject ||
        neighbour.coord == null
      )
        continue;
      return [new Pipe(neighbour.coord), directionIndex];
    }
    return null;
  }
}

const pipeMappings = {
  "═": [0, 1, 0, 1],
  "║": [1, 0, 1, 0],
  "╗": [0, 0, 1, 1],
  "╝": [1, 0, 0, 1],
  "╚": [1, 1, 0, 0],
  "╔": [0, 1, 1, 0],
  ".": [0, 0, 0, 0],
  S: [1, 1, 1, 1],
};

function getPipeMappings(char: string) {
  return pipeMappings[char as keyof typeof pipeMappings].map(Boolean);
}

function getNeighborsCoords(coord: Coord) {
  const topRightBottomLeft: Coord[] = [
    { x: 0, y: -1 },
    { x: 1, y: 0 },
    { x: 0, y: 1 },
    { x: -1, y: 0 },
  ];

  return topRightBottomLeft.map((shift) => {
    const newCoord: Coord = {
      x: coord.x + shift.x,
      y: coord.y + shift.y,
    };
    return coordIsVaild(newCoord) ? newCoord : null;
  });
}

function coordIsVaild(coord: Coord) {
  return (
    typeof lines[coord.y] !== "undefined" &&
    typeof lines[coord.y][coord.x] !== "undefined"
  );
}

function getStartPipeShape() {
  const startY = lines
    .map((line) => (line.indexOf("S") == -1 ? false : true))
    .indexOf(true);
  const startX = lines[startY].indexOf("S");
  const start = new Pipe({ x: startX, y: startY });

  const pipeMapping = [0, 0, 0, 0];

  for (const [directionIndex, neighbour] of start.neighbours.entries()) {
    if (!neighbour.coord) continue;
    const neighbourPipe = new Pipe(neighbour.coord);

    if (neighbourPipe.neighbours[(directionIndex + 2) % 4].connected)
      pipeMapping[directionIndex] = 1;
  }

  return Object.keys(pipeMappings).find(
    (key) =>
      pipeMappings[key as keyof typeof pipeMappings].toString() ===
      pipeMapping.toString(),
  );
}

function findPipeConnectedToStart(): [Pipe, number] | [null, null] {
  const startY = lines
    .map((line) => (line.indexOf("S") == -1 ? false : true))
    .indexOf(true);
  const startX = lines[startY].indexOf("S");
  const start = new Pipe({ x: startX, y: startY });

  for (const [directionIndex, neighbour] of start.neighbours.entries()) {
    if (!neighbour.coord) continue;
    const neighbourPipe = new Pipe(neighbour.coord);

    if (neighbourPipe.neighbours[(directionIndex + 2) % 4].connected)
      return [neighbourPipe, directionIndex];
  }
  return [null, null];
}

const lines = input.split("\n");

interface Coord {
  x: number;
  y: number;
}

let [currentPipe, prevDirectionIndex] = findPipeConnectedToStart();

const charsInLoopCoords: String[] = [];

if (prevDirectionIndex != null) {
  while (currentPipe != null && currentPipe.char != "S") {
    if (prevDirectionIndex == null) break;

    charsInLoopCoords.push(JSON.stringify(currentPipe.coord));

    [currentPipe, prevDirectionIndex] = currentPipe?.getNextPipe(
      (prevDirectionIndex + 2) % 4,
    );
  }

  if (currentPipe) charsInLoopCoords.push(JSON.stringify(currentPipe.coord));
}

console.log(Math.ceil(charsInLoopCoords.length / 2));

let insideCount = 0;
lines.forEach((line, y) => {
  let inside = false;
  let prevWasUp = false;
  line.split("").forEach((char, x) => {
    if (char == "S") char = getStartPipeShape() ?? "";
    const coord: Coord = { x, y };
    if (!charsInLoopCoords.includes(JSON.stringify(coord))) {
      if (inside) insideCount++;
      return;
    }

    const [up, right, bottom, left] = getPipeMappings(char);
    if (left && right) return;

    if (up && bottom) inside = !inside;
    else if (right)
      prevWasUp = up; // start horizontal pipe
    else if (left && prevWasUp != up) inside = !inside; // end horizontal pipe and switch direction
  });
});

console.log(insideCount);
