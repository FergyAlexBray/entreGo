package main

import (
	"fmt"

	entrego "github.com/FergyAlexBray/entreGo/src"
)

func main() {
	core := entrego.Core{}

	grid := [][]int{
		{0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
	}

	// define the start and end points
	start := entrego.Point{0, 6}
	end := entrego.Point{6, 6}

	// find the shortest path
	path := entrego.FindShortestPath(grid, start, end)

	fmt.Println("Shortest path:")
	//fmt.Printf("(%d)\n", path[1])
	for _, p := range path {
		fmt.Printf("(%d, %d)\n", p.X, p.Y)
	}

	entrego.Parser(&core)

	core.Run()

	entrego.EndMessage(core)
}
