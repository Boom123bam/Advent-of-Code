#include <fstream>
#include <iostream>
using namespace std;

class Packet {
  public:
    int start1;
    int start2;
    string data;
    Packet(string);

  private:
    int find_distinct(int);
};
bool stringHasSameChar(string);

Packet::Packet(string data) : data(data) {
    this->start1 = find_distinct(4);
    this->start2 = find_distinct(14);
};

int Packet::find_distinct(int n) {
    string chars;
    for (int i = 0; i < data.length() - n - 1; i++) {
        chars = data.substr(i, n);
        if (!stringHasSameChar(chars))
            return i + n;
    }
    return -1;
};

bool stringHasSameChar(string s) {
    for (int i = 0; i < s.length(); i++) {
        for (int j = 0; j < s.length(); j++) {
            if (i == j)
                continue;
            if (s[i] == s[j])
                return true;
        }
    }
    return false;
}

int main() {
    string input;
    ifstream MyReadFile("input.txt");
    getline(MyReadFile, input);

    Packet p(input);
    cout << p.start1 << "\n";
    cout << p.start2 << "\n";
}
