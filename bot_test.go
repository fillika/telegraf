package telegraf

import (
	"os"
	"strconv"
	"testing"
)

var botApi *BotAPI

func TestMain(m *testing.M) {
	var err error

	token, err := os.ReadFile(".env")

	if err != nil {
		panic(err)
	}

	botApi, err = NewBot(string(token))

	if err != nil {
		panic(err)
	}

	// Call the other tests and exit with their status code
	os.Exit(m.Run())
}

func TestGetUpdatesChannel(t *testing.T) {
	_, err := botApi.GetUpdatesChannel(UpdatesConfig{Offset: 0})

	if err != nil {
		t.Error("Test 'GetUpdatesChannel' failed")
	}
}

func TestSendMessage(t *testing.T) {
	data, err := os.ReadFile("chat-id.env")

	chatID, _ := strconv.Atoi(string(data))

	if err != nil {
		t.Error("Test 'SendMessage' failed. Error:", err)
	}

	_, err = botApi.SendMessage(MessageConfig{
		ChatID: chatID,
		Text:   "Hello, world!",
	})

	if err != nil {
		t.Error("Test 'SendMessage' failed", err)
	}
}
