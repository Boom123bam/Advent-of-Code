file = open("10/input.txt")
# file = open("10/test.txt")
input = file.read().strip()
lines = input.split("\n")
grid = [[int(char) if char != "." else None for char in line] for line in lines]

def pg():
    print("\n".join(["".join([str(num) for num in line]) for line in grid]))


def find_trails(r,c,set=None):
    res = 0
    target = grid[r][c] + 1
    for r_off, c_off in [(0,1),(1,0),(0,-1),(-1,0)]:
        n_r, n_c = r + r_off, c + c_off
        if n_r not in range(0, len(grid)) or n_c not in range(0, len(grid[0])):
            continue
        if grid[n_r][n_c] == target:
            if target == 9:
                res += 1
                if set != None:
                    set[(n_r, n_c)] = True
            else:
                res += find_trails(n_r,n_c,set)
    return res if set == None else len(set)

p1 = 0
p2 = 0
for r, row in enumerate(grid):
    for c, col in enumerate(row):
        if col == 0:
            p1 += find_trails(r,c,{})
            p2 += find_trails(r,c)

print(p1)
print(p2)
