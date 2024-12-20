file = open("12/input.txt")
# file = open("12/test.txt")
input = file.read().strip()
lines = input.split("\n")
scanned = [[0 for _ in range(len(lines[0]))] for _ in range(len(lines))]

def expand_zone(r,c,res):
    res.append((r,c))
    scanned[r][c] = 1
    for nr, nc in [(r+1,c),(r-1,c),(r,c+1),(r,c-1)]:
        if (nr, nc) in res:
            continue
        if nr not in range(0,len(lines)) or nc not in range(0,len(lines[0])):
            continue
        if lines[nr][nc] == lines[r][c]:
            res = expand_zone(nr,nc,res)
    return res

zone_id = 0
zones = {}
for r,row in enumerate(lines):
    for c, col in enumerate(row):
        if not scanned[r][c]:
            zones[zone_id] = expand_zone(r,c,[])
            zone_id += 1

def get_perimeter(zone):
    total = len(zone) * 4
    for (r,c) in zone:
        for nr, nc in [(r+1,c),(r-1,c),(r,c+1),(r,c-1)]:
            if (nr,nc) in zone:
                total -= 1
    return total

p1 = 0
for _,zone in zones.items():
    p1 += len(zone) * get_perimeter(zone)
    # print(len(zone), get_perimeter(zone))

print(p1)

all_edges = []
for r,row in enumerate(lines):
    all_edges.append([])
    for c, col in enumerate(row):
        edges = []
        for nr, nc in [(r+1,c),(r-1,c),(r,c+1),(r,c-1)]:
            if nr not in range(0,len(lines)) or nc not in range(0,len(lines[0])) or lines[nr][nc] != lines[r][c]:
                edges.append(1)
            else:
                edges.append(0)
        all_edges[r].append(edges)

def count_edges(zone):
    count = 0
    seen = []
    for r, c in zone:
        seen.append((r,c))
        p_edges = all_edges[r][c]
        count += p_edges.count(1)
        for n_r, n_c in [(r+1,c),(r-1,c),(r,c+1),(r,c-1)]:
            if (n_r, n_c) not in seen:
                continue
            n_edges = all_edges[n_r][n_c]
            for i in range(4):
                if p_edges[i] and n_edges[i]:
                    count -= 1
    return count


p2 = 0
for _,zone in zones.items():
    p2 += len(zone) * count_edges(zone)

print(p2)
