#include <fstream>
#include <iostream>
#include <list>
#include <stdexcept>
#include <string>
using namespace std;

int main() {
  // Create a text string, which is used to output the text file
  string line;
  int sum = 0;
  ifstream MyReadFile("input.txt");
  int maxElfs[3] = {0,0,0};

  while (getline(MyReadFile, line)) {
    try {
      sum += stoi(line);
    } catch (invalid_argument) {
        for (int i = 0; i < 3; i++) {
            if (maxElfs[i] < sum){
                for (int j=2; j>i; j--){
                    maxElfs[j] = maxElfs[j-1];
                }
                maxElfs[i] = sum;
                break;
            }
        }
      sum = 0;
    }
  }
  MyReadFile.close();

  for (int i = 0; i < 3; i++) {
      if (maxElfs[i] < sum){
          maxElfs[i] = sum;
          break;
      }
  }

  cout << maxElfs[0] << "\n";

  sum = 0;
  for (int i = 0; i < 3; i++){
      sum += maxElfs[i];
  }

  cout << sum << "\n";
  return 0;
}
