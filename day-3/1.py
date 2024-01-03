# first step: find the two compartments
# second step: find same element
# third step: add to score

file = open("./day-3/input2.txt", "r")
input = file.read()
lines = input.split("\n")

def find_score(char):
    if ord(char) > 96:
        return (ord(char) - 96)
    else:
        return (ord(char) - 64 + 26)
    

score = 0
for line in lines:
    element1 = line[:int(len(line)/2)]
    element2 = line[int(len(line)/2):]
    
    common = ""
    for char in element1:
        # if (char.find(element2) != -1):
        if char in element2:
            common = char
    
    score += find_score(common)
    
print(score)