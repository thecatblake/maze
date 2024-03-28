class Cell:
    row = 0
    column = 0
    north = None
    south = None
    east = None
    west = None

    def __init__(self, row, column):
        self.row = row
        self.column = column
        self.links = {}

    def link(self, cell, bidi=True):
        self.links[cell] = True
        if bidi:
            cell.link(self, False)
        return self

    def unlink(self, cell, bidi=True):
        self.links[cell] = False
        if bidi:
            cell.unlink(self, False)
        return self

    def links(self):
        return self.links.keys()

    def linked(self, cell):
        return self.links.get(cell, False)

    def neighbors(self):
        li = []
        if self.north:
            li.append(self.north)
        if self.south:
            li.append(self.south)
        if self.east:
            li.append(self.east)
        if self.west:
            li.append(self.west)
        return li

    