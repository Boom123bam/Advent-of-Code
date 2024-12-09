file = open("02/input.txt")
# file = open("02/test.txt")
input = file.read()
lines = input.split("\n")
if not lines[-1]:
    lines.pop()

def is_safe(nums):
    diffs = [nums[i + 1] - num for i, num in enumerate(nums[:-1])]
    increasing =  all(diff > 0 for diff in diffs)
    decreasing =  all(diff < 0 for diff in diffs)
    if not (increasing or decreasing):
        return False
    if not all(abs(diff) <= 3 for diff in diffs):
        return False
    return True

def is_safe_2(nums):
    for i in range(len(nums)):
        if is_safe(nums[:i] + nums[i+1:]):
            return True
    return False

num_rows = [[int(num) for num in line.split(" ")] for line in lines]

p1 = 0
p2 = 0
for nums in num_rows:
    if is_safe(nums):
        p1 += 1
        p2 += 1
    elif is_safe_2(nums):
        p2 += 1

print(p1)
print(p2)
