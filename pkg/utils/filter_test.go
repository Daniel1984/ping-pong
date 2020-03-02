package utils_test

import (
	"testing"

	"github.com/ping-pong/pkg/utils"
	"gotest.tools/assert"
)

func TestGetUniqueInts(t *testing.T) {
	testcases := []struct {
		testMsg        string
		value          []int
		expectedOutput []int
	}{
		{
			testMsg:        "no duplicates",
			value:          []int{1, 2, 3, 4},
			expectedOutput: []int{1, 2, 3, 4},
		},
		{
			testMsg:        "contains duplicates",
			value:          []int{1, 2, 3, 4, 4, 4},
			expectedOutput: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testMsg, func(t *testing.T) {
			got := utils.GetUniqueInts(tc.value)
			assert.DeepEqual(t, tc.expectedOutput, got)
		})
	}
}
