#include <fstream>
#include <iostream>
using namespace std;

int main() {
  string line;
  ifstream MyReadFile("test.txt");
  while (getline(MyReadFile, line)) {
      cout << line << "\n";
    }
}
