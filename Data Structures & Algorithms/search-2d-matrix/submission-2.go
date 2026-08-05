func searchMatrix(matrix [][]int, target int) bool {

	if posVec := findVector(matrix, target) ; posVec != - 1 {
		size := len(matrix[posVec])
		left, right := 0, size - 1
		fmt.Println(matrix[posVec])
		var mid int

		for left <= right {
			mid = (left + right) / 2

			if matrix[posVec][mid] < target { 
				left = mid + 1
			} else if matrix[posVec][mid] > target {
				right = mid - 1
			} else {
				return true 
			}
	}

	}

	return false

}

func findVector(matrix [][]int, target int) int {
	
	size := len(matrix)
	subSize := len(matrix[0]) - 1
	left, right := 0, size - 1
	var mid int
	
	for left < right {
		mid = (left + right) / 2

		if matrix[mid][0] < target && matrix[mid][subSize] < target {
			left = mid + 1 
		} else if matrix[mid][0] > target && matrix[mid][subSize] > target {
			right = mid - 1
		} else {
			return mid 
		}

	}

	if left == right {
		return left
	}

	return - 1 

}
