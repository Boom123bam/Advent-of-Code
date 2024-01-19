file = open("03/input.txt")
# file = open("03/test.txt")
input = file.read()
lines = input.split("\n")

length = len(lines[0])

if not lines[-1]:
    lines.pop(-1)


# part 1

sum = [0] * length

for line in lines:
    for index, char in enumerate(line):
        sum[index] += 1 if char == "1" else -1

res1 = 0
for num in sum:
    res1 = res1 * 2 + (1 if num > 0 else 0)

res2 = 2 ** len(lines[0]) - res1 - 1

print(res1 * res2)


# part 2


def findMostCommon(lines, index):
    w = 0
    for line in lines:
        if not line:
            break
        w += 1 if line[index] == "1" else -1
    return "1" if w >= 0 else "0"
    # return w > 0


def to_binary(binary_string):
    result = 0
    for char in binary_string:
        result = result * 2 + (1 if char == "1" else 0)
    return result


oxygen = lines.copy()
co2 = lines.copy()

for i in range(length):
    most_common_oxygen = findMostCommon(oxygen, i)
    most_common_co2 = findMostCommon(co2, i)
    if len(oxygen) > 1:
        oxygen = [line for line in oxygen if line[i] == most_common_oxygen]
    if len(co2) > 1:
        co2 = [line for line in co2 if line[i] != most_common_co2]
    # print(most_common, len(oxygen), len(co2))

# print(oxygen, co2)
print(to_binary(oxygen[0]) * to_binary(co2[0]))
