package main

import (
	"math/bits"
	"sort"
)

// given array arr consisting of non-negative intergers and an integer k, find the minimum length of any sub-array of arr such that if all elements in this sub-array are represented in binary notation and concatenated to form a binary string then the number of 1s in the resulting string is at least k. If no such array exists print -1.

func minimumSubArray(arr []uint, length int) int {

	greaterThanOne := make([]int, 0)

	isOne := false

	for _, item := range arr {
		ones := bits.OnesCount(item)
		if ones >= length {
			isOne = true
			break
		}

		if ones > 1 {
			greaterThanOne = append(greaterThanOne, ones)
		}
	}

	if isOne {
		return 1
	}

	sort.Slice(greaterThanOne, func(i, j int) bool {
		return greaterThanOne[i] > greaterThanOne[j]
	})

	count := 0

	for _, item := range greaterThanOne {
		if length <= 0 {
			return count
		}
		length = length - item
		count++
	}

	if length > 0 {
		return count + length
	}

	return -1
}
