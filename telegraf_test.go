package telegraf

import (
	"os"
	"testing"
)

var botAPI *BotAPI

func TestCreateNewBot(t *testing.T) {
	var err error

	token, _ := os.ReadFile(".env")
	botAPI, err = CreateNewBot(string(token))

	if err != nil {
		t.Error("Test 'TestCreateNewBot' failed")
	}
}
