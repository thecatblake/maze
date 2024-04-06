package maze

import (
	"math/rand"
)

func BinaryTree(grid *Grid) {
	for row := range grid.rows {
		for _, cell := range grid.grid[row] {
			var neighbors []*Cell
			if cell.north != nil {
				neighbors = append(neighbors, cell.north)
			}
			if cell.east != nil {
				neighbors = append(neighbors, cell.east)
			}
			if len(neighbors) != 0 {
				neighbor := neighbors[rand.Intn(len(neighbors))]

				if neighbor != nil {
					cell.Link(neighbor, true)
				}
			}
		}
	}
}
