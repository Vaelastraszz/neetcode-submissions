func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	refColor := image[sr][sc]

	if refColor == color {
		return image
	}

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) {
			return
		}

		if image[r][c] != refColor {
			return
		}

		image[r][c] = color

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	dfs(sr, sc)

	return image
}