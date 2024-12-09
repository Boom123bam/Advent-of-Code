file = open("03/input.txt")
# file = open("03/test.txt")
input = file.read()
lines = input.split("\n")
if not lines[-1]:
    lines.pop()

do = True

def readNext(line: str):
    global do
    try:
        while len(line):
            if line.startswith("do()"):
                do = True
                line = line[4:]
            elif line.startswith("don't()"):
                do = False
                line = line[7:]
            elif line.startswith("mul("):
                break;
            else:
                line = line[1:]
        line = line[4:]

        first, second = 0,0

        if not line[0].isdigit():
            return line, None, None
        while line[0].isdigit():
            first = first * 10 + int(line[0])
            line = line[1:]

        if line[0] != "," :
            return line, None, None
        line = line[1:]

        if not line[0].isdigit():
            return line, None, None
        while line[0].isdigit():
            second = second * 10 + int(line[0])
            line = line[1:]

        if line[0] != ")" :
            return line, None, None

        return line[1:], first, second
    except IndexError:
        return line, None, None

p1 = 0
p2 = 0
for line in lines:
    while len(line):
        line, a, b = readNext(line)
        if a != None and b != None:
            p1 += a*b
            if do:
                p2 += a*b
print(p1)
print(p2)
