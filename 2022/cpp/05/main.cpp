#include <cstdio>
#include <fstream>
#include <iostream>
#include <list>
using namespace std;

#define NUM_STACKS 9

struct Node {
    char data;
    Node *up;
    Node *down;

    Node(char d) : data(d), up(nullptr), down(nullptr) {}
};

struct Instruction {
    int amount;
    int from;
    int to;
};

void printStacks(Node *[]);
void move(Node *[], Instruction *);
void move2(Node *[], Instruction *);
void printTops(Node *[]);

int main() {
    string line;
    ifstream MyReadFile("input.txt");

    char c;
    Node *stackPtrs[NUM_STACKS + 1];

    while (getline(MyReadFile, line) && line.substr(0, 2) != " 1") {
        for (int i = 1, j = 0; i < line.length(); i += 4, j++) {
            c = line[i];
            if (c == ' ')
                continue;

            if (stackPtrs[j] == nullptr) {
                stackPtrs[j] = new Node(c);
            } else {
                Node *node = new Node(c);
                node->up = stackPtrs[j];
                stackPtrs[j]->down = node;
                stackPtrs[j] = node;
            }
        }
    }
    // reset ptrs to top
    for (int i = 0; i < NUM_STACKS; i++) {
        if (stackPtrs[i] == nullptr)
            continue;
        for (; stackPtrs[i]->up != nullptr; stackPtrs[i] = stackPtrs[i]->up) {
        }
    }
    getline(MyReadFile, line);

    list<Instruction> instructions;
    while (getline(MyReadFile, line)) {
        Instruction instruction;
        sscanf(line.c_str(), "move %d from %d to %d", &instruction.amount,
               &instruction.from, &instruction.to);
        instruction.from--;
        instruction.to--;
        instructions.push_back(instruction);
    }

    // printStacks(stackPtrs);
    for (Instruction ins : instructions) {
        move(stackPtrs, &ins);
        // move2(stackPtrs, &ins);
        // printStacks(stackPtrs);
    }
    printTops(stackPtrs);
}

void printStacks(Node *stackPtrs[]) {
    string line;
    for (int i = 0; i < NUM_STACKS; i++) {
        line = "";
        cout << i + 1 << ": ";
        for (Node *n = stackPtrs[i]; n != nullptr; n = n->down) {
            line = n->data + line;
        }
        cout << line << "\n";
    }
}

void move(Node *stackPtrs[], Instruction *ins) {
    for (int i = 0; i < ins->amount; i++) {
        Node *tmp = stackPtrs[ins->from];
        // cout << "from " << ins->from + 1 << " to " << ins->to + 1 << "("
        //      << tmp->data << ")\n";
        stackPtrs[ins->from] = tmp->down;
        if (stackPtrs[ins->from] != nullptr)
            stackPtrs[ins->from]->up = nullptr;
        tmp->down = stackPtrs[ins->to];
        if (stackPtrs[ins->to] != nullptr)
            stackPtrs[ins->to]->up = tmp;
        stackPtrs[ins->to] = tmp;
    }
}

void move2(Node *stackPtrs[], Instruction *ins) {
    Instruction i1 = {
        .amount = ins->amount, .from = ins->from, .to = NUM_STACKS};
    Instruction i2 = {.amount = ins->amount, .from = NUM_STACKS, .to = ins->to};
    move(stackPtrs, &i1);
    move(stackPtrs, &i2);
}

void printTops(Node *stackPtrs[]) {
    for (int i = 0; i < NUM_STACKS; i++) {
        if (stackPtrs[i] == nullptr)
            continue;
        cout << stackPtrs[i]->data;
    }
    cout << "\n";
}
