package entrego

type Point struct {
	X int
	Y int
}

func FindShortestPath(grid [][]int, start, end Point) []Point {
	// create a queue for breadth-first search
	queue := []Point{start}

	// keep track of visited cells to avoid loops
	visited := map[Point]bool{start: true}

	// keep track of the previous cell for each cell
	// this will allow us to reconstruct the path at the end
	prev := map[Point]Point{start: Point{-1, -1}}

	// perform breadth-first search
	for len(queue) > 0 {
		// get the next point in the queue
		p := queue[0]
		queue = queue[1:]

		// check if we have reached the end
		if p == end {
			break
		}

		// check all four directions for valid moves
		for _, dir := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			next := Point{p.X + dir.X, p.Y + dir.Y}

			// skip out-of-bounds cells and cells that are obstacles
			if next.X < 0 || next.X >= len(grid) || next.Y < 0 || next.Y >= len(grid[0]) || grid[next.X][next.Y] == 1 {
				continue
			}

			// skip cells that have already been visited
			if visited[next] {
				continue
			}

			// mark cell as visited and add it to the queue
			visited[next] = true
			queue = append(queue, next)
			prev[next] = p
		}
	}

	// reconstruct the path from the previous map
	path := []Point{end}
	p := end
	for p != start {
		p = prev[p]
		path = append(path, p)
	}

	// reverse the path to get the start to end order
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

// func Pathfinder(SpaceMap [][]int) [][]int {
// 	goland := [][]int{}
// 	return goland
// }
