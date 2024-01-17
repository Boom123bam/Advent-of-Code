def change_directory(cd_to):
    global current_path
    if cd_to == "/":
        current_path = []
    elif cd_to == "..":
        current_path.pop(-1)
    else:
        current_path.append(cd_to)


def add_file(line):
    current_dir_dict = get_dir_dict(current_path)
    [meta_data, name] = line.split(" ")
    if meta_data == "dir":
        # add directory
        current_dir_dict.update({name: {}})
    else:
        # add size and name
        current_dir_dict.update({name: int(meta_data)})


def get_dir_dict(path):
    result = files
    for cd_to in path:
        result = result.get(cd_to)
    return result


def get_dir_size(dir_dict):
    global small_sizes_total
    keys = dir_dict.keys()
    size = 0
    for key in keys:
        item = dir_dict[key]
        if isinstance(item, dict):
            dirsize = get_dir_size(item)
            # print(f"size of {key}: {dirsize}")
            if dirsize <= 100000:
                small_sizes_total += dirsize
            size += dirsize
        else:
            size += item
    dir_sizes.append(size)
    return size


# input = open("6/test.txt").read()
input = open("6/input.txt").read()
if input[-1] == "\n":
    input = input[:-1]
lines = input.split("\n")

files = {}
current_path = []
small_sizes_total = 0
dir_sizes = []

# read files
line_index = 0
for line in lines:
    if line.startswith("$ cd"):
        cd_to = line.split(" ")[2]
        change_directory(cd_to)
    elif line == "$ ls":
        continue
    else:
        # add file to files
        add_file(line)

used_space = get_dir_size(files)
print(f"part 1: {small_sizes_total}")
remaining_space = 70000000 - used_space
min_space_to_free = 30000000 - remaining_space
dir_sizes.sort()

for size in dir_sizes:
    if size >= min_space_to_free:
        print(f"part 2: {size}")
        break
