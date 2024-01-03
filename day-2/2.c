#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ROCK 0
#define PAPER 1
#define SCISSORS 2

#define R_SCORE 1
#define P_SCORE 2
#define S_SCORE 3

#define WIN_SCORE 6
#define DRAW_SCORE 3
#define LOSS_SCORE 0

#define LOSS 'X'
#define DRAW 'Y'
#define WIN 'Z'


int getType(char letter){
	char rps[3][2] = {"AX", "BY", "CZ"};
	for(int i=0; i < 3; i++){
		if (rps[i][0] == letter || rps[i][1] == letter) return i;
	}
	return -1;
}

int main(){
	FILE *fptr;

	// Open a file in read mode
	fptr = fopen("day-2/input.txt", "r");

	char lineString[100];
	int totalScore = 0;
	int maxSums[3] = {0, 0, 0};

	char outcome;
	int me, opponent;
	int rpsScores[3] = {R_SCORE, P_SCORE, S_SCORE};

	int beats[3];
	beats[ROCK] = SCISSORS;
	beats[PAPER] = ROCK;
	beats[SCISSORS] = PAPER;

	while (fgets(lineString, 100, fptr))
	{
		opponent = getType(lineString[0]);
		outcome = lineString[2];
		// me = getType(lineString[2]);

		if (outcome == DRAW) {
			// draw
			totalScore += DRAW_SCORE;
			me = opponent;
		}
		else if (outcome == WIN) {
			// win
			totalScore += WIN_SCORE;
			me = beats[beats[opponent]];
		}
		else {
			totalScore += LOSS_SCORE;
			me = beats[opponent];
		}

		totalScore += rpsScores[me];

		// printf("%c %c %d\n",
	 // 		me == ROCK ? 'r' : me == PAPER ? 'p' : 's',
	 // 		opponent == ROCK ? 'r' : me == PAPER ? 'p' : 's',
		// 	totalScore
	 // 	);
	}
	printf("%d\n", totalScore);
}

