# Wave Pattern Array Rearrangement

This Go package implements a solution for rearranging an array into a "wave" pattern according to specified criteria. The implementation handles the arrangement of integers into blocks of size 2x+1, where x is a given parameter ≥ 1.

## Problem Description

Given an array of integers and an integer x (where x ≥ 1), rearrange the array into a "wave" pattern as follows:

1. **Partitioning**: Divide the array into consecutive blocks of 2x+1 elements.
2. **Wave Pattern in Each Block**: For every block (with indices 0 to 2x):
   - The element at index x must be the maximum in that block
   - The elements to the left of index x must be arranged in non-decreasing order
   - The elements to the right of index x must be arranged in non-increasing order

## Installation

```bash
git clone   https://github.com/abd-ulbasit/wavepattern 
cd wave-pattern
```


## Running Tests

To run the tests, navigate to the project directory and execute:

```bash
go test ./wave-pattern -v
```

This will run all test cases and display detailed output including:
- Example from the problem statement
- Invalid input validation
- Multiple block handling
- Pattern validation


## Usage

```go
package main

import (
    "fmt"
    "github.com/abd-ulbasit/wavepattern"
)

func main() {
    arr := []int{3, 6, 5, 10, 7}
    x := 2

    result, err := wavepattern.ArrangeWavePattern(arr, x)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Printf("Rearranged array: %v\n", result)
}
```


## Implementation Details

The solution implements:
1. Input validation for x ≥ 1
2. Array length validation (must be multiple of 2x+1)
3. Block-wise processing
4. Proper arrangement of elements within each block
5. Preservation of original array (returns new array)

## Example

Input:
- Array: [3, 6, 5, 10, 7]
- x = 2

Output:
- [3, 6, 10, 7, 5]

In this output:
- 10 is at index 2 (x)
- [3, 6] is in non-decreasing order
- [7, 5] is in non-increasing order