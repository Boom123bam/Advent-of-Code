file = open("2/input.txt")
input = file.read()
lines = input.split("\n")

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
        # depth += count # p1
        aim += count  # p2
    else:
        # depth -= count # p1
        aim -= count  # p2

print(horizontal * depth)
