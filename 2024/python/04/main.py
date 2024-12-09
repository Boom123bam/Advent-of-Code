file = open("04/input.txt")
# file = open("04/test.txt")
input = file.read()
lines = input.split("\n")
if not lines[-1]:
    lines.pop()

def count(line):
    return line.count("XMAS") + line.count("SAMX")

p1 = 0
for line in lines:
    p1 += count(line)

for i in range(len(lines[0])):
    col = ""
    for line in lines:
        col += line[i]
    p1 += count(col)

# diag
def get_diag(lines, i):
    row = 0
    col = i
    result = ""
    while col >= 0:
        if row < len(lines) and col < len(lines[row]):
            result += lines[row][col]
        col -= 1
        row += 1
    return result

def get_diag_2(lines, i):
    row = 0
    col = len(lines[0]) - i - 1
    result = ""
    while col < len(lines[0]):
        if row < len(lines) and col >= 0:
            result += lines[row][col]
        col += 1
        row += 1
    return result

for i in range(len(lines) + len(lines[0])):
    p1 += count(get_diag(lines, i))
    p1 += count(get_diag_2(lines, i))

print(p1)

def is_x_mas(grid, row, col):
    if grid[row][col] != "A":
        return False
    if row < 1 or row > len(grid)-2:
        return False
    if col < 1 or col > len(grid[0])-2:
        return False
    diag_1 = grid[row-1][col-1] + grid[row+1][col+1]
    diag_2 = grid[row+1][col-1] + grid[row-1][col+1]
    if "M" not in diag_1 or "S" not in diag_1:
        return False
    if "M" not in diag_2 or "S" not in diag_2:
        return False
    return True

p2 = 0
for row in range(len(lines)):
    for col in range(len(lines[row])):
        if is_x_mas(lines, row, col):
            p2 += 1

print(p2)
