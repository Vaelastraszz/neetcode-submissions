func hasDuplicate(nums []int) bool {
    
	existMap := make(map[int]struct{})
	var exist bool
	for _, num := range nums {
		if _,ok := existMap[num]; ok {
			exist = true 
			return exist
		} else {
			existMap[num] = struct{}{}
		}
	}
	return exist
}
