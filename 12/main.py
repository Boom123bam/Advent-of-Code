input = open("12/input.txt").read()
# input = open("12/test.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")

# each node is a set
nodes = {}

# add nodes and connections
for line in lines:
    [a, b] = line.split("-")
    if a in nodes:
        nodes[a].add(b)
    else:
        nodes[a] = {b}

    if b in nodes:
        nodes[b].add(a)
    else:
        nodes[b] = {a}


def find_possible_paths(
    part_1, currentNode="start", path=[], small_cave_twice=None, depth=0
):
    paths = 0
    neighbours = nodes[currentNode]
    for neighbour in neighbours:
        path_small_cave_twice = small_cave_twice
        if neighbour == "start":
            continue
        if neighbour.islower() and neighbour in path:
            if part_1:
                continue
            if path_small_cave_twice:
                if path_small_cave_twice != neighbour:
                    continue
                # neighbour  == path_small_cave_twice
                if path.count(neighbour) == 2:
                    continue
            else:
                path_small_cave_twice = neighbour

        if neighbour == "end":
            # print("-" * depth, neighbour, path, path_small_cave_twice)
            paths += 1
            continue
        path_copy = path.copy()
        path_copy.append(neighbour)
        paths += find_possible_paths(
            part_1, neighbour, path_copy, path_small_cave_twice, depth + 1
        )
    return paths


print(find_possible_paths(part_1=True))
print(find_possible_paths(part_1=False))
