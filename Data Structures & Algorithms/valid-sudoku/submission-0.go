func isValidSudoku(board [][]byte) bool {

	squareToRow := make([][]byte, 0)

	for _, row := range board {
		if !isValidRow(row) {
			return false
		}
	}

	for i := range board {
		if !isValidColumn(i, board) {
			return false
		}
	}

	squareToRow = squareAsRow(board)

	for _, row := range squareToRow {
		if !isValidRow(row) {
			return false
		}
	}

	return true
}

func isValidRow(row []byte) bool {
	checkMap := make(map[byte]struct{})

	for _, b := range row {

		if b == '.' {
			continue
		}

		if _, ok := checkMap[b]; ok {
			return false
		}

		checkMap[b] = struct{}{}
	}

	return true
}

func isValidColumn(j int, board [][]byte) bool {
	checkMap := make(map[byte]struct{})

	for i := 0; i < len(board); i++ {
		num := board[i][j]

		if num == '.' {
			continue
		}

		if _, ok := checkMap[num]; ok {
			return false
		}

		checkMap[num] = struct{}{}
	}

	return true
}

func squareAsRow(board [][]byte) [][]byte {

	squareToRow := make([][]byte, 0)

	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {

			oneRow := make([]byte, 0, 9)

			oneRow = append(oneRow, board[i][j:j+3]...)
			oneRow = append(oneRow, board[i+1][j:j+3]...)
			oneRow = append(oneRow, board[i+2][j:j+3]...)

			squareToRow = append(squareToRow, oneRow)
		}
	}

	return squareToRow
}