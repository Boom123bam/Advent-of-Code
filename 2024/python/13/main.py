file = open("13/input.txt")
# file = open("13/test.txt")
input = file.read().strip()

def parse_ln(ln):
    parts = ln.split(": ")[1].split(", ")
    return int(parts[0][2:]) + 1j * int(parts[1][2:])

machines = [
    [parse_ln(line) for line in lines.split("\n")]
    for lines in input.split("\n\n")
]

def get_presses(machine):
    a,b,prize = machine
    for i in range(100):
        for j in range(100):
            pos = a * i + b * j
            if pos == prize:
                return i,j
    return None, None

p1 = 0
for machine in machines:
    a_c, b_c = get_presses(machine)
    if a_c != None and b_c != None:
        p1 += a_c * 3 + b_c

print(p1)

import math

def find_angle_between(a, b):
    dot_product = a.real * b.real + a.imag * b.imag
    mag_a = abs(a)
    mag_b = abs(b)
    cos_theta = dot_product / (mag_a * mag_b)
    angle_rad = math.acos(cos_theta)
    return angle_rad

def is_close_to_integer(num, tolerance=1e-3):
    return abs(num - round(num)) < tolerance

p2 = 0
for machine in machines:
    a,b,prize = machine
    prize += 10000000000000 + 10000000000000j
    ang_a = find_angle_between(prize, b)
    ang_b = find_angle_between(prize, a)
    a_len = (abs(prize) * math.sin(ang_a)) / math.sin(math.pi - ang_a - ang_b)
    if ang_a == 0:
        continue
    b_len = (a_len * math.sin(ang_b)) / math.sin(ang_a)
    a_c, b_c = a_len/abs(a), b_len/abs(b)
    if is_close_to_integer(a_c) and is_close_to_integer(b_c):
        p2 += round(a_c) * 3 + round(b_c)

print(p2)
