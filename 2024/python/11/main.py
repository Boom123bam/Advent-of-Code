file = open("11/input.txt")
# file = open("11/test.txt")
input = file.read().strip()

nums = [int(n) for n in input.split(" ")]

rocks = {}
for n in nums:
    if n not in rocks:
        rocks[n] = 0
    rocks[n] += 1

def blink(rocks):
    new = {}
    for k, v in rocks.items():
        if k == 0:
            if 1 not in new:
                new[1] = 0
            new[1] += v
            continue

        numlen = len(str(k))
        if numlen % 2 == 0:
            if int(str(k)[:numlen//2]) not in new:
                new[int(str(k)[:numlen//2])] = 0
            new[int(str(k)[:numlen//2])] += v

            if int(str(k)[numlen//2:]) not in new:
                new[int(str(k)[numlen//2:])] = 0
            new[int(str(k)[numlen//2:])] += v
            continue

        if k*2024 not in new:
            new[k*2024] = 0
        new[k*2024] += v
    return new


for i in range(25):
    rocks = blink(rocks)
c = 0
for _,v in rocks.items():
    c += v
print(c)

for i in range(50):
    rocks = blink(rocks)
c = 0
for _,v in rocks.items():
    c += v
print(c)
