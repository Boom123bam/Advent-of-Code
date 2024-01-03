#read input
the_file = open("./day-1/input.txt", "r")
the_input = the_file.read()

the_file.close()

the_array = the_input.split("\n\n")

the_sims=[]

for each_thing in the_array:
    the_thing_list = each_thing.split("\n")
    the_sim = 0
    for each_num in the_thing_list:
        the_sim += int(each_num)
    
    the_sims.append(the_sim)
    
the_sims.sort()
the_sims.reverse()
print(the_sims[0])
print(the_sims[0] + the_sims[1] + the_sims[2])
