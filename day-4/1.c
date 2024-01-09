#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int getNextNum(char lineString[]);

int main()
{
  FILE *fptr;
  // Open a file in read mode
  fptr = fopen("input2.txt", "r");
  char lineString[100];
  int count = 0;
  int s1, s2, e1, e2;

  while (fgets(lineString, 100, fptr))
  {
    // printf("%s\n", lineString);
    s1 = getNextNum(lineString);
    e1 = getNextNum(lineString);
    s2 = getNextNum(lineString);
    e2 = getNextNum(lineString);

    // if ((s2 >= s1 && e2 <= e1) || (s1 >= s2 && e1 <= e2))
    if (!((e1 < s2) || (e2 < s1)))
      count++;
  }

  printf("%d\n", count);

  // Close the file
  fclose(fptr);
}

int getNextNum(char lineString[])
{
  static int index = 0;
  int result = 0;
  char c;

  while ((c = lineString[index++]) != '\n')
  {
    if (c >= '0' && c <= '9')
      result = 10 * result + (c - '0');

    else
      return result;
  }

  index = 0;
  return result;
}

/*

s1 l1 s2 l2
if 2 is inside 1:
s2 >= s1 && s1 + l1 <= s2...

s1-e1 s2-e2
if 2 is inside 1:
s2 >= s1 && e2 <= e1

or s1 >= s2 && e1 <= e2
*/
