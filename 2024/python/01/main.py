file = open("01/input.txt")
# file = open("01/test.txt")
input = file.read()
lines = input.split("\n")
if not lines[-1]:
    lines.pop()

left = []
right = []
for line in lines:
    nums = [int(num) for num in line.split("   ")]
    left.append(nums[0])
    right.append(nums[1])

left.sort()
right.sort()

p1 = 0
for i,l in enumerate(left):
    p1 += abs(l - right[i])
print(p1)

p2 = 0
for num in left:
    p2 += num * right.count(num)
print(p2)
