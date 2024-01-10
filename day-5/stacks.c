#include "./stacks.h"
#include <stdio.h>

#define MAX_NUM_STACKS 50
#define MAX_STACK_HEIGHT 1000

static char stacks[MAX_STACK_HEIGHT][MAX_NUM_STACKS];
static char *stackTopPointers[MAX_NUM_STACKS];

void getStackTopPointers() {
  for (int i = 0; i < MAX_NUM_STACKS; i++) {
    stackTopPointers[i] = &stacks[i][0];
  }
}

void push(int stackIndex, char c) { *stackTopPointers[stackIndex]++ = c; }

void pushBottom(int stackIndex, char c) {
  // find stack bottom
  char *stackBottomPointer = &stacks[stackIndex][0];
  // set current to top
  char *currentPointer = stackTopPointers[stackIndex]++;
  // loop from top to bottom and shift up
  for (; currentPointer > stackBottomPointer; currentPointer--) {
    *currentPointer = *(currentPointer - 1);
  }

  *stackBottomPointer = c;
}

char pop(int stackIndex) { return *--stackTopPointers[stackIndex]; }

void printStackTops(int numStacks) {
  char c;
  for (int i = 0; i < numStacks; i++) {
    c = *(stackTopPointers[i] - 1);
    printf("%c ", c ? c : '_');
  }
  printf("\n");
}

void printStack(int stackIndex) {
  // find stack bottom
  char *currentPointer = &stacks[stackIndex][0];
  // set current to top
  char *stackTopPointer = stackTopPointers[stackIndex] - 1;

  for (; currentPointer <= stackTopPointer; currentPointer++) {
    printf("%c ", *currentPointer);
  }
  printf("\n");
}

void printStacks(int numStacks) {
  for (int i = 0; i < numStacks; i++) {
    printStack(i);
  }
}
