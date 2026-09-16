func maxAreaOfIsland(grid [][]int) int {
    
	visited := make(map[[2]int]bool)
    nbRows := len(grid)
    nbCols := len(grid[0])
    maxSizeIsland := 0

    var dfs func(x, y int) int

    dfs = func(x, y int) int {

        if x < 0 || x >= nbRows || y < 0 || y >= nbCols {
            return 0
        }

        if grid[x][y] == 0 {
            return 0
        }

        if _, ok := visited[[2]int{x, y}]; ok {
            return 0
        }

        visited[[2]int{x, y}] = true
		count := 1

        count += dfs(x-1, y)
        count += dfs(x+1, y)
        count += dfs(x, y-1)
        count += dfs(x, y+1)

		return count
    }

    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == 1 {
                if _, ok := visited[[2]int{row, col}]; !ok {
                    tmpCount := dfs(row, col)
					maxSizeIsland = max(tmpCount, maxSizeIsland)
                }
            }
        }
    }

    return maxSizeIsland 
}
