#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <fstream>
#include <string>
#include <list>
using namespace std;

#define MAX_INSTRUCTIONS 1000

void printList(list<char> matrix[], int numArrays);
void getCrates(int numArrays, list<char> matrix[]);
int getNextNum(string lineString);
void move(list<char> *&matrix, int numArrays, int colFrom, int colTo);

ifstream MyReadFile("testinput.txt");

int main()
{

  string lineString;

  getline(MyReadFile, lineString);

  // get num of arrays
  int numArrays = (lineString.length() + 1) / 4;

  // list<char> matrix[numArrays]; // int arr[5];
  list<char> *matrix = new list<char>[numArrays];

  int instructions[3][MAX_INSTRUCTIONS];
  int counter = 0;

  // get crates
  getCrates(numArrays, matrix);

  // skip blank line
  getline(MyReadFile, lineString);

  // Read Instructions
  while (getline(MyReadFile, lineString))
  {
    instructions[counter][0] = getNextNum(lineString);
    instructions[counter][1] = getNextNum(lineString);
    instructions[counter][2] = getNextNum(lineString);
    counter++;
  }

  move(matrix, 3, 0, 1);

  printList(matrix, numArrays);

  return 0;
}

void move(list<char> *&matrix, int numArrays, int colFrom, int colTo)
{
  matrix[colFrom].reverse();
  char object = matrix[colFrom].front();
  matrix[colFrom].reverse();

  cout << ":" << object;

  (matrix[colFrom]).pop_back();
  (matrix[colTo]).push_back(object);
  printList(matrix, numArrays);
  // 1, 2, 3
}

/*
[
  ["Z","N"],
  ["M","C","D"],
  ["P"],
]
*/

void getCrates(int numArrays, list<char> matrix[])
{
  string lineString;
  do
  {
    if (lineString[1] == '1')
      break;
    for (int column = 0; column < numArrays; column++)
    {
      int index = (column * 4 + 1);
      char c = lineString[index];
      if (c != ' ')
      {
        // add to array
        (matrix[column]).push_front(c);
      }
    }
  } while (getline(MyReadFile, lineString));
}

void printList(list<char> matrix[], int numArrays)
{
  for (int i = 0; i < numArrays; i++)
  {
    list<char> l = matrix[i];
    list<char>::const_iterator j;
    for (j = l.begin(); j != l.end(); ++j)
      cout << *j << " "; // l[j]
    cout << endl;
  }
}

int getNextNum(string lineString)
{
  static int index = 0;

  for (; index < lineString.length(); index++)
  {
    if (lineString[index] >= '0' && lineString[index] <= '9')
    {
      break;
    }
  }
  int result = 0;
  char c = lineString[index];

  while ((c = lineString[index++]) != '\0')
  {
    if (c >= '0' && c <= '9')
      result = 10 * result + (c - '0');

    else
      return result;
  }

  index = 0;
  return result;
}