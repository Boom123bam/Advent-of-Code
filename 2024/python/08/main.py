file = open("08/input.txt")
# file = open("08/test.txt")
input = file.read().strip()
lines = input.split("\n")

is_in_bounds = lambda coord : coord[0] in range(0, len(lines)) and coord[1] in range(0,len(lines[0]))
antennas = {}

for r,line in enumerate(lines):
    for c, char in enumerate(line):
        if char != ".":
            if char not in antennas:
                antennas[char] = []
            antennas[char].append((r,c))

def get_pair_perms(items):
    result = []
    for i,a in enumerate(items[:-1]):
        for b in items[i+1:]:
            result.append((a,b))
    return result

def get_antinodes(coord_1, coord_2, p1):
    r1, c1 = coord_1
    r2, c2 = coord_2
    diff = (r1-r2, c1-c2)
    result = []
    if p1:
        a1, a2 = (r1 + diff[0], c1 + diff[1]),(r2 - diff[0], c2 - diff[1])
        if is_in_bounds(a1):
            result.append(a1)
        if is_in_bounds(a2):
            result.append(a2)
    else:
        a = coord_1
        while(is_in_bounds(a)):
            result.append(a)
            a = (a[0] + diff[0], a[1] + diff[1])

        a = coord_2
        while(is_in_bounds(a)):
            result.append(a)
            a = (a[0] - diff[0], a[1] - diff[1])

    return result

p1_antinodes = {}
p2_antinodes = {}
for symbol,coords in antennas.items():
    perms = get_pair_perms(coords)
    for perm in perms:
        a1 = get_antinodes(perm[0], perm[1], True)
        for a in a1:
            p1_antinodes[a] = True

        a2 = get_antinodes(perm[0], perm[1], False)
        for a in a2:
            p2_antinodes[a] = True

# print(p2_antinodes)
print(len(p1_antinodes))
print(len(p2_antinodes))
