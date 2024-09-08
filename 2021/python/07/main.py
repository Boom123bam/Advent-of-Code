def get_fuel_cost(pos_to, use_fib):
    cost = 0
    for pos in nums:
        if use_fib:
            cost += fib(abs(pos_to - pos))  # part 2
        else:
            cost += abs(pos_to - pos)  # part 1
    return cost


def fib(n):
    sum = 0
    for i in range(n + 1):
        sum += i
    return sum


def get_min_cost(part_1):
    min_cost = 999999999999
    min_num = min(nums)
    max_num = max(nums)
    for pos in range(min_num, max_num):
        # print(pos, get_fuel_cost(pos))
        min_cost = min(min_cost, get_fuel_cost(pos, use_fib=not part_1))

    return min_cost


input = open("07/input.txt").read()
# input = open("07/test.txt").read()
nums = list(map(int, input.split(",")))
print(get_min_cost(part_1=True))
print(get_min_cost(part_1=False))
