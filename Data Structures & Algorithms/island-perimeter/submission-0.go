func islandPerimeter(grid [][]int) int {

    nRows := len(grid)
    nCols := len(grid[0])

    perimeter := 0
    visited := make(map[[2]int]bool)

    var getPerimeter func(x, y int)

    getPerimeter = func(x, y int) {

        neighbors := [][2]int{
            {x + 1, y},
            {x - 1, y},
            {x, y + 1},
            {x, y - 1},
        }

        visited[[2]int{x, y}] = true

        for _, neigh := range neighbors {

            nx, ny := neigh[0], neigh[1]

            // Hors grille
            if nx < 0 || nx >= nRows ||
                ny < 0 || ny >= nCols {
                perimeter++
                continue
            }

            // Eau
            if grid[nx][ny] == 0 {
                perimeter++
                continue
            }

            // Terre déjà visitée
            if visited[[2]int{nx, ny}] {
                continue
            }

            // Terre non visitée : DFS
            getPerimeter(nx, ny)
        }
    }

    // On cherche la première case de terre
    
	search : for x := 0; x < nRows; x++ {
        for y := 0; y < nCols; y++ {

            if grid[x][y] == 1 {
                getPerimeter(x, y)
                break search
            }
        }
    }

    return perimeter
}