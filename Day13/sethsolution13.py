import copy

def inputParser(string):
    #takes a string of input and turns it into a useful list of digits
    string = string.strip()
    string = string.split(':')
    output = [int(digit) for digit in string]
    return output
    
def firewallBuilder(dictionary, list):
    #calls inputParser to add a list of instructions to a dictionary
    dictionary[list[0]] = [] #the value will be a list of zeros
    for i in range(list[1]):
        dictionary[list[0]].append(0) #build out the depth of each layer
    dictionary[list[0]][0] = 1 #here's the location of the security scanner
    flag = True
    dictionary[list[0]] = [dictionary[list[0]],flag] #appends a boolean for direction
    return dictionary

def firewallFinisher(dictionary):
    #makes sure that the firewall is padded with empty values
    checker = set([i for i in range(max(firewall.keys()))])
    checkAgainst = set(firewall.keys())
    for i in checker.difference(checkAgainst):
#        print(i)
        dictionary[i] = 'skip'
#        print(dictionary[i])
#    print(dictionary)
    return dictionary

def wallUpdater(dictionary):
    #updates the firewall after the scanners move, returns the new dict
    for layer in dictionary:
        if not dictionary[layer] == 'skip':
            position = dictionary[layer][0].index(1) #mark where the scanner is
            dictionary[layer][0][position] = 0
            if dictionary[layer][1]: #if the flag is on, move forward
                dictionary[layer][0][position + 1] = 1
                if dictionary[layer][0][-1] == 1:
                    dictionary[layer][1] = False
            else:
                dictionary[layer][0][position - 1] = 1
                if dictionary[layer][0][0] == 1:
                    dictionary[layer][1] = True
    return dictionary

def severity(dictionary):
    #given a firewall, returns the "severity" of attempting to cross
    hypotheticalDictionary = copy.deepcopy(dictionary)
    severity = 0
    for i in range(len(hypotheticalDictionary)): #iterate for as many layers as there are in the firewall
#        print('now at location', i, 'layer is', hypotheticalDictionary[i][0])
        if hypotheticalDictionary[i][0][0] == 1:
#            print('encountered security, adding', i * len(hypotheticalDictionary[i][0]))
            severity += i * len(hypotheticalDictionary[i][0])
        hypotheticalDictionary = wallUpdater(hypotheticalDictionary)
    return severity

def isCaught(dictionary):
    #true if trying to get through the firewall would get you caught
    dictCopy = copy.deepcopy(dictionary)
    for i in range(len(dictCopy)):
        if dictCopy[i][0][0] == 1:
#            print('caught by security at layer', i, 'see', dictCopy[i][0])
            return True
        dictCopy = wallUpdater(dictCopy)
    return False

f = open('sethinput.txt', 'r')
data = f.readlines()
    
firewall = {}

for instruction in data:
    firewall = firewallBuilder(firewall, inputParser(instruction))

firewallFinisher(firewall)

print('answer to part A is', severity(firewall))

delay = 0

while isCaught(firewall):
    delay += 1
    firewall = wallUpdater(firewall)
#    print('delay of', delay, 'isCaught?', isCaught(firewall))
    if delay > 100000:
        break

print('part B answer is', delay)