def get_corrupted_char_or_matches(line):
    matches = {"(": ")", "[": "]", "{": "}", "<": ">"}

    expected_matches = []
    opening_chars = matches.keys()

    for char in line:
        if char in opening_chars:
            expected_matches.append(matches[char])
        elif char == expected_matches[-1]:
            expected_matches.pop(-1)
        else:
            return [char, []]
    return [None, expected_matches]


input = open("10/test.txt").read()
input = open("10/input.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")


opening_scores = {
    ")": 3,
    "]": 57,
    "}": 1197,
    ">": 25137,
}

closing_scores = {
    ")": 1,
    "]": 2,
    "}": 3,
    ">": 4,
}

p1_score = 0
p2_scores = []
for line in lines:
    [corrupted_char, closing_matches] = get_corrupted_char_or_matches(line)
    if corrupted_char:
        p1_score += opening_scores[corrupted_char]
    else:
        score = 0
        while len(closing_matches):
            match = closing_matches.pop(-1)
            score *= 5
            score += closing_scores[match]
        p2_scores.append(score)

p2_scores.sort()

print(p1_score)
print(p2_scores[int((len(p2_scores) - 1) / 2)])
