# 1266. Minimum Time Visiting All Points

**Difficulty:** Easy  
**Type:** Greedy, Math  
**Solved:** Yes

---

## 📋 Problem Explanation

On a 2D plane, there are n points with integer coordinates `points[i] = [xi, yi]`. You need to visit all the points in the **exact order** they appear in the array and return the **minimum time in seconds** required.

### Movement Rules:
In **1 second**, you can perform **ONE** of these actions:
1. Move **vertically** by one unit (up or down)
2. Move **horizontally** by one unit (left or right)
3. Move **diagonally** (one unit vertically AND one unit horizontally simultaneously) - this counts as moving √2 units but takes only 1 second!

### Important Constraints:
- You must visit points in the **given order**
- You can pass through points without counting them as visits
- Find the **minimum total time** to visit all points

---

## 🎯 Basic Requirements to Understand This Problem

To solve this problem, you need to understand:

1. **Coordinate System**: Understanding how to calculate distance between two points on a 2D plane
2. **Manhattan Distance vs Chebyshev Distance**: 
   - Manhattan: |x2-x1| + |y2-y1| (moving only horizontally/vertically)
   - Chebyshev: max(|x2-x1|, |y2-y1|) (optimal with diagonal movement)
3. **Greedy Approach**: The optimal strategy is to move diagonally as much as possible, then move in one direction
4. **Time Calculation**: Understanding that diagonal movement is more efficient than sequential horizontal + vertical movement
5. **Absolute Values**: Using absolute values to calculate distances regardless of direction

---

## 💡 Key Insight

**Why is the answer `max(|x2-x1|, |y2-y1|)`?**

Consider moving from point (x1, y1) to (x2, y2):
- Horizontal distance: |x2 - x1|
- Vertical distance: |y2 - y1|

If we can move diagonally, we move in both directions simultaneously. The limiting factor is the **larger distance**. Once we've covered the smaller distance diagonally, we move in one direction for the remaining distance.

### Visual Representation: Diagonal Movement Advantage

#### Scenario 1: Moving from (0,0) to (3,5)

**Without Diagonal Movement (Sequential):**
```
     y
     |
   5 |           E (3,5) ✓ DESTINATION
     |           |
   4 |           |
     |           |
   3 |     D (3,3)
     |     |
   2 |     |
     |     |
   1 |     |
     |     |
   0 |S---+-----+---
       0   1   2   3   x

Path: (0,0) → (1,0) → (2,0) → (3,0) → (3,1) → (3,2) → (3,3) → (3,4) → (3,5)
Movements: 3 right + 5 up = 8 STEPS (8 seconds)
```

**With Diagonal Movement (Optimal):**
```
     y
     |
   5 |           E (3,5) ✓ DESTINATION
     |          /|
   4 |        /  |
     |      /    |
   3 |    D     |  (3 diagonal steps)
     |  /       |  (2 vertical steps)
   2 | /        |
     |/         |
   1 |/          |
     |/          |
   0 S-----------+---
       0   1   2   3   x

Path (Diagonal First): 
Step 1: (0,0) → (1,1)  [DIAGONAL]
Step 2: (1,1) → (2,2)  [DIAGONAL]
Step 3: (2,2) → (3,3)  [DIAGONAL]
Step 4: (3,3) → (3,4)  [VERTICAL]
Step 5: (3,4) → (3,5)  [VERTICAL]

Movements: 3 diagonal + 2 vertical = 5 STEPS (5 seconds)
Total Time: max(3, 5) = 5 seconds ✓
```

**Key Insight:**
- Horizontal distance (dx) = |3 - 0| = 3
- Vertical distance (dy) = |5 - 0| = 5
- Time needed = **max(3, 5) = 5 seconds**
- We move diagonally for `min(dx, dy) = 3` steps, then move in one direction for `|dx - dy| = 2` steps

---

### Visual Representation: Why Diagonal is Better

```
MOVEMENT COMPARISON:

Horizontal Only (1,0):   1 unit right per second
Vertical Only (0,1):     1 unit up per second
Diagonal (1,1):          1 unit right + 1 unit up per second ✓ MOST EFFICIENT

Movement Grid:
┌──────┬──────┬──────┬──────┐
│      │      │      │      │
│  →   │  →   │  →   │ (3,5)│ ← Destination
│      │      │      │ ↑    │
├──────┼──────┼──────┼──┤
│      │      │      │ ↑    │
│      │      │      │ ↑    │
│      │      │  ↗   │ (3,3) Diagonal done
│      │      │ ↗    │      │
├──────┼──────┼──────┼──────┤
│      │      │ ↗    │      │
│      │  ↗   │      │      │
│  ↗   │      │      │      │
│ (0,0)│      │      │      │
└──────┴──────┴──────┴──────┘

Diagonal Path: S → D → D → D → E → E
Time: 1 + 1 + 1 + 1 + 1 = 5 seconds
```

---

## 📝 Code Explanation (Go)

### Algorithm Flow Diagram

For `points = [[1,1],[3,4],[-1,0]]`:

```
         ┌─────────────────────┐
         │       START         │
         └────────────┬────────┘
                      │
                      v
         ┌─────────────────────┐
         │  totalTime = 0      │
         └────────────┬────────┘
                      │
                      v
         ┌─────────────────────────────┐
         │ Loop: i = 1 to len(points)  │
         └────────────┬────────────────┘
                      │
          ┌───────────┴──────────┐
          │                      │
          v                      v
    ┌───────────────┐      ┌────────────────┐
    │ i < 3? YES    │      │ i < 3? NO      │
    └────────┬──────┘      └────────┬───────┘
             │                      │
    ┌────────v──────────┐           │
    │ i = 1             │           │
    │ Get P1=(1,1)      │           │
    │ Get P2=(3,4)      │           │
    │ dx = 2, dy = 3    │           │
    │ time = max(2,3)=3 │           │
    │ totalTime = 3     │           │
    │ Loop Again        │           │
    └────────┬──────────┘           │
             │                      │
    ┌────────v──────────┐           │
    │ i = 2             │           │
    │ Get P2=(3,4)      │           │
    │ Get P3=(-1,0)     │           │
    │ dx = 4, dy = 4    │           │
    │ time = max(4,4)=4 │           │
    │ totalTime = 7     │           │
    │ Loop Again        │           │
    └────────┬──────────┘           │
             │                      │
    ┌────────v──────────┐           │
    │ i = 3             │           │
    │ i < 3? NO         │───────────┘
    │ EXIT LOOP         │
    └────────┬──────────┘
             │
             v
    ┌────────────────────┐
    │ RETURN totalTime=7 │
    └────────────────────┘
```

---

### Line-by-Line Code Walkthrough

```go
func minTimeToVisitAllPoints(points [][]int) int {
```
- **Function Declaration**: Takes a 2D slice of integers (each inner slice has 2 elements: [x, y])
- Returns an integer (total time in seconds)

```go
	totalTime := 0
```
- **Initialize Counter**: `totalTime` accumulates the time needed to travel between all consecutive points
- Starts at 0 because initially no travel has occurred

```go
	for i := 1; i < len(points); i++ {
```
- **Loop Through Points**: Start from index 1 (not 0)
- **Why index 1?**: We need to calculate distance between consecutive points (i-1 and i)
- **Loop Condition**: Continue until we've processed all pairs of consecutive points
- **Example**: If points = [[1,1],[3,4],[-1,0]], loop runs for i=1 and i=2

```go
		x1, y1 := points[i-1][0], points[i-1][1]
```
- **Extract Previous Point**: Get coordinates of the previous point (where we're coming from)
- `points[i-1][0]` = x-coordinate of previous point
- `points[i-1][1]` = y-coordinate of previous point
- **Scenario**: If i=1, this gets point[0] which is [1,1], so x1=1, y1=1

```go
		x2, y2 := points[i][0], points[i][1]
```
- **Extract Current Point**: Get coordinates of the current point (where we're going to)
- `points[i][0]` = x-coordinate of current point
- `points[i][1]` = y-coordinate of current point
- **Scenario**: If i=1, this gets point[1] which is [3,4], so x2=3, y2=4

```go
		dx := abs(x2 - x1)
```
- **Calculate Horizontal Distance**: Absolute difference between x-coordinates
- `abs()` ensures distance is positive (direction doesn't matter for distance)
- **Scenario**: |3 - 1| = 2 (move 2 units horizontally)

```go
		dy := abs(y2 - y1)
```
- **Calculate Vertical Distance**: Absolute difference between y-coordinates
- `abs()` ensures distance is positive
- **Scenario**: |4 - 1| = 3 (move 3 units vertically)

```go
		totalTime += max(dx, dy)
```
- **Add Time for This Segment**: Use `max(dx, dy)` because:
  - We can move diagonally, covering both horizontal and vertical distance simultaneously
  - The time taken equals the **larger** of the two distances
  - If dx=2 and dy=3, we move diagonally 2 steps (covering 2 horizontal and 2 vertical), then 1 more step vertically = 3 seconds total
- **Scenario**: max(2, 3) = 3 seconds to go from [1,1] to [3,4]

```go
	return totalTime
```
- **Return Result**: Return the accumulated total time to visit all points

### Helper Functions

```go
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```
- **Absolute Value Function**: Returns positive distance regardless of sign
- **Purpose**: Convert negative distances to positive
- **Scenario**: abs(-5) = 5, abs(3) = 3

```go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
- **Maximum Function**: Returns the larger of two numbers
- **Purpose**: Determines the time needed for this segment (Chebyshev distance)
- **Scenario**: max(3, 5) = 5

### Distance Calculation Visualization

#### Example: From (1,1) to (3,4)

```
COORDINATE GRID:
        
     y
     4 |          ★ (3,4) ← DESTINATION
       |          |  ↑
     3 |          |  | dy = 3
       |          |  ↑
     2 |          |
       |          |
     1 |★---------+-  
       |(1,1)     |
       |          dx = 2
     0 |__________→__________
       0    1    2    3    x

DISTANCE FORMULA:
dx = |x2 - x1| = |3 - 1| = 2
dy = |y2 - y1| = |4 - 1| = 3

TIME = max(dx, dy) = max(2, 3) = 3 seconds

WHY max() and not sum()?
- We can move DIAGONALLY (both directions at once)
- We move diagonally min(2,3) = 2 times (covers both x and y for 2 steps)
- Then we move vertically 1 more time (to cover remaining y distance)
- Total: 2 + 1 = 3 = max(2,3) ✓
```

#### Movement Types Comparison Table

```
┌─────────────────────┬──────────┬──────────┬─────────────────────┐
│ Movement Type       │ Cost     │ Example  │ When Optimal        │
├─────────────────────┼──────────┼──────────┼─────────────────────┤
│ Horizontal Only (→) │ 1 sec    │ (0,0)→   │ When dy=0           │
│                     │ per unit │ (5,0)    │ Straight line only  │
├─────────────────────┼──────────┼──────────┼─────────────────────┤
│ Vertical Only (↑)   │ 1 sec    │ (0,0)→   │ When dx=0           │
│                     │ per unit │ (0,5)    │ Straight line only  │
├─────────────────────┼──────────┼──────────┼─────────────────────┤
│ Diagonal (↗)        │ 1 sec    │ (0,0)→   │ MOST EFFICIENT ✓    │
│                     │ covers 2 │ (1,1)    │ Moves in 2 dirs     │
│                     │ units    │ per sec  │ simultaneously      │
├─────────────────────┼──────────┼──────────┼─────────────────────┤
│ Mixed (Optimal)     │ 1 sec    │ (0,0)→   │ ALWAYS BEST ✓       │
│ Diagonal + Linear   │ per step │ (3,5)    │ Use diagonal till   │
│                     │ = max()  │ = 5 secs │ one dist done       │
└─────────────────────┴──────────┴──────────┴─────────────────────┘
```

---

## 📊 Complete Example Walkthrough

**Input:** `points = [[1,1],[3,4],[-1,0]]`

### Visual Map of All Points:

```
        y
        |
    5   |
        |
    4   | P2(3,4) ●
        |    ↗  ↘
    3   |  ↗      ↘
        |↗          ↘
    2   |            ↘
        |              ↘ P1(1,1) ●
    1   |   ↗   ↗ ↗ ↗ ↗
        |↗        ↗
    0   |------●---------|---x
        |  P3(-1,0)
   -1   |
        |
        -1  0  1  2  3  4
```

### Step-by-Step Movement

#### **Leg 1: P1 (1,1) → P2 (3,4)**

```
Coordinates:
From: (1, 1)
To:   (3, 4)

Distance Calculation:
dx = |3 - 1| = 2 (move 2 units RIGHT)
dy = |4 - 1| = 3 (move 3 units UP)

Movement Strategy (using diagonal):
┌──────┬──────┬──────┐
│      │      │      │
│      │      │  P2  │ (3,4) ← GOAL
│      │      │  ●   │
├──────┼──────┼──────┤
│      │      │  ●   │ (3,3) ← Step 3
│      │  ●   │      │ Diagonal done!
├──────┼──────┼──────┤
│      │  ●   │      │ (2,2) ← Step 2
│  ●   │      │      │
├──────┼──────┼──────┤
│  P1  │      │      │
│  ●   │      │      │ (1,1) ← START
└──────┴──────┴──────┘

Step 1: (1,1) → (2,2)  ↗ DIAGONAL    [1 sec]
Step 2: (2,2) → (3,3)  ↗ DIAGONAL    [1 sec]
Step 3: (3,3) → (3,4)  ↑ VERTICAL    [1 sec]

Time: max(2, 3) = 3 seconds ✓
```

**Iteration 1 Details:**
- Previous point: [1,1] → x1=1, y1=1
- Current point: [3,4] → x2=3, y2=4
- dx = |3-1| = **2**
- dy = |4-1| = **3**
- Time for this segment = max(2,3) = **3 seconds**
- totalTime = 0 + 3 = **3**

---

#### **Leg 2: P2 (3,4) → P3 (-1,0)**

```
Coordinates:
From: (3, 4)
To:   (-1, 0)

Distance Calculation:
dx = |-1 - 3| = 4 (move 4 units LEFT)
dy = |0 - 4| = 4 (move 4 units DOWN)

Movement Strategy (equal distances = perfect diagonal):
┌──────┬──────┬──────┬──────┬──────┐
│  P2  │      │      │      │      │
│  ●   │      │      │      │      │ (3,4) ← START
├──────┼──────┼──────┼──────┼──────┤
│      │  ●   │      │      │      │ (2,3) ← Step 1
│      │      │      │      │      │
├──────┼──────┼──────┼──────┼──────┤
│      │      │  ●   │      │      │ (1,2) ← Step 2
│      │      │      │      │      │
├──────┼──────┼──────┼──────┼──────┤
│      │      │      │  ●   │      │ (0,1) ← Step 3
│      │      │      │      │      │
├──────┼──────┼──────┼──────┼──────┤
│      │      │      │      │  P3  │ (-1,0) ← Step 4, GOAL
│      │      │      │      │  ●   │
└──────┴──────┴──────┴──────┴──────┘

Step 1: (3,4) → (2,3)  ↙ DIAGONAL    [1 sec]
Step 2: (2,3) → (1,2)  ↙ DIAGONAL    [1 sec]
Step 3: (1,2) → (0,1)  ↙ DIAGONAL    [1 sec]
Step 4: (0,1) → (-1,0) ↙ DIAGONAL    [1 sec]

Time: max(4, 4) = 4 seconds ✓
Perfect diagonal movement! (equal dx and dy)
```

**Iteration 2 Details:**
- Previous point: [3,4] → x1=3, y1=4
- Current point: [-1,0] → x2=-1, y2=0
- dx = |-1-3| = **4**
- dy = |0-4| = **4**
- Time for this segment = max(4,4) = **4 seconds**
- totalTime = 3 + 4 = **7**

---

### Complete Path Visualization

```
Full Journey: [1,1] → [3,4] → [-1,0]

Coordinate Space (Full View):

        y
        |
    5   |
        |
    4   | [1,1]P1 ↗ ↗ [3,4]P2 ↙ ↙
        |        ↗       ↙
    3   |       ↗       ↙
        |      ↗       ↙
    2   |    ↗       ↙
        |   ↗       ↙
    1   | ↗       ↙
        |/       ↙
    0   +-------[−1,0]P3---x
        |
   -1   |
        |
        -1  0  1  2  3
        
Leg 1 (→): 3 seconds
Leg 2 (↙): 4 seconds
━━━━━━━━━━━━━━━━━
Total: 7 seconds ✓
```

**Output:** 7 ✓

---

## 🔵 Go Implementation

```go
package main

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	// Initialize total time counter to track cumulative seconds
	totalTime := 0

	// Loop through each point starting from index 1 (the second point)
	// We start from 1 because we need a previous point to calculate distance from
	for i := 1; i < len(points); i++ {
		// Extract coordinates of the previous point (where we're coming from)
		x1 := points[i-1][0]  // x-coordinate of previous point
		y1 := points[i-1][1]  // y-coordinate of previous point

		// Extract coordinates of the current point (where we're going to)
		x2 := points[i][0]  // x-coordinate of current point
		y2 := points[i][1]  // y-coordinate of current point

		// Calculate the absolute horizontal distance between the two points
		// abs() ensures the distance is always positive regardless of direction
		dx := int(math.Abs(float64(x2 - x1)))

		// Calculate the absolute vertical distance between the two points
		// abs() ensures the distance is always positive regardless of direction
		dy := int(math.Abs(float64(y2 - y1)))

		// Calculate time for this segment using Chebyshev distance
		// Time = max(dx, dy) because we can move diagonally
		// - Move diagonally for min(dx, dy) steps (covers both directions)
		// - Then move in remaining direction for |dx - dy| steps
		// - Total = min(dx, dy) + |dx - dy| = max(dx, dy)
		segmentTime := int(math.Max(float64(dx), float64(dy)))

		// Add this segment's time to the total
		totalTime += segmentTime
	}

	// Return the total accumulated time to visit all points in order
	return totalTime
}

// Helper function to find maximum of two integers
// Go doesn't have a built-in max function, so we implement it
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function to find absolute value of an integer
// Go doesn't have a built-in abs function for integers, so we implement it
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Alternative optimized version using helper functions (avoids type conversion)
func minTimeToVisitAllPointsOptimized(points [][]int) int {
	totalTime := 0

	for i := 1; i < len(points); i++ {
		x1, y1 := points[i-1][0], points[i-1][1]
		x2, y2 := points[i][0], points[i][1]

		dx := abs(x2 - x1)
		dy := abs(y2 - y1)

		// Use helper max function
		totalTime += max(dx, dy)
	}

	return totalTime
}
```

### Go Code Explanation:

**Main Function Breakdown:**
- **`len(points)`**: Gets the number of points in the slice
- **`points[i-1][0]`**: Accesses x-coordinate of point at index i-1
- **`math.Abs()`**: Calculates absolute value (requires type conversion to float64)
- **`int(math.Max(...))`**: Finds maximum and converts result back to int
- **Loop starts at 1**: We need a previous point for distance calculation

**Type Conversions:**
- Go requires explicit type conversion: `float64(value)` and `int(result)`
- This is more verbose than other languages but ensures type safety

**Alternative Optimized Version:**
- Uses custom `abs()` and `max()` helper functions
- Avoids repeated `math` package calls
- Cleaner, more efficient code (avoids type conversions)
- **Recommended for production** use

### Comparison: `math.Abs()` vs Custom `abs()`

```go
// Option 1: Using math.Abs() (requires type conversion)
dx := int(math.Abs(float64(x2 - x1)))  // More conversions
dy := int(math.Abs(float64(y2 - y1)))
distance := int(math.Max(float64(dx), float64(dy)))

// Option 2: Using custom helpers (cleaner)
dx := abs(x2 - x1)      // Direct calculation
dy := abs(y2 - y1)
distance := max(dx, dy)

// Option 2 is preferred for Go development!
```

---

## ☕ Java Implementation

```java
class Solution {
    public int minTimeToVisitAllPoints(int[][] points) {
        // Initialize total time counter
        int totalTime = 0;
        
        // Loop through each point starting from the second point
        for (int i = 1; i < points.length; i++) {
            // Get coordinates of previous point (where we're coming from)
            int x1 = points[i - 1][0];
            int y1 = points[i - 1][1];
            
            // Get coordinates of current point (where we're going)
            int x2 = points[i][0];
            int y2 = points[i][1];
            
            // Calculate absolute horizontal distance
            int dx = Math.abs(x2 - x1);
            
            // Calculate absolute vertical distance
            int dy = Math.abs(y2 - y1);
            
            // Time to reach current point = max(horizontal, vertical)
            // This works because we can move diagonally
            totalTime += Math.max(dx, dy);
        }
        
        // Return total time to visit all points
        return totalTime;
    }
}
```

### Java Code Explanation:
- **`Math.abs()`**: Java's built-in absolute value function
- **`Math.max()`**: Java's built-in maximum value function
- **Array Declaration**: `int[][]` represents a 2D array
- **`points.length`**: Number of points to visit
- Everything else follows the same logic as the Go version

---

## 🔄 Complexity Analysis

| Aspect | Value |
|--------|-------|
| **Time Complexity** | O(n) where n = number of points |
| **Space Complexity** | O(1) - only using constant extra space |
| **Algorithm Type** | Greedy |

**Why O(n)?** We iterate through each point once, performing constant-time operations for each point.

---

## 🏢 Frequency Asked in Companies

This problem is asked **moderately frequently** in technical interviews, particularly for:

### Companies Known to Ask This Problem:

| Company | Frequency | Role |
|---------|-----------|------|
| **Google** | Medium | SDE, SDE-II |
| **Amazon** | Medium | SDE, SDE-II |
| **Meta** | Low-Medium | Engineer, Senior Engineer |
| **Microsoft** | Low-Medium | Software Engineer |
| **Apple** | Low | Engineer |
| **Adobe** | Low-Medium | Engineer |
| **Bloomberg** | Low | Engineer |
| **Goldman Sachs** | Low | Analyst, Engineer |
| **Uber** | Low-Medium | Software Engineer |
| **Airbnb** | Low | Engineer |

### Why Companies Ask This:
1. **Tests Understanding of Geometry**: Chebyshev distance, Manhattan distance concepts
2. **Greedy Algorithm Thinking**: Identifying optimal substructure
3. **Mathematical Insight**: Understanding why `max(dx, dy)` is the answer
4. **Implementation Skills**: Clean, bug-free code writing
5. **Easy to Medium Level**: Good filter question to assess problem-solving ability

### Interview Tips:
- ✅ **Explain the math**: Why diagonal movement is optimal
- ✅ **Draw diagrams**: Show how diagonal movement saves time
- ✅ **Test edge cases**: Single point, two points, collinear points
- ✅ **Mention complexity**: Show understanding of time/space analysis
- ✅ **Discuss alternatives**: Mention why not using Manhattan distance

---

## 🧪 Test Cases

### Test Case 1: Basic Example
```
Input: points = [[1,1],[3,4],[-1,0]]
Output: 7

Explanation:
Path visualization:
     y
   5 |
   4 | ●(3,4) ← Point 2
   3 |  /↓\
   2 | / ↓ \
   1 |●(1,1)→●→●
   0 | ← ← [−1,0]● ← Point 3
    ━━━━━━━━━━━━━━━
       x
       0  1  2  3

Leg 1: [1,1]→[3,4] = max(2,3) = 3 sec
Leg 2: [3,4]→[-1,0] = max(4,4) = 4 sec
Total: 7 seconds ✓
```

### Test Case 2: Horizontal Movement Only
```
Input: points = [[3,2],[-2,2]]
Output: 5

Explanation:
Path visualization:
     y
   3 |
   2 |●(3,2)────────●(-2,2)
   1 |      ← ← ← ←
   0 |________________
      -2  -1  0  1  2  3   x

Distance: dx = |-2-3| = 5, dy = |2-2| = 0
Time: max(5, 0) = 5 seconds ✓
No vertical movement needed!
```

### Test Case 3: Single Point (Edge Case)
```
Input: points = [[0,0]]
Output: 0

Explanation:
     y
   1 |
   0 |●(0,0) ← Already at first point
  -1 |_________
    -1  0  1   x

No movement needed!
Total: 0 seconds ✓
```

### Test Case 4: All Points Same (Edge Case)
```
Input: points = [[0,0],[0,0],[0,0]]
Output: 0

Explanation:
     y
   1 |
   0 |●●● All same point
  -1 |_________
    -1  0  1   x

Loop iteration 1: [0,0]→[0,0] = max(0,0) = 0
Loop iteration 2: [0,0]→[0,0] = max(0,0) = 0
Total: 0 seconds ✓
```

### Test Case 5: Diagonal Path
```
Input: points = [[0,0],[1,1],[2,2],[3,3]]
Output: 3

Explanation:
Path visualization:
     y
   3 |●───(3,3)
   2 |  \ 
   1 |   ●(2,2)
   0 |●───●───●(1,1)
      0   1   2  3   x
      (0,0)

Step 1: [0,0]→[1,1] = max(1,1) = 1 sec (diagonal)
Step 2: [1,1]→[2,2] = max(1,1) = 1 sec (diagonal)
Step 3: [2,2]→[3,3] = max(1,1) = 1 sec (diagonal)
Total: 3 seconds ✓

Perfect diagonal movement at each step!
```

---

## 📌 Key Takeaways

✅ **Core Concept**: Chebyshev distance is optimal when diagonal movement is allowed  
✅ **Formula**: Time = max(|x2-x1|, |y2-y1|) for each segment  
✅ **Approach**: Greedy - optimal locally leads to global optimum  
✅ **Complexity**: Linear time, constant space  
✅ **Interview Value**: Tests both math understanding and coding ability

---

## 🚀 Quick Reference Card

### Formula Summary
```
For each pair of consecutive points (x1,y1) and (x2,y2):
    
    Time = max(|x2 - x1|, |y2 - y1|)
    
    Where:
    |x2 - x1| = horizontal distance
    |y2 - y1| = vertical distance
    
    Total Time = Sum of all segment times
```

### Decision Tree
```
                    Need to move?
                        |
                        v
                    Is dx == dy?
                   /           \
                 YES            NO
                  |              |
    Perfect       |              v
    Diagonal!    |         Which is larger?
                 |         /           \
    Time =      |       dx > dy       dy > dx
    max(dx,dy)  |       /               \
                |      /                 \
              Time =  dx              Time = dy
              sec     sec               sec
```

### Visual Algorithm Summary
```
┌─────────────────────────────────────────────────────┐
│              ALGORITHM STEPS                        │
├─────────────────────────────────────────────────────┤
│ 1. Initialize totalTime = 0                         │
│ 2. Loop from point[1] to point[n-1]                 │
│ 3. For each iteration:                              │
│    a) Get current and previous coordinates          │
│    b) Calculate dx = |x2 - x1|                      │
│    c) Calculate dy = |y2 - y1|                      │
│    d) Add max(dx, dy) to totalTime                  │
│ 4. Return totalTime                                 │
└─────────────────────────────────────────────────────┘
```

### Common Mistakes to Avoid
```
❌ WRONG: Using dx + dy (Manhattan distance)
   - Ignores diagonal movement capability
   - Gives incorrect (higher) results

❌ WRONG: Using sqrt(dx² + dy²) (Euclidean distance)
   - More accurate geometrically but wrong here
   - Diagonal movement is discrete (1 sec), not continuous

❌ WRONG: Starting loop from index 0
   - No previous point to compare with
   - Should start from index 1

❌ WRONG: Not using absolute values
   - Negative distances don't make sense
   - |x2-x1| handles all direction combinations

✅ RIGHT: Using max(dx, dy) (Chebyshev distance)
   - Accounts for diagonal movement
   - Optimal for 8-directional movement
```

### Implementation Checklist
```
□ Understand Chebyshev distance concept
□ Know why max() is used instead of sum()
□ Implement absolute value calculation
□ Implement max value calculation
□ Loop starts from index 1 (not 0)
□ Accumulate time for all segments
□ Return total accumulated time
□ Test with edge cases (1 point, same points, etc.)
□ Verify time complexity is O(n)
□ Verify space complexity is O(1)
```

### Similar Problems to Practice
```
1. LeetCode 1266 (This problem) - Chebyshev distance
2. LeetCode 483 - Valid Palindrome III (Dynamic Programming)
3. LeetCode 163 - Missing Ranges (Array manipulation)
4. LeetCode 1037 - Valid Boomerang (Geometry)
5. Minimum Steps to One (Greedy approach)
```

---

## 💬 How to Explain This in an Interview

### 30-Second Version
"For each pair of consecutive points, we calculate the horizontal and vertical distances. Since we can move diagonally, the time needed equals the maximum of these two distances. We sum up times for all segments."

### 2-Minute Version
"We use the Chebyshev distance (max of absolute differences) instead of Manhattan distance. This is optimal because diagonal movement covers both directions simultaneously, making it more efficient. The algorithm iterates through consecutive point pairs, calculates the maximum of horizontal and vertical distances for each segment, and sums them up. Time complexity is O(n) since we process each point once."

### 5-Minute Version (with diagram)
"The key insight is that when you can move diagonally, the optimal path moves diagonally as much as possible, then continues in one direction for the remaining distance. The time equals the longer of the two distances.

[Draw a diagram showing movement from (0,0) to (3,5)]

This can be calculated simply as max(|3-0|, |5-0|) = max(3,5) = 5 seconds.

The algorithm iterates through each consecutive pair of points, applies this formula, and accumulates the total time. The time complexity is O(n) - linear in the number of points - and we use constant space. This is optimal and cannot be improved."

---
