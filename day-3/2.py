file = open("./day-3/input2.txt", "r")
input = file.read()
lines = input.split("\n")

def find_score(char):
    if ord(char) > 96:
        return (ord(char) - 96)
    else:
        return (ord(char) - 64 + 26)

score = 0

for i in range (0, len(lines)-2, 3):
    element1 = lines[i]
    element2 = lines[i + 1]
    element3 = lines[i + 2]
    
    common = ""
    for char in element1:
        if char in element2 and char in element3:
            common = char
            
    score += find_score(common)

print(score)