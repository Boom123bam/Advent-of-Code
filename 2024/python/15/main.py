file = open("15/input.txt")
# file = open("15/test.txt")
input = file.read().strip()
map, moves = input.split("\n\n")
grid = map.split("\n")

def find_pos(grid):
    for r, row in enumerate(grid):
        for c, col in enumerate(row):
            if col == "@":
                return r,c
    return -1, -1

def update_grid(grid, move):
    r, c = find_pos(grid)
    if move == '<':
        c2 = c - 1
        while grid[r][c2] != "#":
            if grid[r][c2] == ".":
                grid[r] = grid[r][:c2] + grid[r][c2+1:c+1] + "." + grid[r][c+1:]
                break
            c2 -= 1
    elif move == ">":
        c2 = c + 1
        while grid[r][c2] != "#":
            if grid[r][c2] == ".":
                grid[r] = grid[r][:c] + "." + grid[r][c:c2] + grid[r][c2+1:]
                break
            c2 += 1
    elif move == "^":
        r2 = r - 1
        while grid[r2][c] != "#":
            if grid[r2][c] == ".":
                for row in range(r2, r):
                    grid[row] = grid[row][:c] + grid[row + 1][c] + grid[row][c+1:]
                grid[r] = grid[r][:c] + "." + grid[r][c+1:]
                break
            r2 -= 1
    elif move == "v":
        r2 = r + 1
        while grid[r2][c] != "#":
            if grid[r2][c] == ".":
                for row in range(r2, r, -1):
                    grid[row] = grid[row][:c] + grid[row - 1][c] + grid[row][c+1:]
                grid[r] = grid[r][:c] + "." + grid[r][c+1:]
                break
            r2 += 1
    return grid


for move in moves:
    grid = update_grid(grid, move)
    # print("\n".join(grid) + "\n")

p1 = 0
for r, row in enumerate(grid):
    for c, col in enumerate(row):
        if col == "O":
            p1 += r * 100 + c

print(p1)

grid_2 = [list(line.replace("#", "##").replace(".", "..").replace("O", "[]").replace("@", "@.")) for line in map.split("\n")]

def update_grid_2(grid, move):
    r, c = find_pos(grid)
    if move == '<':
        c2 = c - 1
        while grid[r][c2] != "#":
            if grid[r][c2] == ".":
                for col in range(c2, c):
                    grid[r][col] = grid[r][col + 1]
                grid[r][c] = "."
                break
            c2 -= 1

    elif move == ">":
        c2 = c + 1
        while grid[r][c2] != "#":
            if grid[r][c2] == ".":
                for col in range(c2, c, -1):
                    grid[r][col] = grid[r][col - 1]
                grid[r][c] = "."
                break
            c2 += 1

    elif move == "^":
        affected = get_affected(grid, (r,c), True, [])
        new_grid = [row[:] for row in grid]
        for r,c in affected:
            new_grid[r][c] = "."
        for r,c in affected:
            new_grid[r-1][c] = grid[r][c]
        return new_grid

    elif move == "v":
        affected = get_affected(grid, (r,c), False, [])
        new_grid = [row[:] for row in grid]
        for r,c in affected:
            new_grid[r][c] = "."
        for r,c in affected:
            new_grid[r+1][c] = grid[r][c]
        return new_grid

    return grid

def get_affected(grid, pos, above, affected):
    affected.append(pos)
    r,c = pos
    target_r = r-1 if above else r+1
    if grid[target_r][c] == "#":
        return []
    elif grid[target_r][c] not in affected and grid[target_r][c] == "[":
        affected = get_affected(grid, (target_r, c), above, affected)
        if affected == []:
            return  []
        affected = get_affected(grid, (target_r, c+1), above, affected)
        if affected == []:
            return  []
    elif grid[target_r][c] not in affected and grid[target_r][c] == "]":
        affected = get_affected(grid, (target_r, c), above, affected)
        if affected == []:
            return  []
        affected = get_affected(grid, (target_r, c-1), above, affected)
        if affected == []:
            return  []
    return affected

for move in moves:
    grid_2 = update_grid_2(grid_2, move)
    # print("\n".join(["".join(line) for line in grid_2]))

p2 = 0
for r, row in enumerate(grid_2):
    for c, col in enumerate(row):
        if col == "[":
            p2 += r * 100 + c

print(p2)
