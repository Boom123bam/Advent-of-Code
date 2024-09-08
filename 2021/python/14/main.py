input = open("14/input.txt").read()
# input = open("14/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")


def update_pair_counts(pair_counts):
    pairs = pair_counts.keys()
    new_pair_counts = {pair: 0 for pair in pairs}
    for pair in pairs:
        pair_to_1, pair_to_2 = insertion_pairs_map[pair]
        new_pair_counts[pair_to_1] += pair_counts[pair]
        new_pair_counts[pair_to_2] += pair_counts[pair]
    return new_pair_counts


def get_char_counts(pair_counts):
    global start_polymer_str
    chars = set("".join(pair_counts.keys()))
    # each char is in 2 pairs so is counted twice
    # except the first and last chars, so add them in
    char_double_counts = {char: 0 for char in chars}
    char_double_counts[start_polymer_str[0]] += 1
    char_double_counts[start_polymer_str[-1]] += 1
    pairs = pair_counts.keys()
    for pair in pairs:
        for char in pair:
            char_double_counts[char] += pair_counts[pair]

    return {char: int(count / 2) for char, count in char_double_counts.items()}


def get_min_max_diff(char_counts):
    max_char = max(char_counts.keys(), key=char_counts.get)
    min_char = min(char_counts.keys(), key=char_counts.get)
    return char_counts[max_char] - char_counts[min_char]


start_polymer_str = lines.pop(0)
lines.pop(0)  # skip blank line
instructions = [line.split(" -> ") for line in lines]
insertion_pairs_map = {}
pair_counts = {}

# add instructions to insertion_pairs_map
for pair, char_to_insert in instructions:
    # map ab->c to (ac,cb)
    insertion_pairs_map.update(
        {pair: (pair[0] + char_to_insert, char_to_insert + pair[1])}
    )
    # set pair count to 0
    pair_counts[pair] = 0

# add initial pair counts
for i in range(1, len(start_polymer_str)):
    pair = start_polymer_str[i - 1] + start_polymer_str[i]
    pair_counts[pair] += 1

# part 1
for i in range(10):
    pair_counts = update_pair_counts(pair_counts)

char_counts = get_char_counts(pair_counts)
print(get_min_max_diff(char_counts))

# part 2
for i in range(30):
    pair_counts = update_pair_counts(pair_counts)

char_counts = get_char_counts(pair_counts)
print(get_min_max_diff(char_counts))
