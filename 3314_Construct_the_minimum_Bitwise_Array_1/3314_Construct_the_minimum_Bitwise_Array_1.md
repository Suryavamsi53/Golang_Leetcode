# 3314. Construct the Minimum Bitwise Array I

You are given an array nums consisting of n prime integers.

You need to construct an array ans of length n, such that, for each index i, the bitwise OR of ans[i] and ans[i] + 1 is equal to nums[i], i.e. ans[i] OR (ans[i] + 1) == nums[i].

Additionally, you must minimize each value of ans[i] in the resulting array.

If it is not possible to find such a value for ans[i] that satisfies the condition, then set ans[i] = -1.

 

Example 1:

Input: nums = [2,3,5,7]

Output: [-1,1,4,3]

Explanation:

For i = 0, as there is no value for ans[0] that satisfies ans[0] OR (ans[0] + 1) = 2, so ans[0] = -1.
For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 3 is 1, because 1 OR (1 + 1) = 3.
For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 5 is 4, because 4 OR (4 + 1) = 5.
For i = 3, the smallest ans[3] that satisfies ans[3] OR (ans[3] + 1) = 7 is 3, because 3 OR (3 + 1) = 7.
Example 2:

Input: nums = [11,13,31]

Output: [9,12,15]

Explanation:

For i = 0, the smallest ans[0] that satisfies ans[0] OR (ans[0] + 1) = 11 is 9, because 9 OR (9 + 1) = 11.
For i = 1, the smallest ans[1] that satisfies ans[1] OR (ans[1] + 1) = 13 is 12, because 12 OR (12 + 1) = 13.
For i = 2, the smallest ans[2] that satisfies ans[2] OR (ans[2] + 1) = 31 is 15, because 15 OR (15 + 1) = 31.
 

Constraints:

1 <= nums.length <= 100
2 <= nums[i] <= 1000
nums[i] is a prime number.

---
## Solution

### Go
```go
// minBitwiseArray takes an array of integers and returns a transformed array
func minBitwiseArray(nums []int) []int {
    // Create a result slice 'ans' with the same length as 'nums'
    ans := make([]int, len(nums))

    // Loop through each element 'n' in 'nums' along with its index 'i'
    for i, n := range nums {
        // If the number is even, mark it as impossible (-1) and skip further processing
        if n%2 == 0 {
            ans[i] = -1
            continue
        }

        // Initialize 'p' to track the position of the rightmost 0 bit in 'n'
        p := 0
        // Shift 'n' right by 'p' bits and check if the least significant bit is 1
        // Keep moving 'p' forward until we find a 0 bit
        for ((n >> p) & 1) == 1 {
            p++
        }

        // Once we find the rightmost 0 bit at position 'p',
        // subtract 2^(p-1) from 'n' to flip the bit just before it
        ans[i] = n - (1 << (p - 1))
    }

    // Return the transformed array
    return ans
}
```

---
## Explanation

1️⃣ Clear idea of the problem

You are given an array nums.

For each number n in nums, you must find the smallest integer x such that:

x OR (x + 1) == n


OR is the bitwise OR operator.

If no such x exists, return -1 for that position.

You must minimize x for every element independently.

2️⃣ Minimum knowledge needed to understand the problem

You only need three basic concepts:

🔹 A. Binary representation

Every integer can be written in binary (base-2).

Example:

11 = 1011

🔹 B. Bitwise OR (|)

The OR operator sets a bit to 1 if either bit is 1.

1 | 0 = 1
1 | 1 = 1
0 | 0 = 0

🔹 C. What happens when you add 1 to a number

When adding 1:

Trailing 1s turn into 0

The first 0 becomes 1

Example:

0111 + 1 = 1000

3️⃣ Key observation (most important insight)
🔴 Fact 1:

x | (x + 1) is always an odd number

➡️ Therefore:

If n is even, the answer is impossible

We immediately return -1

🔴 Fact 2:

x | (x + 1) produces a number where:

All bits from the least significant bit up to the first 0 bit are 1

This means:

n must have a block of 1s at the right end

The first 0 above that block is the key position

🔍 Visual Scenarios to Understand the Logic

Let's walk through a few examples to see exactly how the bits behave.

**Scenario 1: The "All Ones" Case (n = 7)**
*   **Input**: `n = 7`
*   **Binary**: `...000111`
*   **Step 1**: Find the first `0` starting from the right.
    *   Bit 0 is `1`
    *   Bit 1 is `1`
    *   Bit 2 is `1`
    *   Bit 3 is `0` -> **Stop here!** (`p = 3`)
*   **Step 2**: We need to turn the bit *just before* this zero (bit `p-1` = 2) into a `0`.
    *   Subtract `2^(3-1) = 2^2 = 4`.
    *   `x = 7 - 4 = 3`.
*   **Verification**:
    *   `x = 3` (`...0011`)
    *   `x + 1 = 4` (`...0100`)
    *   `3 | 4` = `0011 | 0100` = `0111` (which is 7). **Success!**

**Scenario 2: The "Mixed Bits" Case (n = 11)**
*   **Input**: `n = 11`
*   **Binary**: `...001011`
*   **Step 1**: Find the first `0`.
    *   Bit 0 is `1`
    *   Bit 1 is `1`
    *   Bit 2 is `0` -> **Stop here!** (`p = 2`)
*   **Step 2**: Turn bit `p-1` (bit 1) into a `0`.
    *   Subtract `2^(2-1) = 2^1 = 2`.
    *   `x = 11 - 2 = 9`.
*   **Verification**:
    *   `x = 9` (`...1001`)
    *   `x + 1 = 10` (`...1010`)
    *   `9 | 10` = `1001 | 1010` = `1011` (which is 11). **Success!**

**Scenario 3: The "Even Number" Case (n = 2)**
*   **Input**: `n = 2`
*   **Logic**: `x | (x+1)` always produces an odd number because the last bit is always `1`.
*   **Result**: Since 2 is even, it's impossible. Return `-1`.

4️⃣ Right algorithm to choose (and why)
Step-by-step algorithm

For each number n:

If n is even
→ store -1

If n is odd

Scan bits from right to left

Find the rightmost 0 bit at position p

Subtract 2^(p-1) from n

This gives the smallest possible x that satisfies:

x | (x + 1) == n

5️⃣ Why this algorithm works (intuition)

The OR result becomes n because:

Adding 1 flips trailing 1s

OR fills them back

Subtracting 2^(p-1):

Minimizes x

Preserves the OR result

6️⃣ Data structures used

Slice ([]int)

To store input and output

No extra data structures

Only integer variables

This is optimal.

7️⃣ Correct Go implementation
func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, n := range nums {

		// Step 1: Even numbers are impossible
		if n%2 == 0 {
			ans[i] = -1
			continue
		}

		// Step 2: Find rightmost 0 bit
		p := 0
		for ((n >> p) & 1) == 1 {
			p++
		}

		// Step 3: Compute minimum x
		ans[i] = n - (1 << (p - 1))
	}

	return ans
}

8️⃣ Time and Space Complexity
⏱ Time Complexity

Each number scans at most log₂(n) bits

Maximum value is 1000 → at most 10 bits

O(n log n)

💾 Space Complexity

Output array only

No auxiliary space

O(1) extra space

9️⃣ Final takeaway (one-line summary)

To minimize x, find the first zero bit in n from the right and clear the bit just below it — even numbers are impossible.