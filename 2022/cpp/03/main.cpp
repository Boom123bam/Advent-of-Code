#include <fstream>
#include <iostream>
#include <string>
using namespace std;

bool contains(string, char);
int find_priority(char);
string find_common(string, string);

int main() {
    string line;
    ifstream MyReadFile("input.txt");
    string s1, s2;
    int sumP1, sumP2;
    string comm;
    int lineNum;
    while (getline(MyReadFile, line)) {
        s1 = line.substr(0, line.length() / 2);
        s2 = line.substr(line.length() / 2);
        sumP1 += find_priority(find_common(s1, s2)[0]);

        if (lineNum % 3 == 0) {
            if (lineNum != 0) {
                sumP2 += find_priority(comm[0]);
            }
            comm = line;
        } else {
            comm = find_common(comm, line);
        }
        lineNum++;
    }
    sumP2 += find_priority(comm[0]);
    cout << sumP1 << "\n";
    cout << sumP2 << "\n";
}

bool contains(string s, char target) {
    for (char c : s) {
        if (c == target) {
            return true;
        }
    }
    return false;
}

string find_common(string s1, string s2) {
    string result;
    for (char c : s1) {
        if (contains(s2, c)) {
            result += c;
        }
    }
    return result;
}

int find_priority(char c) {
    if (c >= 'a' and c <= 'z') {
        return c - 'a' + 1;
    }
    return c - 'A' + 27;
}
