def get_fuel_cost(pos_to):
    cost = 0
    for pos in nums:
        # cost += abs(pos_to - pos)  # part 1
        cost += fib(abs(pos_to - pos))  # part 2
    return cost


def fib(n):
    sum = 0
    for i in range(n + 1):
        sum += i
    return sum


# input = open("7/test.txt").read()
input = open("7/input.txt").read()
nums = list(map(int, input.split(",")))


min_cost = 999999999999
min_num = min(nums)
max_num = max(nums)
for pos in range(min_num, max_num):
    # print(pos, get_fuel_cost(pos))
    min_cost = min(min_cost, get_fuel_cost(pos))

print(min_cost)
