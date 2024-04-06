package main

import (
	"fmt"
	"maze/maze"
)

func main() {
	grid := maze.NewGrid(4, 4)
	maze.BinaryTree(grid)
	fmt.Println(grid.String())
}
