package main

import (
	"fmt"
)

// Approach 1: Brute Force
// Time Complexity: O(n^3) - Generation O(n^2) * Validation O(n)
// Space Complexity: O(k)
func lengthOfLongestSubstringBruteForce(s string) int {
	n := len(s)
	maxLen := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if isAllUnique(s, i, j) {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}
	return maxLen
}

func isAllUnique(s string, start, end int) bool {
	seen := make(map[byte]bool)
	for k := start; k <= end; k++ {
		if seen[s[k]] {
			return false
		}
		seen[s[k]] = true
	}
	return true
}

// Approach 2: Sliding Window
// Time Complexity: O(2n) = O(n)
// Space Complexity: O(min(n, k))
func lengthOfLongestSubstringSlidingWindow(s string) int {
	n := len(s)
	set := make(map[byte]bool)
	maxLen, left := 0, 0

	for right := 0; right < n; right++ {
		// If duplicate found, shrink window from left
		for set[s[right]] {
			delete(set, s[left])
			left++
		}
		set[s[right]] = true
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// Approach 3: Optimized Sliding Window (Direct Index Jump)
// Time Complexity: O(n)
// Space Complexity: O(min(n, k))
func lengthOfLongestSubstringOptimized(s string) int {
	n := len(s)
	maxLen := 0
	charMap := make(map[byte]int)

	for left, right := 0, 0; right < n; right++ {
		if index, found := charMap[s[right]]; found {
			// Ensure left only moves forward
			if index+1 > left {
				left = index + 1
			}
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		charMap[s[right]] = right
	}
	return maxLen
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{" ", 1},
		{"au", 2},
	}

	fmt.Println("--- Testing Optimized Sliding Window ---")
	for _, tc := range testCases {
		result := lengthOfLongestSubstringOptimized(tc.input)
		fmt.Printf("Input: %q | Expected: %d | Result: %d\n", tc.input, tc.expected, result)
	}

	fmt.Println("\n--- Testing Brute Force ---")
	for _, tc := range testCases {
		result := lengthOfLongestSubstringBruteForce(tc.input)
		fmt.Printf("Input: %q | Expected: %d | Result: %d\n", tc.input, tc.expected, result)
	}
}
