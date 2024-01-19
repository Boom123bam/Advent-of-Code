file = open("02/input.txt")
# file = open("02/test.txt")
input = file.read()
lines = input.split("\n")


def get_ans(part_1):
    horizontal = 0
    depth = 0
    aim = 0
    for line in lines:
        if not line:
            break
        [direction, count] = line.split(" ")
        count = int(count)

        if direction == "forward":
            horizontal += count
            depth += aim * count
        elif direction == "down":
            if part_1:
                depth += count  # p1
            else:
                aim += count  # p2
        else:
            if part_1:
                depth -= count  # p1
            else:
                aim -= count  # p2

    return horizontal * depth


print(get_ans(part_1=True))
print(get_ans(part_1=False))
