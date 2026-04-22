# 3. Longest Substring Without Repeating Characters

## Problem Statement
Given a string `s`, find the length of the **longest substring** without repeating characters.

### Examples

**Example 1:**
- **Input:** `s = "abcabcbb"`
- **Output:** `3`
- **Explanation:** The answer is "abc", with the length of 3.

**Example 2:**
- **Input:** `s = "bbbbb"`
- **Output:** `1`
- **Explanation:** The answer is "b", with the length of 1.

**Example 3:**
- **Input:** `s = "pwwkew"`
- **Output:** `3`
- **Explanation:** The answer is "wke", with the length of 3.

### Constraints
- `0 <= s.length <= 5 * 10^4`
- `s` consists of English letters, digits, symbols, and spaces.

---

## Interview Guide: Google Preparation

For a Google interview, how you arrive at the solution is as important as the code itself.

1.  **Clarify the Character Set**: Ask if the string is ASCII (128 chars), Extended ASCII (256 chars), or Unicode. This affects your choice between an array and a map.
2.  **Start with Brute Force**: Briefly explain the $O(n^3)$ approach. It shows you can think through the problem logically before optimizing.
3.  **Optimize with Sliding Window**: This is the most common pattern for substring problems.
4.  **Dry Run**: Be ready to trace through an example like `"abcabcbb"` using pointers on a whiteboard.

---

## Solutions

### Approach 1: Brute Force (Iterative Check)
Check every possible substring and verify if it's unique.

#### Thought Process:
- A substring is defined by its start index `i` and end index `j`.
- We use two nested loops to generate all substrings.
- For each substring, we use a helper function to check for duplicates using a set.

#### Implementation (Go):
```go
func lengthOfLongestSubstringBruteForce(s string) int {
    n := len(s)
    maxLen := 0
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if isAllUnique(s, i, j) {
                if j - i + 1 > maxLen {
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
```

#### Complexity Analysis:
- **Time Complexity**: $O(n^3)$
    - $O(n^2)$ for generating all substrings.
    - $O(n)$ for checking uniqueness.
- **Space Complexity**: $O(k)$, where $k$ is the character set size.

---

### Approach 2: Sliding Window (Using a Set)
We use a "window" that moves across the string. If a duplicate is found, we shrink the window from the left.

#### Thought Process:
1. Initialize two pointers `left` and `right` at index 0.
2. Expand `right` and add characters to a set.
3. If `s[right]` is already in the set, it means we have a duplicate.
4. Shrink the window by moving `left` forward and removing characters from the set until the duplicate is gone.
5. Update the max length at each step.

#### Implementation (Go):
```go
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    set := make(map[byte]bool)
    maxLen, left := 0, 0
    
    for right := 0; right < n; right++ {
        for set[s[right]] {
            delete(set, s[left])
            left++
        }
        set[s[right]] = true
        if right - left + 1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
```

#### Complexity Analysis:
- **Time Complexity**: $O(2n) \approx O(n)$. Each character is visited twice (once by `left`, once by `right`).
- **Space Complexity**: $O(min(n, k))$ for the hash set.

---

### Approach 3: Optimized Sliding Window (Direct Index Jump)
Instead of moving `left` one by one, we store the *next* possible position for `left` in a map.

#### Thought Process:
1. Use a map `charToIndex` to store the last seen position of each character.
2. When we encounter a character `s[right]` that we've seen before, we can jump `left` directly to `charToIndex[s[right]] + 1`.
3. **Crucial**: Ensure we only move `left` forward (never backwards).

#### Implementation (Go):
```go
func lengthOfLongestSubstring(s string) int {
    n := len(s)
    maxLen := 0
    // Maps character to its index
    charMap := make(map[byte]int) 
    
    for left, right := 0, 0; right < n; right++ {
        if idx, found := charMap[s[right]]; found {
            // Jump left to the position after the last seen duplicate
            if idx + 1 > left {
                left = idx + 1
            }
        }
        // Update max length
        if right - left + 1 > maxLen {
            maxLen = right - left + 1
        }
        // Store current index
        charMap[s[right]] = right
    }
    return maxLen
}
```

#### Complexity Analysis:
- **Time Complexity**: $O(n)$. Single pass.
- **Space Complexity**: $O(min(n, k))$ for the map.

---

## Edge Cases to Mention
- **Empty String**: `""` $\rightarrow$ `0`
- **Single Character**: `"a"` $\rightarrow$ `1`
- **All Same Characters**: `"bbbbb"` $\rightarrow$ `1`
- **All Unique Characters**: `"abcdef"` $\rightarrow$ `6`

## Interviewer Follow-ups
- **"What if the string is Unicode?"**
  - Use `map[rune]int` instead of `map[byte]int`.
- **"Can we improve performance further?"**
  - If the character set is restricted (e.g., ASCII), use an array `[128]int` instead of a map to avoid hashing overhead and improve cache locality.

## Summary for Google Interviews
The **Optimized Sliding Window** is the expected production-level solution. Always demonstrate the evolution from Brute Force ($O(n^3)$) $\rightarrow$ Sliding Window ($O(n)$ with 2 visits per char) $\rightarrow$ Optimized Sliding Window ($O(n)$ with 1 visit per char).