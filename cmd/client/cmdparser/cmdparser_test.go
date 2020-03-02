package cmdparser_test

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/ping-pong/cmd/client/cmdparser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAskForUserID(t *testing.T) {
	testcases := []struct {
		testMsg     string
		inputVal    string
		expectedID  int
		expectedErr func(*testing.T, error)
	}{
		{
			testMsg:    "string",
			inputVal:   "foo\n",
			expectedID: 0,
			expectedErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		{
			testMsg:    "int with string in string",
			inputVal:   "1a\n",
			expectedID: 0,
			expectedErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		{
			testMsg:    "number too large",
			inputVal:   fmt.Sprintf("%d%d\n", math.MaxInt64, 1),
			expectedID: 0,
			expectedErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		{
			testMsg:    "success",
			inputVal:   "1\n",
			expectedID: 1,
			expectedErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testMsg, func(t *testing.T) {
			stdin := &bytes.Buffer{}
			stdin.Write([]byte(tc.inputVal))

			cmdp := cmdparser.New(stdin)
			gotID, err := cmdp.AskForUserID()

			tc.expectedErr(t, err)
			assert.Equal(t, tc.expectedID, gotID)
		})
	}
}

func TestAskForFriendIDs(t *testing.T) {
	testcases := []struct {
		testMsg        string
		inputVal       string
		inputUserID    int
		expectedOutput []int
		expectedErr    func(*testing.T, error)
	}{
		{
			testMsg:        "contains non ints",
			inputVal:       "1,2,a,3",
			inputUserID:    4,
			expectedOutput: nil,
			expectedErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		{
			testMsg:        "contains self in friends list",
			inputVal:       "1,2,3,4",
			inputUserID:    4,
			expectedOutput: nil,
			expectedErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		{
			testMsg:        "valid input",
			inputVal:       "1,2,3,4",
			inputUserID:    5,
			expectedOutput: []int{1, 2, 3, 4},
			expectedErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testMsg, func(t *testing.T) {
			stdin := &bytes.Buffer{}
			stdin.Write([]byte(tc.inputVal))

			cmdp := cmdparser.New(stdin)
			gotID, err := cmdp.AskForFriendIDs(tc.inputUserID)

			tc.expectedErr(t, err)
			assert.Equal(t, tc.expectedOutput, gotID)
		})
	}
}
