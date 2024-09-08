f = open("03/input.txt", "r").read()
# f = """#1 @ 1,3: 4x4
# #2 @ 3,1: 4x4
# #3 @ 5,5: 2x2
# """
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()


parsed_lines = []

for line in lines:
    [id, _, xy, size] = line.split(" ")
    [x, y] = [int(i) for i in xy[:-1].split(",")]
    [w, h] = [int(i) for i in size.split("x")]
    parsed_lines.append([int(id[1:]), x, y, w, h])

grid = [[0] * 1000 for _ in range(1000)]


def print_grid():
    for line in grid:
        print("".join([str(i) for i in line]))


def get_num_overlaps():
    count = 0
    for line in grid:
        for num in line:
            if num > 1:
                count += 1
    return count


id_without_overlap = -1

for data in parsed_lines:
    [id, x, y, w, h] = data
    for i in range(x, x + w):
        for j in range(y, y + h):
            grid[j][i] += 1

# print_grid(
print(get_num_overlaps())

for data in parsed_lines:
    [id, x, y, w, h] = data
    has_overlap = False

    for i in range(x, x + w):
        for j in range(y, y + h):
            if grid[j][i] > 1:
                has_overlap = True

    if not has_overlap:
        print(id)
        break
