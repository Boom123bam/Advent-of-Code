test_input = "A Y\nB X\nC Z"

order = {
    "r":"s",
    "s":"p",
    "p":"r"
         }

shape = {
    "r":1,
    "p":2,
    "s":3
        }

file = open("./day-2/input.txt", "r")
input = file.read()
# input = test_input
input = input.replace("A","r")
input = input.replace("B","p")
input = input.replace("C","s")
input = input.replace("X","r")
input = input.replace("Y","p")
input = input.replace("Z","s")

file.close()

lines = input.split("\n")

points = 0
for round in lines:
    [opponent, me] = round.split(" ")
    points += shape[me]
    if (opponent == me): 
        #draw
        print(round, "draw")
        points += 3
    elif (order[me] == opponent):
        #win
        print(round, "win")
        points += 6
    else:
        #loser
        print(round, "bad")

print(points)
