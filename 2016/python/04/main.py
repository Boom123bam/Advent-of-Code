input = """aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]"""

input = open("04/input.txt", "r").read()


def parse_line(line):
    parts = line.replace("[", "-")[:-1].split("-")
    name = " ".join(parts[:-2])
    id = int(parts[-2])
    checksum = parts[-1]
    return [name, id, checksum]


rooms = [parse_line(line) for line in input.split("\n") if line]

alphabet = "abcdefghijklmnopqrstuvwxyz"


def get_expected_checksum(name):
    counts = {}
    for letter in name:
        if letter in alphabet:
            if letter not in counts:
                counts[letter] = 1
            else:
                counts[letter] += 1
    return "".join(
        sorted(
            counts.keys(),
            key=lambda k: counts[k] * 26 - alphabet.index(k),
            reverse=True,
        )[:5]
    )


def shift(input, amount):
    min_amount = amount % len(alphabet)
    result = list(input)
    for i, letter in enumerate(input):
        if letter not in alphabet:
            continue
        result[i] = alphabet[(alphabet.index(letter) + min_amount) % len(alphabet)]
    return "".join(result)


sum = 0
for room in rooms:
    [name, id, checksum] = room
    if get_expected_checksum(name) == checksum:
        sum += id
        print(shift(name, id), id)

print(sum)
