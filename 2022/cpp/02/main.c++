#include <fstream>
#include <iostream>
#include <map>
using namespace std;
#define ROCK 1
#define PAPER 2
#define SCISSORS 3

#define LOSS 0
#define DRAW 3
#define WIN 6

int RPS[3] = {ROCK, PAPER, SCISSORS};
int score(int me, int opp);

int main() {
    map<char, int> oppMap = {{'A', ROCK}, {'B', PAPER}, {'C', SCISSORS}};
    map<char, int> meMap = {{'X', ROCK}, {'Y', PAPER}, {'Z', SCISSORS}};
    string line;
    // ifstream MyReadFile("test.txt");
    ifstream MyReadFile("input.txt");
    int opp, me;
    int sumP1, sumP2;
    while (getline(MyReadFile, line)) {
        opp = oppMap[line[0]];
        me = meMap[line[2]];
        sumP1 += score(me, opp);

        // Part 2
        switch (line[2]) {
        case 'X':
            // lose
            for (int i = 0; i < 3; i++) {
                if (RPS[i] == opp) {
                    me = RPS[i == 0 ? 2 : i - 1];
                }
            }
            break;
        case 'Y':
            // draw
            me = opp;
            break;
        case 'Z':
            // win
            for (int i = 0; i < 3; i++) {
                if (RPS[i] == opp) {
                    me = RPS[i == 2 ? 0 : i + 1];
                }
            }
            break;
        }
        sumP2 += score(me, opp);
        // cout << me << opp << "  ";
        // cout << score(me, opp) << "\n";
    }
    cout << sumP1 << "\n";
    cout << sumP2 << "\n";
}

int score(int me, int opp) {
    if (me == opp) {
        // cout << "d";
        return DRAW + me;
    }
    int beatsMe;
    for (int i = 0; i < 3; i++) {
        if (RPS[i] == me) {
            beatsMe = RPS[i == 2 ? 0 : i + 1];
        }
    }
    if (opp == beatsMe) {
        // cout << "l";
        return LOSS + me;
    }
    // cout << "w";
    return WIN + me;
}
