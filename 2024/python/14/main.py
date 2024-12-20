file = open("14/input.txt")
# file = open("14/test.txt")
input = file.read().strip()
# w,h = 11, 7
w,h = 101, 103
cycles = 100

def parse(line):
    parts = [part[2:] for part in line.split(" ")]
    p = int(parts[0].split(",")[0]) + 1j * int(parts[0].split(",")[1])
    v = int(parts[1].split(",")[0]) + 1j * int(parts[1].split(",")[1])
    return p,v

bots = [parse(line) for line in input.split("\n")]
quads = {q:0 for q in [(0,0),(0,1),(1,0),(1,1)]}

def get_end_pos(p, v, cycles):
    end = p + v * cycles
    end = end.real % w + 1j*(end.imag % h)
    return end

for p,v in bots:
    end = get_end_pos(p, v, cycles)
    if end.real == (w-1)/2 or end.imag == (h-1)/2:
        continue
    q_x = 1
    q_y = 1
    if end.real < (w-1)/2:
        q_x = 0
    if end.imag < (h-1)/2:
        q_y = 0
    quads[(q_x,q_y)] += 1

p1 = 1
for _,v in quads.items():
    p1 *= v
print(p1)

def get_pic(cycle):
    pic = [[0 for i in range(w)] for j in range(h)]
    for p, v in bots:
        end = get_end_pos(p, v, cycle)
        pic[int(end.imag)][int(end.real)] += 1
        if pic[int(end.imag)][int(end.real)] > 1:
            return None
    return pic

cycle = 0
pic = None
while pic == None:
    pic = get_pic(cycle)
    cycle += 1

print(cycle - 1)
print("\n".join(["".join(["#" if n else "." for n in line]) for line in pic]))
