func numIslands(grid [][]byte) int {
    visited := make(map[[2]int]bool)
    nbRows := len(grid)
    nbCols := len(grid[0])
    countIslands := 0

    var dfs func(x, y int)

    dfs = func(x, y int) {

        if x < 0 || x >= nbRows || y < 0 || y >= nbCols {
            return
        }

        if grid[x][y] == '0' {
            return
        }

        if _, ok := visited[[2]int{x, y}]; ok {
            return
        }

        visited[[2]int{x, y}] = true

        dfs(x-1, y)
        dfs(x+1, y)
        dfs(x, y-1)
        dfs(x, y+1)
    }

    for row := range grid {
        for col := range grid[row] {
            if grid[row][col] == '1' {
                if _, ok := visited[[2]int{row, col}]; !ok {
                    countIslands++
                    dfs(row, col)
                }
            }
        }
    }

    return countIslands
}