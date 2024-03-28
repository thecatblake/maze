from cell import Cell


class Grid:
    rows = 0
    columns = 0

    def __init__(self, rows, columns):
        self.rows = rows
        self.columns = columns

        self.grid = self.prepare_grid()
        self.configure_cells()

    def prepare_grid(self):
        return [[Cell(j, i) for i in range(self.columns)] for j in range(self.rows)]

    def configure_cells(self):
        for row in range(self.rows):
            for col in range(self.columns):
                cell = self.grid[row][col]
                cell.north = self[row - 1, col]
                cell.south = self[row + 1, col]
                cell.east = self[row, col + 1]
                cell.west = self[row, col - 1]

    def __getitem__(self, item):
        if isinstance(item, tuple) and len(item) == 2:
            row, col = item

            if 0 <= row < self.rows and 0 <= col < self.columns:
                return self.grid[row][col]
            else:
                return None

    def each_cell(self):
        for row in range(self.rows):
            for col in range(self.columns):
                yield self[row, col]

    def __str__(self):
        output = "+" + "---+" * self.columns + "\n"
        for row in self.grid:
            top = "|"
            bottom = "+"
            for cell in row:
                if cell is None:
                    cell = Cell(-1, 1)
                body = "   "
                east_boundary = " " if cell.linked(cell.east) else "|"
                top += body + east_boundary
                south_boundary = "   " if cell.linked(cell.south) else "---"
                bottom += south_boundary + "+"
            output += top + "\n" + bottom + "\n"

        return output
