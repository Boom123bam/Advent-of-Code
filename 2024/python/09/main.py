file = open("09/input.txt")
# file = open("09/test.txt")
input = file.read().strip()
id = 0
disk = []

for i, num in enumerate(input):
    if i % 2 == 0:
        disk += [id] * int(num)
        id += 1
    else:
        disk += [None] * int(num)

last_id = id - 1

first_free = disk.index(None)
last_taken = len(disk) - 1
while disk[last_taken] == None:
    last_taken -= 1

while last_taken != first_free - 1:
    disk[first_free], disk[last_taken] = disk[last_taken], None
    first_free = disk.index(None)
    while disk[last_taken] == None:
        last_taken -= 1

p1 = 0
for i,num in enumerate(disk[:first_free]):
    p1 += i * num
print(p1)

# print("".join([str(i) if i != None else "." for i in disk]))

id = 0
disk = []

for i, num in enumerate(input):
    if i % 2 == 0:
        disk.append((id,int(num)))
        id += 1
    else:
        disk.append((None,int(num)))


def attempt_move(disk, id):
    from_pos = 0
    from_amount = 0
    for pos, tup in enumerate(disk):
        if tup[0] == id:
            from_amount = tup[1]
            from_pos = pos
            break

    for pos, tup in enumerate(disk[:from_pos]):
        to_id, to_amount = tup
        if to_id == None and to_amount >= from_amount:
            tmp = disk.pop(from_pos)
            disk.insert(from_pos, (None, from_amount))
            disk[pos] = (None, to_amount - from_amount)
            disk.insert(pos, (tmp))
            return disk
    return disk

def pdisk(disk):
    r = ""
    for id,amount in disk:
        r += (str(id) if id != None else ".") * amount
    print(r)

for i in range(last_id, 0, -1):
    disk = attempt_move(disk, i)
    # pdisk(disk)

p2 = 0
i = 0
for tup in disk:
    id, amount = tup
    if id == None:
        i += amount
        continue
    for _ in range(amount):
        p2 += i * id
        i += 1

print(p2)
