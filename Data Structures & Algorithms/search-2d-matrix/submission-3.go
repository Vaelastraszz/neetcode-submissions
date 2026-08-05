func searchMatrix(matrix [][]int, target int) bool {
	row := findVector(matrix, target)
	if row == -1 {
		return false
	}

	left, right := 0, len(matrix[row])-1

	for left <= right {
		mid := left + (right-left)/2

		switch {
		case matrix[row][mid] < target:
			left = mid + 1
		case matrix[row][mid] > target:
			right = mid - 1
		default:
			return true
		}
	}

	return false
}

func findVector(matrix [][]int, target int) int {
	left, right := 0, len(matrix)-1
	lastCol := len(matrix[0]) - 1

	for left <= right {
		mid := left + (right-left)/2

		if target < matrix[mid][0] {
			right = mid - 1
		} else if target > matrix[mid][lastCol] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}