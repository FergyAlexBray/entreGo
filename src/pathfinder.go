package entrego

// Position represents a position on the grid
type Position struct {
	row, col int
}

// pathfinder finds the shorter path in a 2D slice of integers avoiding objects
// and returns the next position compared to the start point closer to the destination point
func Pathfinder(grid [][]int, start, dest Position, avoid []Position) Position {
	// Create a set of positions that have been visited
	visited := make(map[Position]struct{})
	// Create a queue to store the positions to visit
	queue := []Position{}
	// Add the start position to the queue
	queue = append(queue, start)

	// While there are positions in the queue
	for len(queue) > 0 {
		// Get the first position in the queue
		pos := queue[0]
		queue = queue[1:]
		// Add the position to the visited set
		visited[pos] = struct{}{}
		// If the position is the destination, return it
		if pos == dest {
			return pos
		}

		// Get the row and column of the position
		row, col := pos.row, pos.col
		// Loop through the neighbors of the position
		for _, p := range []Position{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}} {
			// If the neighbor is out of bounds or is an object to avoid, skip it
			if p.row < 0 || p.col < 0 || p.row >= len(grid) || p.col >= len(grid[0]) {
				continue
			}
			for _, a := range avoid {
				if a == p {
					continue
				}
			}
			// If the neighbor has not been visited, add it to the queue
			if _, ok := visited[p]; !ok {
				queue = append(queue, p)
			}
		}
	}

	// If the destination is not reachable, return the start position
	return start
}

// func main() {
// grid := [][]int{
// 	{1, 0, 1, 1, 1},
// 	{1, 1, 1, 0, 1},
// 	{1, 0, 1, 1, 1},
// 	{1, 1, 1, 1, 1},
// 	{1, 1, 1, 1, 1},
// }
// start := Position{0, 0}
// dest := Position{4, 4}
// avoid := []Position{
// 	{0, 1},
// 	{1, 3},
// 	{2, 1},
// }

// nextPos := pathfinder(grid, start, dest, avoid)
// fmt.Printf("Next position: (%d, %d)\n", nextPos.row, nextPos.col)
// }

// func Pathfinder(SpaceMap [][]int) [][]int {
// 	goland := [][]int{}
// 	return goland
// }
