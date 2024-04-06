package maze

type Cell struct {
	row    int
	column int
	north  *Cell
	south  *Cell
	east   *Cell
	west   *Cell
	links  map[*Cell]bool
}

func newCell(row int, col int) *Cell {
	cell := Cell{row: row, column: col}
	cell.links = make(map[*Cell]bool)
	return &cell
}

func (cell *Cell) Link(target *Cell, bidi bool) {
	cell.links[target] = true

	if bidi {
		target.Link(cell, false)
	}
}

func (cell *Cell) Unlink(target *Cell, bidi bool) {
	delete(cell.links, target)

	if bidi {
		target.Unlink(cell, false)
	}
}

func (cell *Cell) linked(target *Cell) bool {
	_, ok := cell.links[target]

	if !ok {
		return false
	}

	return true
}

func (cell *Cell) neighbors() []*Cell {
	var neighbors []*Cell

	if cell.north != nil {
		neighbors = append(neighbors, cell.north)
	}

	if cell.south != nil {
		neighbors = append(neighbors, cell.south)
	}

	if cell.east != nil {
		neighbors = append(neighbors, cell.east)
	}

	if cell.west != nil {
		neighbors = append(neighbors, cell.west)
	}

	return neighbors
}
