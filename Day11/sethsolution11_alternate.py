f = open('sethinput.txt', 'r')
data = f.read()
data = data.strip()
data = 'ne,ne,sw,sw'
data = data.split(',')

class hex(object):
    locations = []
    
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z
        hex.locations.append(self)
    
    def __str__(self):
        return "Hex (%s, %s, %s)"%(self.x, self.y, self.z)
    
    def ne(self):
        self.x += 1
        self.z += -1
        hex.locations.append(self)

        
    def n(self):
        self.y += 1
        self.z += -1
        hex.locations.append(self)

    
    def nw(self):
        self.x += -1
        self.y += 1
        hex.locations.append(self)

    def sw(self):
        self.x += -1
        self.z += 1
        hex.locations.append(self)

        
    def s(self):
        self.y += -1
        self.z += 1
        hex.locations.append(self)
        
    def se(self):
        self.x += 1
        self.y += -1
        hex.locations.append(self)

    def distance(self, other):
        dx = abs(self.x - other.x)
        dy = abs(self.y - other.y)
        dz = abs(self.z - other.z)
        return (dx + dy + dz) / 2
    
#    def __cmp__(self, other): this probably doesn't make sense to implement
#        if self.x > other.x:
#            return 1
#        if self.x < other.x:
#            return -1
#        if self.y > other.y:
#            return 1
#        if self.y < other.y:
#            return -1
#        if self.z > other.z:
#            return 1
#        if self.z < other.z:
#            return -1

mover = hex(0,0,0)

locations = []

for direction in data:
    if direction == 'ne':
        mover.ne()
    elif direction == 'n':
        mover.n()
    elif direction == 'nw':
        mover.nw()
    elif direction == 'sw':
        mover.sw()
    elif direction == 's':
        mover.s()
    elif direction == 'se':
        mover.se()
    else:
        print('direction does not make sense')
    mover.locations.append([mover.x,mover.y,mover.z])

origin = hex(0,0,0)

print('part a answer is', origin.distance(mover))