func shortestPathBinaryMatrix(grid [][]int) int {
	nRows := len(grid)
	nCols := len(grid[0])
	visited := make(map[[2]int]bool)

	if grid[0][0] == 1 || grid[nRows-1][nCols-1] == 1 {
		return -1
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0,0})
	visited[[2]int{0, 0}] = true

	distance := 1

	for len(queue) > 0 {

		queueLen := len(queue)

		for i:=0; i < queueLen; i++ {
			cell := queue[0]
			queue = queue[1:]
			posX, posY := cell[0], cell[1]

			if posX == nRows-1 && posY == nCols - 1 {
				return distance
			}

			children := [][2]int{
				{posX + 1, posY},
				{posX - 1, posY},
				{posX, posY + 1},
				{posX, posY - 1},
				{posX + 1, posY + 1},
				{posX + 1, posY - 1},
				{posX - 1, posY + 1},
				{posX - 1, posY - 1},
			}

			for _, child := range children {
				childX, childY := child[0], child[1]
				_, ok := visited[child]
				
				if min(childX, childY) < 0 || childX >= nRows ||
				childY >= nCols || ok || grid[childX][childY] == 1 {
					continue
				}

				visited[child] = true
				queue = append(queue, child)
			}

		}

		distance ++

	}

	return -1
}
