package leetcode

import (
	"fmt"
	"testing"
)

// TestFindCircleNum 测试 findCircleNum 函数
func TestFindCircleNum(t *testing.T) {
	tests := []struct {
		isConnected [][]int
		expected    int
	}{
		{
			isConnected: [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}},
			expected:    1,
		},
		{
			isConnected: [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			expected:    2,
		},
		{
			isConnected: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			expected:    3,
		},
		{
			isConnected: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
			expected:    1,
		},
		{
			isConnected: [][]int{{1}},
			expected:    1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.isConnected), func(t *testing.T) {
			result := findCircleNum(tt.isConnected)
			if result != tt.expected {
				t.Errorf("findCircleNum(%v) = %d, want %d", tt.isConnected, result, tt.expected)
			}
		})
	}
}
