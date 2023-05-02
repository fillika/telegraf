package telegraf

import (
	"fmt"
	"os"
	"testing"
)

var bot *BotAPI

func Test_createUrlWithTokenAndMethod(t *testing.T) {
	token := "x2222hsf"
	url := createUrlWithTokenAndMethod(token, methods.getMe)

	if url != fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methods.getMe) {
		t.Error("Incorrect url")
	}
}

func Test_getMe(t *testing.T) {
	token, _ := os.ReadFile(".env")

	var err error
	bot, err = getMe(string(token))

	if err != nil {
		t.Error("Error while getting me")
	}
}

func Test_getUpdates(t *testing.T) {
	_, err := getUpdates(bot)

	if err != nil {
		t.Error("Error while getting updates")
	}
}
