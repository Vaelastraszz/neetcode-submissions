func orangesRotting(grid [][]int) int {
	nbRows := len(grid)
	nbCols := len(grid[0])
	mapFreshFruits := make(map[[2]int]bool)

	type Position struct {
		x int
		y int
	}

	queue := []Position{}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 2 {
				queue = append(queue, Position{
					x: row,
					y: col,
				})
			}

			if grid[row][col] == 1 {
				mapFreshFruits[[2]int{row, col}] = true
			}
		}
	}

	if len(mapFreshFruits) == 0 {
		return 0
	}

	minutesElapsed := -1

	for len(queue) > 0 {
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			current := queue[0]
			queue = queue[1:]

			x := current.x
			y := current.y

			voisins := [4]Position{
				{x + 1, y},
				{x - 1, y},
				{x, y + 1},
				{x, y - 1},
			}

			for _, voisin := range voisins {
				nx := voisin.x
				ny := voisin.y

				if nx >= 0 && nx < nbRows &&
					ny >= 0 && ny < nbCols &&
					grid[nx][ny] == 1 {

					queue = append(queue, Position{nx, ny})
					delete(mapFreshFruits, [2]int{nx, ny})
					grid[nx][ny] = 2
				}
			}
		}

		minutesElapsed++
	}

	if len(mapFreshFruits) > 0 {
		return -1
	}

	return minutesElapsed
}