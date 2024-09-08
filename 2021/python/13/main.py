input = open("13/input.txt").read()
# input = open("13/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]


def read_fold_line(line):
    return line.split(" ")[-1][0]


def plot_dots():
    for x, y in dots:
        grid[y][x] = 1


def print_grid():
    for row in grid:
        row_str = ""
        for dot in row:
            row_str += "█" if dot else " "
        print(row_str)
    return


def add_each_item(target, addition):
    for i in range(len(target)):
        target[i] += addition[i]
    return


def fold_grid(along_x_axis):
    if along_x_axis:
        for y_i in range(int((len(grid) - 1) / 2)):
            y_i = add_each_item(grid[y_i], grid.pop(-1))
        grid.pop(-1)
    else:
        for y_i in range(len(grid)):
            for x_i in range(int((len(grid[y_i]) - 1) / 2)):
                grid[y_i][x_i] += grid[y_i].pop(-1)
            grid[y_i].pop(-1)
    return


def count_dots():
    count = 0
    for row in grid:
        for num in row:
            if num:
                count += 1
    return count


[dots_lines, folds_lines] = [section.split("\n") for section in input.split("\n\n")]
dots = [[int(num) for num in line.split(",")] for line in dots_lines]
folds = [read_fold_line(line) for line in folds_lines]

max_x_i = max([dot[0] for dot in dots])
max_y_i = max([dot[1] for dot in dots])

grid = [[0] * (max_x_i + 1) for y in range(max_y_i + 1)]

plot_dots()

# part 1
# reversed because line x = 0 is y axis
fold_grid(along_x_axis=(folds.pop(0) == "y"))

print(count_dots())

# part 2
for fold in folds:
    fold_grid(along_x_axis=(fold == "y"))

print_grid()
