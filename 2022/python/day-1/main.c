#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int cmpfunc(const void *a, const void *b)
{
  return (*(int *)a - *(int *)b);
}

int main()
{
  FILE *fptr;

  // Open a file in read mode
  fptr = fopen("input.txt", "r");

  char lineString[100];
  int currentSum = 0;
  int maxSums[3] = {0, 0, 0};
  while (fgets(lineString, 100, fptr))
  {
    if (lineString[0] == '\n')
    {
      if (currentSum > maxSums[0])
      {
        maxSums[0] = currentSum;
        qsort(maxSums, 3, sizeof(int), cmpfunc);
      }
      currentSum = 0;
      continue;
    }

    int len = strlen(lineString);

    if (lineString[len - 1] == '\n') // cut new line char
    {
      lineString[len - 1] = '\0';
    }

    currentSum += atoi(lineString);
  }

  if (currentSum > maxSums[0])
  {
    maxSums[0] = currentSum;
    qsort(maxSums, 3, sizeof(int), cmpfunc);
  }

  // // Close the file
  fclose(fptr);

  printf("part 1: %d\npart 2: %d\n", maxSums[2], maxSums[0] + maxSums[1] + maxSums[2]);
}