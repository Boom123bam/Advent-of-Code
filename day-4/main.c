#include <stdio.h>

int getNextNum(char lineString[]);

int main() {
  FILE *fptr;
  // Open a file in read mode
  fptr = fopen("day-4/input2.txt", "r");
  char lineString[100];
  int p1count = 0;
  int p2count = 0;
  int s1, s2, e1, e2;

  while (fgets(lineString, 100, fptr)) {
    // printf("%s\n", lineString);
    s1 = getNextNum(lineString);
    e1 = getNextNum(lineString);
    s2 = getNextNum(lineString);
    e2 = getNextNum(lineString);

    if ((s2 >= s1 && e2 <= e1) || (s1 >= s2 && e1 <= e2))
      p1count++;
    if (!((e1 < s2) || (e2 < s1)))
      p2count++;
  }

  printf("part 1: %d\npart 2: %d\n", p1count, p2count);

  // Close the file
  fclose(fptr);
}

int getNextNum(char lineString[]) {
  static int index = 0;
  int result = 0;
  char c;

  while ((c = lineString[index++]) != '\n') {
    if (c >= '0' && c <= '9')
      result = 10 * result + (c - '0');

    else
      return result;
  }

  index = 0;
  return result;
}
