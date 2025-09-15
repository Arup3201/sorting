package sorting

func QuickSort(nums []int, low, high int) []int {
	if low >= high || low < 0 {
		return nums
	}

	pivot := partition(nums, low, high)

	nums = QuickSort(nums, low, pivot-1)
	nums = QuickSort(nums, pivot+1, high)

	return nums
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]

	i := low - 1
	for j := low; j <= high-1; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
