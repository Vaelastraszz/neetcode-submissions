func rotate(nums []int, k int) {


	if len(nums) < 2 {
		return
	}
	
	n := len(nums)

	for i:=0; i<k; i ++ {

		prev, next := nums[0], nums[1]
		
		for j:=1; j<n; j++ {
			
			nums[j] = prev
			prev = next
			if j == n -1 {
				nums[0] = next
			} else {
				next = nums[j+1]
			}
			
		}
	}

}
