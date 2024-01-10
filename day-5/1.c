#include "./stacks.h"
#include <stdio.h>

#define isNum(c) (c >= '0' && c <= '9')
#define getLetterIndexInString(i) (i * 4 + 1)
#define MAX_INSTRUCTIONS 1000

int getNumCols(char *firstCharPointer);
int getNextNum(char lineString[]);

int main() {
  FILE *fptr;
  fptr = fopen("day-5/testinput.txt", "r");
  char lineString[100];
  int numCols;
  int lastInstructionIndex = 0;
  int colFrom, colTo, amount;
  char c;

  getStackTopPointers();

  int instructions[MAX_INSTRUCTIONS][3];

  // read first line to get number of columns
  fgets(lineString, 100, fptr);
  numCols = getNumCols(lineString);

  // read crates
  do {
    // read letters in line
    for (int i = 0; i < numCols; i++) {
      c = lineString[(getLetterIndexInString(i))];

      if (c != ' ') {
        pushBottom(i, c);
      }
    }
  } while (fgets(lineString, 100, fptr) && lineString[1] != '1');

  fgets(lineString, 100, fptr); // skip blank line

  // read instructions
  for (int i = 0; fgets(lineString, 100, fptr); i++) {
    if (lineString[0] == '\n')
      break;
    instructions[i][0] = getNextNum(lineString);
    instructions[i][1] = getNextNum(0);
    instructions[i][2] = getNextNum(0);
    lastInstructionIndex++;
  }

  // execute instructions
  for (int i = 0; i < lastInstructionIndex; i++) {
    amount = instructions[i][0];
    colFrom = instructions[i][1] - 1;
    colTo = instructions[i][2] - 1;
    for (int count = 0; count < amount; count++) {
      c = pop(colFrom);
      push(colTo, c);
    }
  }

  // printStacks(numCols);
  printStackTops(numCols);
}

int getNumCols(char *firstCharPointer) {
  int lineLen;
  for (lineLen = 0; *firstCharPointer++ != '\n'; lineLen++) {
  }
  return (lineLen + 1) / 4;
}

int getNextNum(char *firstCharPointer) {
  static char *currentCharPointer;
  if (firstCharPointer)
    currentCharPointer = firstCharPointer;
  int result = 0;

  while (!isNum(*currentCharPointer)) {
    currentCharPointer++;
  }
  while (isNum(*currentCharPointer)) {
    result = result * 10 + *currentCharPointer - '0';
    currentCharPointer++;
  }

  return result;
}
