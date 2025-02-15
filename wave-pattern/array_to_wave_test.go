// main_test.go
package wavepattern

import (
	"testing"
)

func TestArrangeWavePattern(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		x       int
		wantErr bool
	}{
		{
			name:    "Example from problem statement",
			arr:     []int{3, 6, 5, 10, 7},
			x:       2,
			wantErr: false,
		},
		{
			name:    "Invalid x value",
			arr:     []int{1, 2, 3},
			x:       0,
			wantErr: true,
		},
		{
			name:    "Invalid array length",
			arr:     []int{1, 2, 3, 4},
			x:       2,
			wantErr: true,
		},
		{
			name:    "Multiple blocks",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:       2,
			wantErr: false,
		},
		{
			name:    "Multiple blocks",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			x:       3,
			wantErr: false,
		},
		{
			name:    "Multiple blocks",
			arr:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
			x:       3,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ArrangeWavePattern(tt.arr, tt.x)
			// print input and output side by side
			t.Logf("Input: %v\nOutput: %v\n", tt.arr, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ArrangeWavePattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Verify the result follows wave pattern rules
				if !IsValidWavePattern(got, tt.x) {
					t.Errorf("ArrangeWavePattern() result does not follow wave pattern rules: %v", got)
				}

				// Verify all original elements are present
				if !haveSameElements(got, tt.arr) {
					t.Errorf("ArrangeWavePattern() result has different elements than input: got %v, want elements from %v", got, tt.arr)
				}
			}
		})
	}
}

// haveSameElements checks if two slices have the same elements (ignoring order)
func haveSameElements(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[int]int)

	// Count elements in first slice
	for _, v := range a {
		countMap[v]++
	}

	// Subtract counts for second slice
	for _, v := range b {
		countMap[v]--
		if countMap[v] == 0 {
			delete(countMap, v)
		}
	}

	return len(countMap) == 0
}

func TestIsValidWavePattern(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		x        int
		expected bool
	}{
		{
			name:     "Valid pattern 1",
			arr:      []int{3, 6, 10, 7, 5},
			x:        2,
			expected: true,
		},
		{
			name:     "Valid pattern 2",
			arr:      []int{3, 6, 10, 6, 3},
			x:        2,
			expected: true,
		},
		{
			name:     "Invalid - peak not maximum",
			arr:      []int{3, 6, 5, 7, 4},
			x:        2,
			expected: false,
		},
		{
			name:     "Invalid - left side not non-decreasing",
			arr:      []int{6, 3, 10, 7, 5},
			x:        2,
			expected: false,
		},
		{
			name:     "Invalid - right side not non-increasing",
			arr:      []int{3, 6, 10, 5, 7},
			x:        2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidWavePattern(tt.arr, tt.x)
			if result != tt.expected {
				t.Errorf("IsValidWavePattern() = %v, want %v for array %v", result, tt.expected, tt.arr)
			}
		})
	}
}
