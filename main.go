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

	//print next path
	fmt.Printf("(%d)\n", path[1])

	//print all path
	// for _, p := range path {
	// 	fmt.Printf("(%d, %d)\n", p.X, p.Y)
	// }

	transpalette := []entrego.Transpalette{
		{"transpalette_2", "GO", "[2,2]", "colis2", "BLUE"},
		{"transpalette_2", "TAKE", "[2,2]", "colis2", "GREEN"},
		{"transpalette_2", "WAIT", "", "", ""},
	}
	camion := []entrego.Camion{
		{"camion_b", "WAITING", "200/1000"},
		{"camion_lol", "GONE", "0/1000"},
	}

	entrego.Display(10, transpalette, camion)

	entrego.Parser(&core)

	core.Run()

	entrego.EndMessage(core)
}
