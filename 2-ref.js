const testInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`;

const lines = testInput.split("\n");

const games = lines.reduce((games, line) => {
  const [id, reaches] = line.split(": ");
  const gameId = Number(id.split(" ")[1]);
  const game = { gameId, red: 0, green: 0, blue: 0 };
  reaches.split("; ").forEach((reach) => {
    reach.split(", ").forEach((color) => {
      const [amount, name] = color.split(" ");
      if (game[name] < Number(amount)) game[name] = Number(amount);
    });
  });
  games.push(game);
  return games;
}, []);
console.log(games);

//part1
const part1 = games.reduce(
  (sum, game) =>
    (sum +=
      game.red <= 12 && game.green <= 13 && game.blue <= 14
        ? game.gameId
        : 0),
  0
);
console.log(part1);

//part2
const part2 = games.reduce(
  (sum, game) => (sum += game.red * game.green * game.blue),
  0
);
console.log(part2);
