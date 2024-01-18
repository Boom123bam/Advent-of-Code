def print_board(board):
    for row in board:
        row_str = ""
        for num in row:
            num_str = "##" if num == -1 else str(num)
            row_str += num_str + (3 - len(num_str)) * " "
        print(row_str)


def replace_nums(target, num_to):
    for tb_index in range(len(boards)):
        for row_index in range(len(boards[0])):
            for col_index in range(len(boards[0][0])):
                if boards[tb_index][row_index][col_index] == target:
                    boards[tb_index][row_index][col_index] = num_to


def find_winning_board_index():
    for board_index in range(len(boards)):
        if has_won(boards[board_index]):
            return board_index
    return -1


def find_remaining_board_index():
    remaining_board_index = -1
    for board_index in range(len(boards)):
        if not has_won(boards[board_index]):
            if remaining_board_index != -1:
                # more than one remain
                return -1
            remaining_board_index = board_index

    return remaining_board_index


def has_won(board):
    # check rows
    for row in board:
        if check_row(row):
            return True

    # check cols
    for col_index in range(len(board[0])):
        if check_col(board, col_index):
            return True

    return False


def check_row(row):
    for num in row:
        if num != -1:
            return False
    return True


def check_col(board, col_index):
    for row in board:
        if row[col_index] != -1:
            return False
    return True


def find_sum(board):
    sum = 0
    for row in board:
        for num in row:
            if num != -1:
                sum += num
    return sum


# read file
file = open("4/input.txt")
input = file.read()

# remove blank line
if input[-1] == "\n":
    input = input[:-1]

# parse input
sections = input.split("\n\n")

nums = list(map(int, sections.pop(0).split(",")))

boards = []
for section in sections:
    board = []
    for line in section.split("\n"):
        num_strings = line.strip().replace("  ", " ").split(" ")
        board.append(list(map(int, num_strings)))

    boards.append(board)

winning_index = -1

while winning_index == -1:
    last_num = nums.pop(0)
    replace_nums(last_num, -1)
    winning_index = find_winning_board_index()

sum = find_sum(boards[winning_index])
print(sum * last_num)

# part 2

# find remaining board index
remaining_board_index = -1

while remaining_board_index == -1:
    last_num = nums.pop(0)
    replace_nums(last_num, -1)
    remaining_board_index = find_remaining_board_index()

# play until it wins
while not has_won(boards[remaining_board_index]):
    last_num = nums.pop(0)
    replace_nums(last_num, -1)
sum = find_sum(boards[remaining_board_index])

print(sum * last_num)
