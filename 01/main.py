file = open("1/input.txt")
input = file.read()
lines = input.split()

p1_count = 0

for i in range(1, len(lines)):
    if int(lines[i]) > int(lines[i - 1]):
        p1_count += 1


p2_count = 0
prev_sum = 9999999
sum = 0

for i in range(2, len(lines)):
    sum = int(lines[i]) + int(lines[i - 1]) + int(lines[i - 2])
    if sum > prev_sum:
        p2_count += 1
    prev_sum = sum


print(p1_count)
print(p2_count)
