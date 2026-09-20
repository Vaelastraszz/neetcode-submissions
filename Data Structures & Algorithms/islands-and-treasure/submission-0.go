func islandsAndTreasure(grid [][]int) {
	nbRows := len(grid)
	nbCols := len(grid[0])

	const INF = 2147483647

	type Position struct {
		x int
		y int
	}

	queue := []Position{}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				queue = append(queue, Position{
					x: row,
					y: col,
				})
			}
		}
	}

	distance := 0

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
					grid[nx][ny] == INF {

					queue = append(queue, Position{nx, ny})
					grid[nx][ny] = distance + 1
				}
			}
		}

		distance++
	}
}