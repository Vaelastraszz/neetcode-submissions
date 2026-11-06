func searchInsert(nums []int, target int) int {
	// Initialiser left et right 
	left, right := 0, len(nums) - 1 
	insertPos := len(nums)
	
	// Tant que la queue est inferieure ou egal a la tete
	for left <= right {
	// On calcule mid
		mid := (left + right) / 2
 	// Si on trouve target a nums[mid] on retourne mid
		if nums[mid] == target {
			return mid
		}	
	// Si le nums[mid] est inferieur a target on bouge left a mid + 1
		if nums[mid] < target {
			left = mid + 1
		}
	// Si mid est supérieur a target on bouge right a mid - 1 et on sauve
	// la position d insertion
		if nums[mid] > target {
			right = mid - 1
			insertPos = mid
		}
	}

	// on retourne la position d insertion
	return insertPos
}
