f = open("05/input.txt", "r").read()
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()


polymer = lines[0]
# polymer = "dabAcCaCBAcCcaDA"

alphabet = "abcdefghijklmnopqrstuvwxyz"


def react(polymer):
    i = 1
    while i < len(polymer):
        if (
            polymer[i - 1].lower() == polymer[i].lower()
            and polymer[i - 1] != polymer[i]
        ):
            polymer = polymer[: i - 1] + polymer[i + 1 :]

        i += 1
    return polymer


def get_final_len(polymer):
    newPolymer = react(polymer)
    while newPolymer != polymer:
        polymer = newPolymer
        newPolymer = react(polymer)

    return len(polymer)


print(get_final_len(polymer))


shortestLen = 99999
for letter in alphabet:
    l = get_final_len(polymer.replace(letter, "").replace(letter.upper(), ""))
    if l < shortestLen:
        shortestLen = l

print(shortestLen)
