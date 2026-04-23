package main

import (
	"fmt"
)

// distance calculates the sum of distances for each element in nums.
// arr[i] is the sum of |i - j| for all j such that nums[j] == nums[i] and j != i.
func distance(nums []int) []int64 {
	n := len(nums)
	ans := make([]int64, n)

	// groups maps each number to its list of indices in increasing order.
	groups := make(map[int][]int)
	for i, num := range nums {
		groups[num] = append(groups[num], i)
	}

	// For each group of same values, calculate the prefix and suffix distance sums.
	for _, indices := range groups {
		k := len(indices)
		if k <= 1 {
			continue
		}

		// Calculate total sum of indices for this group to help find rightSum.
		var totalSum int64
		for _, idx := range indices {
			totalSum += int64(idx)
		}

		var leftSum int64
		for i, idx := range indices {
			val := int64(idx)
			countLeft := int64(i)
			countRight := int64(k - 1 - i)

			// rightSum is the sum of indices to the right of the current index.
			rightSum := totalSum - leftSum - val

			// Distance sum = (i * current_idx - sum_of_left_indices) + (sum_of_right_indices - (k-1-i) * current_idx)
			ans[idx] = (countLeft*val - leftSum) + (rightSum - countRight*val)

			// Add current index to leftSum for the next element in the group.
			leftSum += val
		}
	}

	return ans
}

func main() {
	// Test case 1
	nums1 := []int{1, 3, 1, 1, 2}
	fmt.Printf("Input: %v\nOutput: %v\n", nums1, distance(nums1))

	// Test case 2
	nums2 := []int{0, 5, 3}
	fmt.Printf("Input: %v\nOutput: %v\n", nums2, distance(nums2))
}
