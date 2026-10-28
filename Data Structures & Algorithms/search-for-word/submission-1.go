import "slices"

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	byteW := []byte(word)
	firstLetter := byteW[0]
	searchSlice := []byte{}
	nRows := len(board) - 1
	nCols := len(board[0]) - 1

	var search func(posX, posY int, visited map[[2]int]bool) bool
	search = func(posX, posY int, visited map[[2]int]bool) bool {

		// Vérifier les limites avant d'accéder au plateau.
		if posX < 0 || posX > nRows || posY < 0 || posY > nCols {
			return false
		}

		pos := len(searchSlice)

		// Vérifier que la case n'a pas déjà été visitée.
		if visited[[2]int{posX, posY}] {
			return false
		}

		// Vérifier que la lettre correspond à celle attendue.
		if board[posX][posY] != byteW[pos] {
			return false
		}
		// Ajouter la lettre au chemin courant.
		searchSlice = append(searchSlice, board[posX][posY])

		// Le mot complet a été trouvé.
		if slices.Equal(searchSlice, byteW) {
			return true
		}

		visited[[2]int{posX, posY}] = true
		neighbors := [][2]int{
			{posX - 1, posY},
			{posX + 1, posY},
			{posX, posY - 1},
			{posX, posY + 1},
		}

		for _, neighbor := range neighbors {

			if search(neighbor[0], neighbor[1], visited) {
				return true
			}
		}

		// Backtracking : annuler les choix de cette branche.
		delete(visited, [2]int{posX, posY})
		searchSlice = searchSlice[:len(searchSlice)-1]
		return false

	}

	for row := 0; row <= nRows; row++ {
		for col := 0; col <= nCols; col++ {
			if board[row][col] == firstLetter {
				searchSlice = []byte{}
				if search(row, col, make(map[[2]int]bool)) {
					return true
				}
			}
		}
	}
	return false
}