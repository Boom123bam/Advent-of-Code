#ifndef STACKS_H
#define STACKS_H

void getStackTopPointers();
void push(int stackIndex, char c);
void pushBottom(int stackIndex, char c);
char pop(int stackIndex);
void printStackTops(int num_stacks);
void printStack(int stackIndex);
void printStacks(int numStacks);

#endif
