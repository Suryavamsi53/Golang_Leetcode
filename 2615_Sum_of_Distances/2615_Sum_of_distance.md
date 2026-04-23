# 2615. Sum of Distances

## Problem Description
You are given a 0-indexed integer array `nums`. There is a new array `arr` of the same length where `arr[i]` is the sum of `|i - j|` for all `j` such that `nums[j] == nums[i]` and `j != i`.

Return the array `arr`.

---

## Approach 1: Brute Force (Intuitive)
For every index `i`, we iterate through the entire array and find all indices `j` where `nums[j] == nums[i]`. We take the absolute difference `|i - j|` and add it to a running sum.

### Steps:
1. Initialize an array `ans` of length `n` with zeros.
2. For `i` from `0` to `n-1`:
   - For `j` from `0` to `n-1`:
     - If `i != j` and `nums[i] == nums[j]`:
       - `ans[i] += abs(i - j)`
3. Return `ans`.

### Complexity Analysis:
- **Time Complexity**: $O(n^2)$ where $n$ is the length of `nums`. In the worst case (all elements same), we do $n$ operations for each of the $n$ elements.
- **Space Complexity**: $O(1)$ (excluding the output array).

> [!WARNING]
> This approach will result in **Time Limit Exceeded (TLE)** for $n = 10^5$.

---

## Approach 2: Grouping with Prefix Sum (Optimized)
Instead of looking at all elements for every `i`, we can group indices of the same value together. For a value $X$, let its indices be sorted: $p_0, p_1, p_2, \dots, p_{k-1}$.

For the $i$-th occurrence (index $p_i$), the sum of distances is:
$$\text{Sum} = \sum_{j=0}^{k-1} |p_i - p_j|$$

Breaking this into two parts (indices before $p_i$ and indices after $p_i$):
1. **Left Part** (indices $j < i$): Since $p_i > p_j$, $|p_i - p_j| = p_i - p_j$.
   $$\sum_{j=0}^{i-1} (p_i - p_j) = (i \times p_i) - \sum_{j=0}^{i-1} p_j$$
2. **Right Part** (indices $j > i$): Since $p_i < p_j$, $|p_i - p_j| = p_j - p_i$.
   $$\sum_{j=i+1}^{k-1} (p_j - p_i) = \left( \sum_{j=i+1}^{k-1} p_{j} \right) - ((k - 1 - i) \times p_i)$$

### Mathematical Derivation of the Optimized Formula:

For a fixed value, let the indices be $P = \{p_0, p_1, \dots, p_{k-1}\}$.
For an index $p_i$, the total distance $D(p_i)$ is:
$$D(p_i) = \sum_{j=0}^{k-1} |p_i - p_j|$$

Split the sum into $j < i$ and $j > i$:
$$D(p_i) = \sum_{j=0}^{i-1} (p_i - p_j) + \sum_{j=i+1}^{k-1} (p_j - p_i)$$

Expanding the sums:
$$D(p_i) = \left( \sum_{j=0}^{i-1} p_i - \sum_{j=0}^{i-1} p_j \right) + \left( \sum_{j=i+1}^{k-1} p_j - \sum_{j=i+1}^{k-1} p_i \right)$$

Since $p_i$ is a constant in these sums:
- $\sum_{j=0}^{i-1} p_i = i \times p_i$
- $\sum_{j=i+1}^{k-1} p_i = (k - 1 - i) \times p_i$

Substitute back:
$$D(p_i) = (i \times p_i - \text{prefixSum}_{i-1}) + (\text{suffixSum}_{i+1} - (k - 1 - i) \times p_i)$$

Where:
- $\text{prefixSum}_{i-1} = \sum_{j=0}^{i-1} p_j$
- $\text{suffixSum}_{i+1} = \sum_{j=i+1}^{k-1} p_j = \text{totalSum} - \text{prefixSum}_{i}$

This allows us to calculate the distance for any index $p_i$ in $O(1)$ time after an initial $O(k)$ pass to find the `totalSum`.

### Complexity Analysis:
- **Time Complexity**: $O(n)$ where $n$ is the length of `nums`. We iterate through the array once to group elements ($O(n)$) and then iterate through each group to calculate distances ($O(n)$ total).
- **Space Complexity**: $O(n)$ to store the indices of each element in a hash map.

---

## Implementation (Go)

```go
func distance(nums []int) []int64 {
    n := len(nums)
    ans := make([]int64, n)
    
    // Group indices by their values
    groups := make(map[int][]int)
    for i, num := range nums {
        groups[num] = append(groups[num], i)
    }
    
    // Calculate sum of distances for each group
    for _, indices := range groups {
        totalSum := int64(0)
        for _, idx := range indices {
            totalSum += int64(idx)
        }
        
        var leftSum int64 = 0
        k := len(indices)
        for i, idx := range indices {
            val := int64(idx)
            countLeft := int64(i)
            countRight := int64(k - 1 - i)
            
            rightSum := totalSum - leftSum - val
            
            ans[idx] = (countLeft * val - leftSum) + (rightSum - countRight * val)
            
            leftSum += val
        }
    }
    
    return ans
}
```

---

## Example Calculation
`nums = [1, 3, 1, 1, 2]`

**Value 1 indices**: `[0, 2, 3]`
- `totalSum = 0 + 2 + 3 = 5`, `k = 3`
- `i = 0, idx = 0`:
  - `leftSum = 0`, `countLeft = 0`
  - `rightSum = 5 - 0 - 0 = 5`, `countRight = 2`
  - `ans[0] = (0*0 - 0) + (5 - 2*0) = 5`
- `i = 1, idx = 2`:
  - `leftSum = 0`, `countLeft = 1`
  - `rightSum = 5 - 0 - 2 = 3`, `countRight = 1`
  - `ans[2] = (1*2 - 0) + (3 - 1*2) = 2 + 1 = 3`
- `i = 2, idx = 3`:
  - `leftSum = 0 + 2 = 2`, `countLeft = 2`
  - `rightSum = 5 - 2 - 3 = 0`, `countRight = 0`
  - `ans[3] = (2*3 - 2) + (0 - 0*3) = 4 + 0 = 4`

**Result**: `[5, 0, 3, 4, 0]` (after processing other values)