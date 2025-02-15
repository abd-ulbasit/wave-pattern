// main.go
package wavepattern

import (
	"errors"
	"sort"
)

// ArrangeWavePattern rearranges an array into a wave pattern based on the given x value
// Returns error if input array length is not a multiple of (2x+1) or if x < 1
func ArrangeWavePattern1(arr []int, x int) ([]int, error) {
	if x < 1 {
		return nil, errors.New("x must be greater than or equal to 1")
	}

	blockSize := 2*x + 1
	if len(arr)%blockSize != 0 {
		return nil, errors.New("array length must be a multiple of (2x+1)")
	}

	result := make([]int, len(arr))
	copy(result, arr)

	// Process each block
	for i := 0; i < len(arr); i += blockSize {
		block := result[i : i+blockSize]
		rearrangeBlock(block, x)
	}

	return result, nil
}

// rearrangeBlock handles the wave pattern arrangement for a single block
func rearrangeBlock(block []int, x int) {
	// Find the maximum element and its position
	maxIdx := 0
	maxVal := block[0]
	for i, val := range block {
		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}

	// Create temporary slice for the rearranged block
	temp := make([]int, len(block))
	copy(temp, block)

	// Place the maximum element at index x
	block[x] = maxVal

	// Sort left side in non-decreasing order
	leftSide := make([]int, 0, x)
	for i := 0; i < len(block); i++ {
		if i != maxIdx && temp[i] != maxVal {
			leftSide = append(leftSide, temp[i])
		}
		if len(leftSide) == x {
			break
		}
	}
	sort.Ints(leftSide)
	copy(block[:x], leftSide)

	// Sort right side in non-increasing order
	rightSide := make([]int, 0, x)
	for i := 0; i < len(block); i++ {
		if i != maxIdx && temp[i] != maxVal && len(rightSide) < x {
			rightSide = append(rightSide, temp[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rightSide)))
	copy(block[x+1:], rightSide)
}

func IsValidWavePattern(arr []int, x int) bool {
	if x < 1 || len(arr)%(2*x+1) != 0 {
		return false
	}

	blockSize := 2*x + 1

	// Check each block
	for i := 0; i < len(arr); i += blockSize {
		end := i + blockSize
		if end > len(arr) {
			end = len(arr)
		}

		block := arr[i:end]
		if !isValidBlock(block, x) {
			return false
		}
	}

	return true
}

func isValidBlock(block []int, x int) bool {
	if len(block) != 2*x+1 {
		return false
	}

	// Check if element at index x is maximum
	peak := block[x]
	for _, v := range block {
		if v > peak {
			return false
		}
	}

	// Check left side (non-decreasing)
	for i := 1; i < x; i++ {
		if block[i] < block[i-1] {
			return false
		}
	}

	// Check right side (non-increasing)
	for i := x + 1; i < len(block)-1; i++ {
		if block[i] < block[i+1] {
			return false
		}
	}

	return true
}

func ArrangeWavePattern(arr []int, x int) ([]int, error) {
	if x < 1 {
		return nil, errors.New("x must be greater than or equal to 1")
	}

	blockSize := 2*x + 1
	if len(arr)%blockSize != 0 {
		return nil, errors.New("array length must be a multiple of (2x+1)")
	}

	// Sort in ascending order
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	result := make([]int, len(arr))
	right := len(arr) - 1

	// Fill peaks first (working with largest numbers)
	for i := 0; i*blockSize < len(arr); i++ {
		peakIndex := i*blockSize + x
		result[peakIndex] = sorted[right]
		right--
	}

	// Fill remaining positions alternating left and right of peaks
	for dist := 1; dist <= x; dist++ {
		for i := 0; i*blockSize < len(arr); i++ {
			peakIndex := i*blockSize + x

			// Place element to the left of peak
			leftIndex := peakIndex - dist
			if leftIndex >= i*blockSize {
				result[leftIndex] = sorted[right]
				right--
			}

			// Place element to the right of peak
			rightIndex := peakIndex + dist
			if rightIndex < (i+1)*blockSize {
				result[rightIndex] = sorted[right]
				right--
			}
		}
	}

	return result, nil
}
