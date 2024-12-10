file = open("05/input.txt")
# file = open("05/test.txt")
input = file.read().strip()
parts = input.split("\n\n")

rules = [[int(num) for num in nums.split("|")] for nums in parts[0].split("\n")]
updates = [[int(num) for num in nums.split(",")] for nums in parts[1].split("\n")]


rulesMap = {}
for rule in rules:
    if rule[1] not in rulesMap:
        rulesMap[rule[1]] = []
    rulesMap[rule[1]].append(rule[0])

def update_is_correct(nums):
    for i, num in enumerate(nums):
        if i == len(nums) - 1:
            continue
        if num not in rulesMap:
            continue
        not_allowed_after = rulesMap[num]
        for j in range(i+1, len(nums)):
            if nums[j] in not_allowed_after:
                temp = nums.pop(j)
                nums.insert(i, temp)
                return False, nums
    return True, nums

p1 = 0
p2 = 0

for update in updates:
    correct, new = update_is_correct(update)
    if correct:
        p1 += update[int((len(update)-1)/2)]
    else:
        while not correct:
            correct, new = update_is_correct(update)
        p2 += new[int((len(new)-1)/2)]

print(p1)
print(p2)
