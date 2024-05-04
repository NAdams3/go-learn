package main

func increasingTriplet(nums []int) bool {

	start := len(nums) - 1

	largest := 0
	mid := 0

	found := false

	for i := start; i >= 0; i-- {
		item := nums[i]
		if i == start {
			largest = item
			continue
		}
		if item > largest {
			largest = item
			continue
		}

		if item < largest && item > mid {
			mid = item
			continue
		}

		if item < mid {
			found = true
			break
		}

	}

	return found
}
