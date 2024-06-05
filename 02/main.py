f = open("02/input.txt", "r").read()
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()

triple_count = 0
double_count = 0

letters = "abcdefghijklmnopqrstuvwxyz"


def has_double(line):
    chars = list(line)
    for letter in letters:
        if chars.count(letter) == 2:
            return True
    return False


def has_triple(line):
    chars = list(line)
    for letter in letters:
        if chars.count(letter) == 3:
            return True
    return False


for line in lines:
    if has_double(line):
        double_count += 1
    if has_triple(line):
        triple_count += 1

print(double_count * triple_count)


def get_diff_char_index(id1, id2):
    # return -1 if more than 1 different
    # assume equal len
    idx = -1
    for i in range(len(id1)):
        if id1[i] != id2[i]:
            if idx != -1:
                return -1
            idx = i
    return idx


for i in range(len(lines) - 1):
    for j in range(i + 1, len(lines)):
        result = get_diff_char_index(lines[i], lines[j])
        if result != -1:
            print(lines[i][:result] + lines[i][result + 1 :])
