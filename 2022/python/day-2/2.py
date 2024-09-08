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

file.close()

lines = input.split("\n")

points = 0
for round in lines:
    [opponent, outcome] = round.split(" ")
    if outcome == "X":
        # lose
        me = order[opponent]
        points += shape[me] #shape points
    elif outcome == "Y":
        # draw
        me = opponent
        points += shape[me] #shape points
        points += 3 #draw points
    else:
        # win
        me = order[order[opponent]]
        points += shape[me] #shape points
        points += 6 #draw points

print(points)