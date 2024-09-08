let input = await Deno.readTextFile("06/input.txt");
if (input[input.length - 1] == "\n") {
  input = input.slice(0, -1);
}

function getHandTypeScore(hand: string) {
  const cards = hand.split("").sort();
  let cardsCount: number[] = [];
  cards.forEach((card, index) => {
    const lastCard = index == 0 ? null : cards[index - 1];
    if (card == lastCard) cardsCount[cardsCount.length - 1]++;
    else cardsCount.push(1);
  });
  cardsCount = cardsCount.sort();

  const combinations = [
    "1,1,1,1,1",
    "1,1,1,2",
    "1,2,2",
    "1,1,3",
    "2,3",
    "1,4",
    "5",
  ];

  return combinations.indexOf(cardsCount.toString());
}

function compareCards(card1: string, card2: string) {
  const order = "AKQJT98765432";
  let card1Index = order.indexOf(card1),
    card2Index = order.indexOf(card2);
  return card2Index - card1Index;
}

const handsData = input.split("\n").map((line) => {
  const [hand, bid] = line.split(" ");
  return { hand, bid: Number(bid) };
});

// sort by rank

handsData.sort((a, b) => {
  const aHandTypeScore = getHandTypeScore(a.hand);
  const bHandTypeScore = getHandTypeScore(b.hand);
  if (aHandTypeScore != bHandTypeScore) return aHandTypeScore - bHandTypeScore;
  else {
    // compare first different card
    for (let i = 0; i < a.hand.length; i++) {
      const [cardA, cardB] = [a.hand[i], b.hand[i]];
      if (cardA != cardB) return compareCards(cardA, cardB);
    }
    return 0;
  }
});

const winnings = handsData.reduce((total, handsData, index) => {
  const rank = index + 1;
  return total + handsData.bid * rank;
}, 0);

console.log(winnings);
