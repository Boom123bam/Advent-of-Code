input = open("09/input.txt").read()
# input = open("09/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")


def line_to_num_list(line):
    return list(map(int, list(line)))


def yx_exists(yx):
    [row_i, col_i] = yx
    if row_i < 0 or col_i < 0:
        return False
    return row_i < len(height_matrix) and col_i < len(height_matrix[0])


def is_low_point(yx):
    height = height_matrix[y][x]
    neighbours = get_neighbours(yx)
    for neighbour in neighbours:
        if yx_exists(neighbour):
            if height >= height_matrix[neighbour[0]][neighbour[1]]:
                return False
    return True


def find_lower_neighbours(yx):
    height = height_matrix[yx[0]][yx[1]]
    neighbours = get_neighbours(yx)
    lower_yxs = []
    for neighbour in neighbours:
        neighbour_height = height_matrix[neighbour[0]][neighbour[1]]

        if neighbour_height < height:
            lower_yxs.append(neighbour)

    return lower_yxs


def get_neighbours(yx):
    offsets = [[1, 0], [0, 1], [-1, 0], [0, -1]]
    neighbours = []
    for offset in offsets:
        neighbour = [yx[0] + offset[0], yx[1] + offset[1]]
        if yx_exists(neighbour):
            neighbours.append(neighbour)
    return neighbours


def find_basin_size(low_point_yx):
    points_in_basin = []
    points_in_basin.append(low_point_yx)
    while expand_basin(points_in_basin):
        # print_basin(points_in_basin)
        pass
    return len(points_in_basin)


def print_basin(points_in_basin):
    for y in range(len(height_matrix)):
        for x in range(len(height_matrix[0])):
            print("█" if [y, x] in points_in_basin else height_matrix[y][x], end="")
        print()
    print()


def expand_basin(points_in_basin):
    points_to_add = []
    for yx in points_in_basin:
        neighbours = get_neighbours(yx)
        for neighbour in neighbours:
            if neighbour in points_in_basin:
                continue
            neighbour_lower_neighbours = find_lower_neighbours(neighbour)
            add_point = True
            for neighbour_lower_neighbour in neighbour_lower_neighbours:
                if neighbour_lower_neighbour not in points_in_basin:
                    add_point = False
                    break
            if (
                add_point
                and neighbour not in points_to_add
                and height_matrix[neighbour[0]][neighbour[1]] != 9
            ):
                points_to_add.append(neighbour)
    points_in_basin.extend(points_to_add)
    return len(points_to_add)


height_matrix = list(map(line_to_num_list, lines))

p1_sum = 0
low_points = []

for y in range(len(height_matrix)):
    for x in range(len(height_matrix[0])):
        if is_low_point([y, x]):
            low_points.append([y, x])
            p1_sum += height_matrix[y][x] + 1

print(p1_sum)

basin_sizes = []

for low_point in low_points:
    size = find_basin_size(low_point)
    basin_sizes.append(size)

basin_sizes.sort(reverse=True)

print(basin_sizes[0] * basin_sizes[1] * basin_sizes[2])
