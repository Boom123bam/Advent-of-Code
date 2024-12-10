file = open("06/input.txt")
# file = open("06/test.txt")
input = file.read().strip()
grid = [list(line) for line in input.split("\n")]

dirs = ["^", ">", "v", "<"]
offs = [[-1, 0], [0, 1], [1, 0], [0, -1]]
r, c = -1, -1

for rn, row in enumerate(grid):
    for cn, letter in enumerate(row):
        if letter == "^":
            r,c = rn, cn


def getPositions(grid, r,c):
    dir = 0
    p1 = 0
    while True:
        peekR, peekC = r + offs[dir][0], c + offs[dir][1]
        if peekR not in range(0,len(grid)) or peekC not in range(0,len(grid[0])):
            p1 += 1
            break
        peekChar = grid[peekR][peekC]
        if peekChar == ".":
            r, c = peekR, peekC
            grid[r][c] = dirs[dir]
            p1 += 1
        elif peekChar == "#":
            dir += 1
            if dir == 4:
                dir = 0
        elif peekChar == dirs[dir]:
            # loop detected
            return -1
        else:
            r, c = peekR, peekC
            grid[r][c] = dirs[dir]
    # print("\n".join(["".join(row) for row in grid]))
    return p1

print(getPositions([row[:] for row in grid], r, c))

p2 = 0
for rn, row in enumerate(grid):
    for cn, char in enumerate(row):
        if char != ".":
            continue
        gridCopy = [row[:] for row in grid]
        gridCopy[rn][cn] = "#"
        if getPositions(gridCopy, r, c) == -1:
            p2 += 1

print(p2)
