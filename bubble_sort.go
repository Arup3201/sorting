package sorting

func BubbleSort(nums []int) []int {
	n := len(nums)
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < n; i++ {
			if nums[i] < nums[i-1] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				sorted = false
			}
		}
		n--
	}

	return nums
}
