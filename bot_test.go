package telegraf

import (
	"os"
	"strconv"
	"testing"
)

var botApi *BotAPI
var chatID int

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

	data, _ := os.ReadFile("chat-id.env")
	chatID, err = strconv.Atoi(string(data))

	if err != nil {
		panic(err)
	}

	// Call the other tests and exit with their status code
	os.Exit(m.Run())
}

func TestGetMe(t *testing.T) {
	_, err := botApi.GetMe()

	if err != nil {
		t.Error("Test 'GetMe' failed", err)
	}
}

func TestCopyMessage(t *testing.T) {
	msg := NewCopyMessage(chatID, chatID, 11)
	msg.DisableNotification = true

	_, err := botApi.CopyMessage(msg)

	if err != nil {
		t.Error("Test 'CopyMessage' failed", err)
	}
}

func TestForwardMessage(t *testing.T) {
	msg := NewForwardMessage(chatID, chatID, 10)
	msg.DisableNotification = true

	_, err := botApi.ForwardMessage(msg)

	if err != nil {
		t.Error("Test 'ForwardMessage' failed", err)
	}
}

func TestGetUpdatesChannel(t *testing.T) {
	_, err := botApi.GetUpdatesChannel(UpdatesConfig{Offset: 0})

	if err != nil {
		t.Error("Test 'GetUpdatesChannel' failed")
	}
}

func TestSendMessage(t *testing.T) {
	msg := NewMessage(chatID, "Hello, world!")

	_, err := botApi.SendMessage(msg)

	if err != nil {
		t.Error("Test 'SendMessage' failed", err)
	}
}
