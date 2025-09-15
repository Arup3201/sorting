package sorting

func InsertionSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		j := i
		x := nums[j]
		for j > 0 && nums[j-1] > x {
			nums[j] = nums[j-1]
			j--
		}
		nums[j] = x
	}

	return nums
}
