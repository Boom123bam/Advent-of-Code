def mutate_count(num_count):
    new_count = [0] * 9
    for num, count in enumerate(num_count):
        if num == 0:
            new_count[6] += count
            new_count[8] += count
        else:
            new_count[num - 1] += count
    return new_count


# input = open("6/test.txt").read()
input = open("6/input.txt").read()
if input[-1] == "\n":
    input = input[:-1]

nums = list(map(int, input.split(",")))
num_count = [0] * 9

for num in nums:
    num_count[num] += 1


for i in range(256):
    num_count = mutate_count(num_count)

total = 0
for count in num_count:
    total += count
print(total)
