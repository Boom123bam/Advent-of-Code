f = open("01/input.txt", "r").read()
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()

nums = [int(x) for x in lines]
# nums = [+1, -2, +3, +1]
print(sum(nums))

freq = 0
seen_freqs = [0]

i = 0
while True:
    freq += nums[i]
    if freq in seen_freqs:
        break
    seen_freqs.append(freq)
    i += 1
    if i == len(nums):
        i = 0

print(freq)
