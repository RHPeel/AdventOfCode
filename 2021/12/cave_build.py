#DCS Day 12

def fileimport(input_file1):
    input1 = open(input_file1, 'r')
    input1 = input1.read()
    return input1

def listmaker(a):
    b = a.split('\n')
    b.remove('')
    return b

class Cave:
    def __init__(self,a):
        self.id = a
        self.is_small = ''
        self.adjacent_nodes = set()

    def determine_size(self):
        if self.id.islower() == True:
            self.is_small = True 
        else:
            self.is_small = False

    def find_adjacent_caves(self, initial_input):
        for item in initial_input:
            if item[0] == self.id:
                self.adjacent_nodes.add(item[1])
            elif item[1] == self.id:
                self.adjacent_nodes.add(item[0])
        
data_input = fileimport('day_12_input_test.txt')
input_list = listmaker(data_input)

def get_list_of_caves(a):
    b = []
    for item in a:
        c = item.split('-')
        b.append(c)
    return b

def flatten_list(a):
    e = [cell for row in a for cell in row]
    f = []
    for item in e:
        if item not in f:
            f.append(item)
    return f

connections = get_list_of_caves(input_list)
my_list = flatten_list(connections)

def get_cave_map():
    cave_dict = {}

    for i in range(0,len(my_list)):
        cave_dict[my_list[i]] = Cave(my_list[i])
        cave_dict[my_list[i]].determine_size()
        cave_dict[my_list[i]].find_adjacent_caves(connections)
    return cave_dict
