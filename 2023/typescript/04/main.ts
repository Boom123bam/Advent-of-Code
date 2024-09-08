let input = await Deno.readTextFile("04/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

const lines = input.split("\n");

interface Card {
  winning: number[];
  have: number[];
  amount: number;
}

const cardsData: Card[] = [];

lines.forEach((line) => {
  const parts = line.split(":")[1].split("|");

  const card: Card = {
    amount: 1,
    winning: parts[0].match(/\d+/g)?.map((st) => Number(st)) ?? [],
    have: parts[1].match(/\d+/g)?.map((st) => Number(st)) ?? [],
  };
  cardsData.push(card);
});

const sum1 = cardsData.reduce((acc, cardData, cardIndex) => {
  let matches = 0;
  cardData.have.forEach((num) => {
    if (cardData.winning.includes(num)) matches++;
  });

  for (let i = 1; i <= matches; i++) {
    if (typeof cardsData[cardIndex + i] !== "undefined")
      cardsData[cardIndex + i].amount += cardData.amount;
  }

  return matches ? acc + 2 ** (matches - 1) : acc;
}, 0);

console.log(sum1);

let sum2 = cardsData.reduce((acc, card) => card.amount + acc, 0);

console.log(sum2);
