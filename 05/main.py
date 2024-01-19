def print_count_map():
    for row in count_map:
        row_str = ""
        for count in row:
            row_str += str(count) if count else "."
        print(row_str)
    print()


def sign(num):
    if num > 0:
        return 1
    elif num < 0:
        return -1
    else:
        return 0


input = open("05/input.txt").read()
# input = open("05/test.txt").read()
input_lines = input.split("\n")
if not input_lines[-1]:
    input_lines = input_lines[:-1]

# read lines
lines = []
max_x = 0
max_y = 0
for input_line in input_lines:
    [x1, y1, x2, y2] = list(map(int, (input_line.replace(" -> ", ",").split(","))))
    lines.append((x1, y1, x2, y2))
    max_x = max(max_x, x1, x2)
    max_y = max(max_y, y1, y2)


def get_overlaps(part_1):
    count_map = [[0] * (max_x + 1) for _ in range(max_y + 1)]

    for line in lines:
        [x1, y1, x2, y2] = line
        # uncomment for part 1
        if part_1 and not (x1 == x2 or y1 == y2):
            continue
        x_sign = sign(x2 - x1)
        y_sign = sign(y2 - y1)
        x = x1
        y = y1

        while x != x2 or y != y2:
            count_map[y][x] += 1
            x += x_sign
            y += y_sign
        count_map[y][x] += 1

    overlaps = 0

    for row in count_map:
        for count in row:
            if count >= 2:
                overlaps += 1
    return overlaps


print(get_overlaps(part_1=True))
print(get_overlaps(part_1=False))
