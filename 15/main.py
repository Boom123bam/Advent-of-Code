input = open("15/input.txt").read()
# input = open("15/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]


def find_neighbour_coords(yx):
    (y, x) = yx
    offsets = [(1, 0), (0, 1), (-1, 0), (0, -1)]
    neighbours = []
    for o_y, o_x in offsets:
        neighbour = (y + o_y, x + o_x)
        if coord_is_valid(neighbour):
            neighbours.append(neighbour)
    return neighbours


def coord_is_valid(yx):
    (y, x) = yx
    if y < 0 or x < 0:
        return False
    if y >= len(grid) or x >= len(grid[0]):
        return False
    return True


def dijkstra(start_yx, end_yx):
    all_yx = [(y, x) for y in range(len(grid)) for x in range(len(grid[0]))]
    # set every dist to high number to represent not calculated
    dists = {xy: 9999 for xy in all_yx}
    # set start dist to 0
    dists[start_yx] = 0
    # queue = all nodes
    queue = all_yx
    while len(queue):
        # find smallest dist in q
        min_yx = min(queue, key=dists.get)
        # for neighbour in q
        self_dist = dists[min_yx]
        for neighbour in find_neighbour_coords(min_yx):
            neighbour_dist = dists[neighbour]
            if neighbour_dist <= self_dist:
                continue
            new_dist = self_dist + grid[neighbour[0]][neighbour[1]]
            # if dist is smaller update dist
            if new_dist < dists[neighbour]:
                dists[neighbour] = new_dist
        # remove from q
        queue.remove(min_yx)
    return dists[end_yx]


def increment_wrap(num, amount):
    # increment num by amount and resets to 1 if it reaches 9
    return (num + amount - 1) % 9 + 1


def expand_horizontal(line, factor):
    res = []
    for i in range(factor):
        res.extend([increment_wrap(num, i) for num in line])
    return res


def expand_vertical(factor):
    original_size = len(grid)
    for i in range(factor - 1):
        for line_i in range(original_size):
            grid.append([increment_wrap(num, i + 1) for num in grid[line_i]])


def print_grid():
    for line in grid:
        print("".join(map(str, line)))


grid = [[int(num) for num in list(line)] for line in input.split("\n")]

print(dijkstra((0, 0), (len(grid) - 1, len(grid[0]) - 1)))

# expand horizontally and vertically
grid = [expand_horizontal(line, 5) for line in grid]
expand_vertical(5)
print(dijkstra((0, 0), (len(grid) - 1, len(grid[0]) - 1)))

# it took 15 mins lol
