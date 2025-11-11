#include <fstream>
#include <iostream>
#include <string>
using namespace std;

struct section {
    int aStart;
    int aEnd;
    int bStart;
    int bEnd;
};
bool inRange(int &, int &, int &);
bool intersects(section &);

int main() {
    string line;
    ifstream MyReadFile("input.txt");
    section s;
    int countP1;
    int countP2;
    while (getline(MyReadFile, line)) {
        sscanf(line.c_str(), "%d-%d,%d-%d", &s.aStart, &s.aEnd, &s.bStart,
               &s.bEnd);
        if (s.aStart <= s.bStart && s.aEnd >= s.bEnd) {
            countP1++;
        } else if (s.bStart <= s.aStart && s.bEnd >= s.aEnd) {
            countP1++;
        }
        if (intersects(s))
            countP2++;
    }
    cout << countP1 << "\n";
    cout << countP2 << "\n";
}

bool inRange(int &start, int &end, int &n) { return n >= start and n <= end; }

bool intersects(section &s) {
    return inRange(s.aStart, s.aEnd, s.bStart) ||
           inRange(s.aStart, s.aEnd, s.bEnd) ||
           inRange(s.bStart, s.bEnd, s.aStart) ||
           inRange(s.bStart, s.bEnd, s.aEnd);
}
