package advent

func CountIncreasing(nums []int) int {
	count := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}

	return count
}

func CountIncreasingWindowOf3(nums []int) int {
	count := 0
	previousSum := 0

	for i := 2; i < len(nums); i++ {
		sum := nums[i] + nums[i-1] + nums[i-2]

		if sum > previousSum && i != 2 {
			count++
		}

		previousSum = sum
	}

	return count
}

func CountIncreasingWindow(nums []int, window int) int {
	count := 0
	previousSum := 0

	for i := window - 1; i < len(nums); i++ {
		sum := 0

		for j := 0; j < window; j++ {
			sum += nums[i-j]
		}

		if sum > previousSum && i != window-1 {
			count++
		}

		previousSum = sum
	}

	return count
}
