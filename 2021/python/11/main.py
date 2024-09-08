def to_num_arr(nums_str):
    return list(map(int, list(nums_str)))


def increment_all():
    for y in range(len(octopus_grid)):
        octopus_grid[y] = [x + 1 for x in octopus_grid[y]]


def find_flash_pos():
    for y in range(len(octopus_grid)):
        for x in range(len(octopus_grid)):
            if octopus_grid[y][x] >= 10:
                return (y, x)
    return None


def increment_positive_neighbours(yx):
    # include 0
    neighbours = find_neighbours(yx)
    for y, x in neighbours:
        if octopus_grid[y][x] >= 0:
            octopus_grid[y][x] += 1


def find_neighbours(yx):
    # includes self
    [y, x] = yx
    neighbours = []
    for y_off in range(-1, 2):
        for x_off in range(-1, 2):
            neighbour_y = y + y_off
            neighbour_x = x + x_off
            if neighbour_x < 0 or neighbour_y < 0:
                continue
            if neighbour_x >= 10 or neighbour_y >= 10:
                continue
            neighbours.append([neighbour_y, neighbour_x])
    return neighbours


def reset_flashes():
    for y in range(len(octopus_grid)):
        for x in range(len(octopus_grid)):
            if octopus_grid[y][x] == -1:
                octopus_grid[y][x] = 0


def print_grid():
    for row in octopus_grid:
        row_str = ""
        for num in row:
            row_str += "█" if num == -1 else str(num)
        print(row_str)
    print()


def step():
    global flash_count
    increment_all()
    flash_pos = find_flash_pos()
    while flash_pos:
        flash_count += 1
        octopus_grid[flash_pos[0]][flash_pos[1]] = -1
        increment_positive_neighbours(flash_pos)
        flash_pos = find_flash_pos()
    # print_grid()
    reset_flashes()


input = open("11/input.txt").read()
# input = open("11/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")

flash_count = 0
octopus_grid = list(map(to_num_arr, lines))
prev_flash_count = 0


i = 0
while True:
    i += 1
    step()
    if flash_count - prev_flash_count == 100:
        print(i)
        break
    if i == 100:
        print(flash_count)
    prev_flash_count = flash_count
