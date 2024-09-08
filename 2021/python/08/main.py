def find_3(nums, one):
    # find num of len 5 that contains both chars in one
    for num_index in range(len(nums)):
        num = nums[num_index]
        if len(num) != 5:
            continue
        if one[0] not in num:
            continue
        if one[1] not in num:
            continue
        return nums.pop(num_index)


def find_6(nums, one):
    # find num of len 6 that does not contain both chars in one
    for num_index in range(len(nums)):
        num = nums[num_index]
        if len(num) != 6:
            continue
        if one[0] not in num or one[1] not in num:
            return nums.pop(num_index)


def find_9(nums, three, four):
    chars = list(three)
    for char in four:
        if char not in chars:
            chars.append(char)
    nine = "".join(sorted(chars))
    nums.remove(nine)
    return nine


def find_5(nums, four):
    for num_index in range(len(nums)):
        not_common = 0
        for letter in nums[num_index]:
            if letter not in four:
                not_common += 1
        if not_common == 2:
            return nums.pop(num_index)


def get_ordered_nums(nums_left):
    ordered_nums = [None] * 10
    ordered_nums[1] = nums_left.pop(0)
    ordered_nums[7] = nums_left.pop(0)
    ordered_nums[4] = nums_left.pop(0)
    ordered_nums[8] = nums_left.pop(-1)

    ordered_nums[3] = find_3(nums_left, ordered_nums[1])
    ordered_nums[6] = find_6(nums_left, ordered_nums[1])
    ordered_nums[9] = find_9(nums_left, ordered_nums[3], ordered_nums[4])

    # 0, 2, 5 left
    ordered_nums[0] = nums_left.pop(-1)
    ordered_nums[5] = find_5(nums_left, ordered_nums[4])
    ordered_nums[2] = nums_left.pop(0)

    return ordered_nums


input_str = open("08/input.txt").read()
# input_str = open("08/test.txt").read()
if input_str[-1] == "\n":
    input_str = input_str[:-1]
lines = input_str.split("\n")


# read input
inputs = []
outputs = []

for line in lines:
    [first, second] = line.split(" | ")

    input = list(map(lambda str: "".join(sorted(str)), first.split(" ")))
    input.sort(key=len)
    inputs.append(input)

    output = list(map(lambda str: "".join(sorted(str)), second.split(" ")))
    outputs.append(output)

# part 1
unique_count = 0
unique_segment_counts = [2, 4, 3, 7]  # 1, 4, 7, 8

for digits in outputs:
    for digit in digits:
        if len(digit) in unique_segment_counts:
            unique_count += 1

print(unique_count)

# part 2
sum = 0

for row_index in range(len(inputs)):
    order = get_ordered_nums(inputs[row_index])
    output = 0
    for num in outputs[row_index]:
        output = output * 10 + order.index(num)
    sum += output

print(sum)
