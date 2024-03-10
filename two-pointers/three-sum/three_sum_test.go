package three_sum

import (
	"testing"
)

func TestFindSumOfThree(t *testing.T) {

	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			got := findSumOfThree(tt.input, tt.target)
			if got != tt.expected {
				t.Errorf("findSumOfThree() = %v, want %v", got, tt.expected)
			}
		})
	}
}
