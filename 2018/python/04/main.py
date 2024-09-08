f = open("04/input.txt", "r").read()
# f = open("04/test.txt", "r").read()
lines = f.split("\n")
if lines[-1] == "":
    lines.pop()

parsed_lines = []


def printList(list):
    for line in list:
        print(line)
    print()


for line in lines:
    [date_time, action] = line[1:].split("] ")
    [date, time] = date_time.split(" ")
    [year, month, day] = [int(x) for x in date.split("-")]
    [hour, minute] = [int(x) for x in time.split(":")]

    id = None
    if action.endswith("begins shift"):
        id = int(action.split(" ")[1][1:])
        action = "begins shift"

    parsed_lines.append([year, month, day, hour, minute, action, id])

# sort by time
parsed_lines.sort()

guards_total_sleep = {}

current_id = None
for i, line in enumerate(parsed_lines):
    [year, month, day, hour, minute, action, id] = line
    if action == "begins shift":
        current_id = id
    else:
        parsed_lines[i][-1] = current_id
        if action == "falls asleep":
            [year2, month2, day2, hour2, minute2, _, _] = parsed_lines[i + 1]

            if current_id not in guards_total_sleep:
                guards_total_sleep[current_id] = 0

            guards_total_sleep[current_id] += (
                minute2 + hour2 * 60 + day2 * 60 * 24
            ) - (minute + hour * 60 + day * 60 * 24)


most_sleep_id = max(guards_total_sleep, key=guards_total_sleep.get)


target_guard_data = list(filter(lambda line: line[-1] == most_sleep_id, parsed_lines))
# sort by hour and min
target_guard_data.sort(key=lambda line: line[3] * 60 + line[4])


num_sleeping_days = 0
max_sleeping_days = 0
max_sleeping_min = None
for i, line in enumerate(target_guard_data):
    [_, _, _, hour, minute, action, _] = line

    if action == "falls asleep":
        num_sleeping_days += 1
    elif action == "wakes up":
        num_sleeping_days -= 1
    if (
        i + 1 < len(target_guard_data)
        and target_guard_data[i + 1][3] == hour
        and target_guard_data[i + 1][4] == minute
    ):
        continue
    if num_sleeping_days > max_sleeping_days:
        max_sleeping_days = num_sleeping_days
        max_sleeping_min = minute

print(most_sleep_id * max_sleeping_min)


guards_num_sleeping_days = {}
parsed_lines.sort(key=lambda line: line[3] * 60 + line[4])
max_sleeping_days = 0
max_sleeping_min = None
max_sleeping_guard_id = None


for i, line in enumerate(parsed_lines):
    [_, _, _, hour, minute, action, id] = line

    if id not in guards_num_sleeping_days:
        guards_num_sleeping_days[id] = 0

    if action == "falls asleep":
        guards_num_sleeping_days[id] += 1
    elif action == "wakes up":
        guards_num_sleeping_days[id] -= 1
    if (
        i + 1 < len(parsed_lines)
        and parsed_lines[i + 1][3] == hour
        and parsed_lines[i + 1][4] == minute
    ):
        continue
    if guards_num_sleeping_days[id] > max_sleeping_days:
        max_sleeping_days = guards_num_sleeping_days[id]
        max_sleeping_min = minute
        max_sleeping_guard_id = id

print(max_sleeping_min * max_sleeping_guard_id)
