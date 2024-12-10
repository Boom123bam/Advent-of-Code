file = open("07/input.txt")
# file = open("07/test.txt")
input = file.read().strip()
lines = [(int(line.split(": ")[0]), [int(n) for n in line.split(": ")[1].split(" ")]) for line in input.split("\n")]


def get_perms(len, p1):
    def perm_next(perms):
        res = []
        for p in perms:
            if p1:
                res += [p+['+']] + [p+['*']]
            else:
                res += [p+['+']] + [p+['*']] + [p+['|']]
        return res
    res = [[]]
    for i in range(len):
        res = perm_next(res)
    return res

def eq_is_valid(val, nums, p1):
    perms = get_perms(len(nums) - 1, p1)
    for perm in perms:
        if calc_result(nums,perm) == val:
            return True
    return False

def calc_result(nums, perm):
    result = nums[0]
    for i in range(len(nums) - 1):
        if perm[i] == '+':
            result += nums[i + 1]
        elif perm[i] == '*':
            result *= nums[i + 1]
        else:
            result = int(str(result) + str(nums[i+1]))
    return result

p1 = 0
for val, nums in lines:
    if eq_is_valid(val, nums, True):
        p1 += val
print(p1)

p2 = 0
for val, nums in lines:
    if eq_is_valid(val, nums, False):
        p2 += val
print(p2)
