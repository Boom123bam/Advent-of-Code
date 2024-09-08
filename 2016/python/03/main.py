f = open("03/input.txt", "r").read()
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()


def valid(triangle):
    l = sorted(triangle)
    return l[0] + l[1] > l[2]


parsed_lines = [[int(n) for n in l.split()] for l in lines]
possible = [line for line in parsed_lines if valid(line)]
print(len(possible))

count = 0
for y in range(0, len(parsed_lines), 3):
    for x in range(3):
        if valid([parsed_lines[y][x], parsed_lines[y + 1][x], parsed_lines[y + 2][x]]):
            count += 1

print(count)
