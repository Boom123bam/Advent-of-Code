const fs = require("node:fs");

function readFileAsync(filePath) {
  return new Promise((resolve, reject) => {
    fs.readFile(filePath, "utf8", (err, data) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(data);
    });
  });
}

(async () => {
  const input = await readFileAsync("day-1/input.txt");

  const sections = input
    .split("\n\n")
    .map((section) => section.split("\n").map(Number));
  const sums = sections
    .map((section) => section.reduce((sum, num) => sum + num), 0)
    .sort((a, b) => b - a);
  console.log(sums[0]);
  console.log(sums.slice(0, 3).reduce((sum, num) => sum + num, 0));
})();
