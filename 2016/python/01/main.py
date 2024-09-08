input = "R4, R5, L5, L5, L3, R2, R1, R1, L5, R5, R2, L1, L3, L4, R3, L1, L1, R2, R3, R3, R1, L3, L5, R3, R1, L1, R1, R2, L1, L4, L5, R4, R2, L192, R5, L2, R53, R1, L5, R73, R5, L5, R186, L3, L2, R1, R3, L3, L3, R1, L4, L2, R3, L5, R4, R3, R1, L1, R5, R2, R1, R1, R1, R3, R2, L1, R5, R1, L5, R2, L2, L4, R3, L1, R4, L5, R4, R3, L5, L3, R4, R2, L5, L5, R2, R3, R5, R4, R2, R1, L1, L5, L2, L3, L4, L5, L4, L5, L1, R3, R4, R5, R3, L5, L4, L3, L1, L4, R2, R5, R5, R4, L2, L4, R3, R1, L2, R5, L5, R1, R1, L1, L5, L5, L2, L1, R5, R2, L4, L1, R4, R3, L3, R1, R5, L1, L4, R2, L3, R5, R3, R1, L3"
# input = "R8, R4, R4, R8"
instructions = input.split(", ")

directions = [[0, 1], [1, 0], [0, -1], [-1, 0]]
current_dir = 0
pos = [0, 0]

visited = []
first_visited_twice = None
for instruction in instructions:
    direction, amount = instruction[0], int(instruction[1:])

    if direction == "L":
        current_dir -= 1
    else:
        current_dir += 1

    if current_dir == 4 or current_dir == -4:
        current_dir = 0

    for i in range(amount):
        pos[0] += directions[current_dir][0]
        pos[1] += directions[current_dir][1]
        if first_visited_twice is None and pos in visited:
            first_visited_twice = pos.copy()
        visited.append(pos.copy())

print(abs(pos[0]) + abs(pos[1]))
print(abs(first_visited_twice[0]) + abs(first_visited_twice[1]))
