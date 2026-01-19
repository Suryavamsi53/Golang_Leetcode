# 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold

Given a m x n matrix mat and an integer threshold, return the maximum side-length of a square with a sum less than or equal to threshold or return 0 if there is no such square.

 

Example 1:


Input: mat = [[1,1,3,2,4,3,2],[1,1,3,2,4,3,2],[1,1,3,2,4,3,2]], threshold = 4
Output: 2
Explanation: The maximum side length of square with sum less than 4 is 2 as shown.

Example 2:
Input: mat = [[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2],[2,2,2,2,2]], threshold = 1
Output: 0
 

Constraints:

m == mat.length
n == mat[i].length
1 <= m, n <= 300
0 <= mat[i][j] <= 10^4
0 <= threshold <= 10^5

---
## Solution

### Go
```go
package main


func maxSideLength(mat [][]int, threshold int) int {

	// m = number of rows in matrix
	// n = number of columns in matrix
	m, n := len(mat), len(mat[0])

	// ------------------------------------------------
	// STEP 1: Build Prefix Sum Matrix
	// ------------------------------------------------
	// prefix[i][j] will store the sum of elements
	// from mat[0][0] to mat[i-1][j-1]
	// Extra row and column (m+1, n+1) avoid boundary checks

	prefix := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		prefix[i] = make([]int, n+1)
	}

	// Build the prefix sum using dynamic programming
	// i -> row index in prefix matrix
	// j -> column index in prefix matrix
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			// Current cell value
			// + sum above
			// + sum on the left
			// - overlapping top-left area
			prefix[i][j] =
				mat[i-1][j-1] +
					prefix[i-1][j] +
					prefix[i][j-1] -
					prefix[i-1][j-1]
		}
	}

	// ------------------------------------------------
	// Helper function: Check if a k x k square is possible
	// ------------------------------------------------
	// Returns true if ANY square of size k has sum <= threshold
	canForm := func(k int) bool {

		// i = starting row of square
		for i := 0; i+k <= m; i++ {

			// j = starting column of square
			for j := 0; j+k <= n; j++ {

				// Calculate sum of k x k square
				// using prefix sum in O(1)
				sum :=
					prefix[i+k][j+k] -
						prefix[i][j+k] -
						prefix[i+k][j] +
						prefix[i][j]

				// If this square satisfies the condition
				if sum <= threshold {
					return true
				}
			}
		}

		// No valid square of size k found
		return false
	}

	// ------------------------------------------------
	// STEP 2: Binary Search on Square Size
	// ------------------------------------------------

	// Minimum square size = 0
	// Maximum square size = min(m, n)
	low, high := 0, min(m, n)

	// Stores the largest valid square size found
	ans := 0

	for low <= high {

		// Try middle size
		mid := (low + high) / 2

		// Check if a square of size mid is possible
		if canForm(mid) {

			// mid is valid, save it
			ans = mid

			// Try to find a bigger square
			low = mid + 1

		} else {

			// mid is too large, try smaller
			high = mid - 1
		}
	}

	// Return the largest valid square size
	return ans
}

// ------------------------------------------------
// Helper function to find minimum of two integers
// ------------------------------------------------
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
---
## Explanation

🔍 Problem in Simple Words

You are given:

A matrix (m × n) of non-negative numbers

An integer threshold

Your task:

👉 Find the largest square sub-matrix (same number of rows and columns) such that
the sum of all numbers inside that square ≤ threshold.

If no square satisfies this, return 0.

🧠 Key Challenges

Matrix size can be up to 300 × 300

Brute force checking every square is too slow

So we need to:

⚡ Quickly calculate the sum of any square

🔍 Efficiently try different square sizes

🧩 Step 1: Prefix Sum (The Superpower)
❓ Why Prefix Sum?

Without prefix sum:

A k × k square → k² additions

Too slow for large matrices ❌

With prefix sum:

Any square sum → O(1) time ✅

📌 Prefix Sum Definition

Create a matrix prefix of size (m+1) × (n+1) where:

prefix[i][j] =
sum of all elements in mat
from (0,0) to (i-1, j-1)

📐 Prefix Sum Formula
prefix[i][j] =
    mat[i-1][j-1]
  + prefix[i-1][j]
  + prefix[i][j-1]
  - prefix[i-1][j-1]

🔁 Loop Holds (Important!)
for i := 1; i <= m; i++ {      // row
    for j := 1; j <= n; j++ {  // column
        prefix[i][j] = ...
    }
}

Variable	Meaning
i	current row
j	current column
prefix[i][j]	sum of rectangle from top-left
🧪 Example Prefix Matrix

Matrix:

1 1 3
1 1 3
1 1 3


Prefix:

0  0  0  0
0  1  2  5
0  2  4 10
0  3  6 15


👉 prefix[2][2] = 4
(sum of 2×2 top-left square)

🧩 Step 2: Getting Any Square Sum in O(1)

To calculate the sum of a square:

Side length = k

Top-left corner = (r, c)

📐 Square Sum Formula
sum =
prefix[r+k][c+k]
- prefix[r][c+k]
- prefix[r+k][c]
+ prefix[r][c]

🔁 Square Scanning Loops
for r := 0; r+k <= m; r++ {
    for c := 0; c+k <= n; c++ {
        // check square starting at (r, c)
    }
}

Variable	Meaning
r	top row of square
c	left column
k	square size
🧪 Scenario: k = 2

Square at (0,0):

1 1
1 1


Calculation:

prefix[2][2] - prefix[0][2] - prefix[2][0] + prefix[0][0]
= 4 - 0 - 0 + 0
= 4 ≤ threshold ✅

🧩 Step 3: Binary Search on Square Size
❓ Why Binary Search?

If:

A square of size k works

Size k+1 fails

Then:
👉 Any size larger will also fail

This makes the answer monotonic → perfect for binary search.

🔁 Binary Search Loop
low := 0
high := min(m, n)

for low <= high {
    mid := (low + high) / 2

    if canForm(mid) {
        answer = mid
        low = mid + 1   // try bigger
    } else {
        high = mid - 1  // try smaller
    }
}

🧪 Binary Search Example
Try	Size	Result
1	1×1	✅
2	2×2	✅
3	3×3	❌

✔️ Final Answer = 2

🧠 Final Loop Flow (Crystal Clear)
Binary Search Loop
 └── canForm(size)
     └── Row Loop
         └── Column Loop
             └── O(1) square sum check

🎯 Final Intuition (Easy to Remember)

“Prefix sum tells me the weight of any square instantly.
Binary search helps me find the biggest square that fits under the limit.”