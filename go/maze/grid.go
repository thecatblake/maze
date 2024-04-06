package maze

import (
	"strings"
)

type Grid struct {
	rows    int
	columns int
	grid    [][]*Cell
}

func NewGrid(rows int, columns int) *Grid {
	grid := Grid{rows: rows, columns: columns}
	grid.Prepare()
	grid.Configure()
	return &grid
}

func (grid *Grid) Prepare() {
	grid.grid = make([][]*Cell, grid.rows)
	for row := range grid.rows {
		grid.grid[row] = make([]*Cell, grid.columns)
		for col := range grid.columns {
			cell := newCell(row, col)
			grid.grid[row][col] = cell
		}
	}
}

func (grid *Grid) Configure() {
	for row := range grid.rows {
		for col := range grid.columns {
			cell := grid.GetCell(row, col)
			cell.north = grid.GetCell(row-1, col)
			cell.south = grid.GetCell(row+1, col)
			cell.east = grid.GetCell(row, col+1)
			cell.west = grid.GetCell(row, col-1)
		}
	}
}

func (grid *Grid) GetCell(row int, col int) *Cell {
	if row < 0 || row >= grid.rows || col < 0 || col >= grid.columns {
		return nil
	}

	return grid.grid[row][col]
}

func (grid *Grid) String() string {
	output := "+" + strings.Repeat("---+", grid.columns) + "\n"
	for row := range grid.rows {
		top := "|"
		bottom := "+"
		for _, cell := range grid.grid[row] {
			if cell == nil {
				cell = newCell(-1, -1)
			}
			body := "   "
			easternBoundary := "|"

			if cell.linked(cell.east) {
				easternBoundary = " "
			}
			top += body + easternBoundary
			southBoundary := "---"

			if cell.linked(cell.south) {
				southBoundary = "   "
			}
			bottom += southBoundary + "+"
		}
		output += top + "\n" + bottom + "\n"
	}

	return output
}
