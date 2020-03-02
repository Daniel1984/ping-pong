package cmdparser

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

// CmdParser - holds pointer to bufio reader used to get users input
type CmdParser struct {
	reader *bufio.Reader
}

// New returns pointer to new instance of CmdParser
func New(stdin io.Reader) *CmdParser {
	return &CmdParser{
		reader: bufio.NewReader(stdin),
	}
}

// AskForUserID prompts user to enter id and runs validations against it
func (cmdp *CmdParser) AskForUserID() (int, error) {
	fmt.Print("Enter your user ID: ")
	uid, _ := cmdp.reader.ReadString('\n')

	uidInt, err := strconv.Atoi(strings.TrimSpace(uid))
	if err != nil {
		return 0, fmt.Errorf("user ID must be a valid integer: [0..%d]", math.MaxInt64)
	}

	return uidInt, nil
}

// AskForFriendIDs - prompts user to enter coma separated ints as his friend
// ids, runs validations and returns list of friend ids
func (cmdp *CmdParser) AskForFriendIDs(uid int) ([]int, error) {
	fmt.Print("Enter your friend ID or ID's separated by coma: ")
	idStr, _ := cmdp.reader.ReadString('\n')

	splitIds := strings.Split(strings.TrimSpace(idStr), ",")
	friendIds := []int{}

	for _, id := range splitIds {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("friend user IDs must be valid integers: [0..%d] separated by coma", math.MaxInt64)
		}

		if idInt == uid {
			return nil, fmt.Errorf("you can't be in your firends list")
		}

		friendIds = append(friendIds, idInt)
	}

	return friendIds, nil
}
