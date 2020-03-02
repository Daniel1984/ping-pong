package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ping-pong/cmd/client/cmdparser"
	"github.com/ping-pong/pkg/client"
	"github.com/ping-pong/pkg/models"
	"github.com/ping-pong/pkg/sigtermhandler"
	"github.com/ping-pong/pkg/utils"
)

/*
* program designed to ask user for his and his friend IDs. When data is
* collected and validated, call to server, with appropriate payload, is made to
* establish connection and listen for incoming messages
 */
func main() {
	cmdp := cmdparser.New(os.Stdin)

	uid, err := cmdp.AskForUserID()
	checkFatal(err)

	friendIDs, err := cmdp.AskForFriendIDs(uid)
	checkFatal(err)

	payload := &models.Message{
		UserID:  uid,
		Friends: utils.GetUniqueInts(friendIDs),
	}

	payloadBytes, err := json.Marshal(payload)
	checkFatal(err)

	c, err := client.Start(payloadBytes)
	checkFatal(err)

	sigtermhandler.Init(func() {
		log.Print("closing down connection")
		c.Close()
	})
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
